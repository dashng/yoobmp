package bmp

type BmpHandler interface {
	UnmarshalCommonHeader(headerData []byte) *CommonHeader
}

type Handler struct {
}

func (bmpHandler *Handler) UnmarshalCommonHeader(headerData []byte) *CommonHeader {
	bmpHeader, _ := UnmarshalCommonHeader(headerData)
	return bmpHeader
}
