package blockchain

import "net/url"
import "net/http"

type Requestor interface {
	Do(req *http.Request) (*http.Response, error)
}
type BlockChainClient struct {
	httpClient Requestor
	endpoint   url.URL
}

func Block()  {
	
}