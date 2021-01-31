package main

import (
	"github.com/dashng/yoobmp/bmp"
	"github.com/dashng/yoobmp/bmpserver"
)

func main() {
	server, _ := bmpserver.NewYooServer(32124)
	server.Handle(&bmp.Handler{})
}
