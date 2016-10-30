package blockchain

import (
	"io"
	"encoding/json"
)

type ResponseHandler struct {
	json.Decoder
	json.Encoder
	io.Closer

}

func NewReadWriteCloser() ResponseHandler {
	return ResponseHandler{}
}
