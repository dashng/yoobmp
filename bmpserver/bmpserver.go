package bmpserver

import (
	"fmt"
	"io"
	"log"
	"net"

	"github.com/dashng/yoobmp/bmp"
	"github.com/golang/glog"
)

// TCPServer interface
type TCPServer interface {
	run()
	Handle(bmpHandle bmp.BmpHandler)
}

// YooServer listening to tcp port for receiving bmp data
type YooServer struct {
	tcpPort          int32
	tcpListener      net.TCPListener
	commonHeaderData []byte
	bmpHandler       bmp.BmpHandler
}

func (bmpServer *YooServer) getTCPListener(port int32) *net.TCPListener {
	addr, err := net.ResolveTCPAddr("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		panic(err)
	}
	l, err := net.ListenTCP("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}
	return l
}

func (bmpServer *YooServer) run() {
	defer bmpServer.tcpListener.Close()
	for {
		conn, err := bmpServer.tcpListener.Accept()
		if err != nil {
			fmt.Println(err)
			log.Fatal(err)
		}
		go bmpServer.worker(conn)
	}
}

func (bmpServer *YooServer) worker(client net.Conn) {
	for {
		headerMsg := make([]byte, bmp.CommonHeaderLength)
		if _, err := io.ReadAtLeast(client, headerMsg, bmp.CommonHeaderLength); err != nil {
			glog.Errorf("fail to read from client %+v with error: %+v", client.RemoteAddr(), err)
			return
		}
		fmt.Println(headerMsg)
		// Recovering common header first
		header, err := bmpServer.bmpHandler.UnmarshalCommonHeader(headerMsg[:bmp.CommonHeaderLength])
		if err != nil {
			glog.Errorf("fail to recover BMP message Common Header with error: %+v", err)
			continue
		}
		// Allocating space for the message body
		msg := make([]byte, int(header.MessageLength)-bmp.CommonHeaderLength)
		if _, err := io.ReadFull(client, msg); err != nil {
			glog.Errorf("fail to read from client %+v with error: %+v", client.RemoteAddr(), err)
			return
		}

		fullMsg := make([]byte, int(header.MessageLength))
		copy(fullMsg, headerMsg)
		copy(fullMsg[bmp.CommonHeaderLength:], msg)
	}
}

// Handle parse the received bmp data
func (bmpServer *YooServer) Handle(bmpHandler bmp.BmpHandler) {
	bmpServer.bmpHandler = bmpHandler
	bmpServer.run()
}

// NewYooServer initialize YooServer
func NewYooServer(port int32) (TCPServer, error) {
	yooServer := YooServer{}
	yooServer.tcpListener = *(yooServer.getTCPListener(port))
	return &yooServer, nil
}
