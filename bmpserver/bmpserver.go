package bmpserver

import (
	"fmt"
	"io"
	"log"
	"net"

	"github.com/dashng/yoobmp/bmp"
)

// TCPServer interface
type TCPServer interface {
	run()
	Handle(bmpHandle bmp.BmpHandler)
}

// YooServer listening to tcp port for receiving bmp data
type YooServer struct {
	tcpPort          int32
	tcpListener      net.Listener
	commonHeaderData []byte
	bmpHandler       bmp.BmpHandler
}

func (bmpServer *YooServer) getTCPListener(port int32) net.Listener {
	client, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatal(err)
	}
	return client
}

func (bmpServer *YooServer) run() {
	// defer bmpServer.tcpListener.Close()
	for {
		conn, err := bmpServer.tcpListener.Accept()
		if err != nil {
			fmt.Println(err)
			log.Fatal(err)
		}

		fmt.Println("555555555555")
		go bmpServer.worker(conn)
	}
}

func (bmpServer *YooServer) worker(conn net.Conn) {
	defer conn.Close()
	for {
		commonHeaderMsg := make([]byte, bmp.CommonHeaderLength)
		if _, err := io.ReadAtLeast(conn, commonHeaderMsg, bmp.CommonHeaderLength); err != nil {
			fmt.Println("fail to read from client %+v with error: %+v", conn.RemoteAddr(), err)
			return
		}
		bmpHeader, handlerErr := bmpServer.bmpHandler.UnmarshalCommonHeader(commonHeaderMsg[:bmp.CommonHeaderLength])
		if handlerErr != nil {
			fmt.Println("parse header error: %+v", handlerErr)
			continue
		}
		fmt.Println(bmpHeader)
		fmt.Println(commonHeaderMsg[:bmp.CommonHeaderLength])
		// bmpBody := make([]byte, bmpHeader.MessageLength)
		// count, err := io.ReadFull(conn, bmpBody)
		// fmt.Println(count)
		// fmt.Println("===============")
		// fmt.Println(bmpBody)
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
	yooServer.tcpListener = yooServer.getTCPListener(port)
	return &yooServer, nil
}
