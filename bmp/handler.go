package bmp

import "fmt"

type BmpHandler interface {
	UnmarshalCommonHeader(headerData []byte)
}

type Handler struct {
}

func (bmpHandler *Handler) UnmarshalCommonHeader(headerData []byte) {
	fmt.Println("===============", headerData)
}
