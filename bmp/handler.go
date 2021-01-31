package bmp

import "fmt"

type BmpHandler interface {
	UnmarshalCommonHeader(headerData []byte)
}

type Handler struct {
}

func (bmpHandler *Handler) UnmarshalCommonHeader(headerData []byte) {
	header, _ := UnmarshalCommonHeader(headerData)
	fmt.Println(headerData)
	fmt.Println(header)
}
