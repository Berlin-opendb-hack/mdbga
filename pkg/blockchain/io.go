package blockchain

type ResponseHandler struct {
}

func NewReadWriteCloser() *ResponseHandler {
	return &ResponseHandler{}
}

func (h *ResponseHandler) Read(p []byte) (int, error) {
	return 0, nil
}
func (h *ResponseHandler) Write(p []byte) (int, error) {
	return 0, nil
}
func (h *ResponseHandler) Close() error {
	return nil
}