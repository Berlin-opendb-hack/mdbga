package blockchain

import "net/url"
import (
	"net/http"
	"github.com/shopspring/decimal"
	"net/rpc"
	"time"
)

type Transaction struct {
	Amount decimal.Decimal `json:"amount"`
	Address string `json:"amount"`
	Date time.Time `json:"date"`
	Fee decimal.Decimal `json:"fee_amount"`
	ExchangeRate decimal.Decimal `json:"exchange_rate"`
}

type Requestor interface {
	Do(req *http.Request) (*http.Response, error)
}
type BlockChainClient struct {
	rpcClient *rpc.Client
	endpoint   url.URL
}

func NewBlockChainClient(endpoint url.URL, rpcClient *rpc.Client) (*BlockChainClient, error) {
	return &BlockChainClient{
		rpcClient : rpcClient,
		endpoint: endpoint,
	}, nil
}

func (cl *BlockChainClient) SendToAddress(address string, amount decimal.Decimal) error  {
	return cl.rpcClient.Call("sendtoaddress", []string{address, amount.StringFixed(2)}, nil)
}
func (cl *BlockChainClient) ListTransactions() ([]Transaction, error) {
	transactions := []Transaction{}
	err := cl.rpcClient.Call("listtransactions", nil, transactions)
	if nil != err {
		return []Transaction{}, err
	}
	return transactions, nil
}