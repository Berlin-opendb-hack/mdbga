package blockchain

import "net/url"
import (
	"net/http"
	"github.com/shopspring/decimal"
	"net/rpc"
)

type Transaction struct {
	amount string `json:"amount"`
}

type Requestor interface {
	Do(req *http.Request) (*http.Response, error)
}
type BlockChainClient struct {
	rpcClient rpc.Client
	endpoint   url.URL
}

func NewBlockChainClient(endpoint url.URL, rpcClient rpc.Client) (*BlockChainClient, error) {
	return &BlockChainClient{
		rpcClient : rpcClient,
		endpoint: endpoint,
	}
}

func (cl *BlockChainClient) SendToAddress(address string, amount decimal.Decimal)  {
	cl.rpcClient.Call("sendtoaddress", []string{address, amount.StringFixed(2)})
}
func (cl *BlockChainClient) ListTransactions()  {
	transactions := []Transaction{}
	cl.rpcClient.Call("listtransactions", nil, transactions)
}