package bmpserver

import (
	"bufio"
	"fmt"
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
		for {
			commonHeader := make([]byte, bmp.CommonHeaderLength)
			_, err = bufio.NewReader(conn).Read(commonHeader)
			if err != nil {
				log.Printf("Error: %+v", err.Error())
				return
			}
			bmpHeader := bmpServer.bmpHandler.UnmarshalCommonHeader(commonHeader)
			fmt.Println(bmpHeader.Serialize())
		}
	}
}

// Handle parse the received bmp data
func (bmpServer *YooServer) Handle(bmpHandler bmp.BmpHandler) {
	fmt.Println(bmpHandler)
	bmpServer.bmpHandler = bmpHandler
	bmpServer.run()
}

// NewYooServer initialize YooServer
func NewYooServer(port int32) (TCPServer, error) {
	yooServer := YooServer{}
	yooServer.tcpListener = *(yooServer.getTCPListener(port))
	return &yooServer, nil
}
