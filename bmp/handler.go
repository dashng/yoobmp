package bmp

type BmpHandler interface {
	UnmarshalCommonHeader(headerData []byte) (*CommonHeader, error)
}

type Handler struct {
}

func (bmpHandler *Handler) UnmarshalCommonHeader(headerData []byte) (*CommonHeader, error) {
	bmpHeader, err := UnmarshalCommonHeader(headerData)
	return bmpHeader, err
}
