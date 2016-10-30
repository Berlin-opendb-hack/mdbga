package main

import (
	"github.com/Berlin-opendb-hack/mdbga/app"
	"github.com/goadesign/goa"
	"github.com/Berlin-opendb-hack/mdbga/pkg/bank"
	"net/url"
	"os"
	"net/http"
	"github.com/shopspring/decimal"
	"github.com/pkg/errors"
	"time"
	"github.com/Berlin-opendb-hack/mdbga/pkg/blockchain"
	"net/rpc/jsonrpc"
	"github.com/Berlin-opendb-hack/opendb-mock/pkg/accounting"
)

// BlockchainController implements the blockchain resource.
type BlockchainController struct {
	*goa.Controller
}

// NewBlockchainController creates a blockchain controller.
func NewBlockchainController(service *goa.Service) *BlockchainController {


	return &BlockchainController{Controller: service.NewController("BlockchainController")}
}

// GetBlockchainTransfers runs the GetBlockchainTransfers action.
func (c *BlockchainController) GetBlockchainTransfers(ctx *app.GetBlockchainTransfersBlockchainContext) error {
	rpcClient := jsonrpc.NewClient(blockchain.NewReadWriteCloser())
	blockChn, err := blockchain.NewBlockChainClient(url.URL{
		Host: os.Getenv(blockchainHost),
		Scheme: os.Getenv(blockchainScheme),
		Path: os.Getenv(blockchainPath),
	}, rpcClient)
	if nil != err {
		return ctx.InternalServerError(err)
	}
	transactions, err := blockChn.ListTransactions()
	if nil != err {
		return ctx.InternalServerError(err)
	}
	res := app.OpendbHackTransferCollection{}
	for _, transaction := range transactions {
		dateString := transaction.Date.Format(time.RFC3339)
		trnsfer := app.OpendbHackTransfer {
			Date: &dateString,
			Amount: transaction.Amount.StringFixed(2),
			ExchangeRate: transaction.ExchangeRate.String(),
			Identifier: "",

		}
		res = append(res, &trnsfer)
	}
	return ctx.OK(res)
}

// PostBlockchainTransfer runs the PostBlockchainTransfer action.
func (c *BlockchainController) PostBlockchainTransfer(ctx *app.PostBlockchainTransferBlockchainContext) error {

	token := ctx.Value("token")
	if nil == token {
		return ctx.Unauthorized()
	}
	bankEndpoint := url.URL{
		Scheme: os.Getenv(bankScheme),
		Host: os.Getenv(bankHost),
		Path: os.Getenv(bankPath),
	}
	client := &http.Client{}
	bankClient := bank.NewBankClient(client, bankEndpoint)
	accounts, err := bankClient.GetAccounts(token.(string))
	if nil != err {
		return ctx.InternalServerError(err)
	}
	if 0 == len(accounts) {
		return ctx.InternalServerError(errors.New("Expecting at least one account"))
	}
	payload := ctx.Payload
	amount, err := decimal.NewFromString(payload.Amount)

	account := accounts[0]
	balance := decimal.NewFromFloat(account.Balance)
	if -1 == balance.Cmp(amount) {
		return ctx.BadRequest(errors.New("Amount not available in account"))
	}
	if nil != err {
		return ctx.InternalServerError(err)
	}
	transfer := bank.Transfer{
		Amount: amount.StringFixed(2),
		CreditorIBAN: os.Getenv(masterAccountIban),
		CreditorBIC: os.Getenv(masterAccountIban),
		CreditorName: os.Getenv(masterAccountHolder),
		DebtorIBAN: account.Iban,
		DebtorBIC: os.Getenv(defalultBic),
		DebtorName: "Kunde",
		Currency: "EUR",
		RemittanceInformation: payload.Address,
	}
	transferDate := time.Now()
	if nil != payload.Date {
		transferDate, err = time.Parse(time.RFC3339, *payload.Date)
		if nil != err {
			return ctx.InternalServerError(err)
		}
		if transferDate.Before(time.Now()) {
			transferDate = time.Now()
		}
	}
	resId, err := bankClient.CreateTransfer(token.(string), transfer)
	if nil != err {
		return ctx.InternalServerError(err)
	}
	rpcClient := jsonrpc.NewClient(blockchain.NewReadWriteCloser())
	blockChn, err := blockchain.NewBlockChainClient(url.URL{
		Host: os.Getenv(blockchainHost),
		Scheme: os.Getenv(blockchainScheme),
		Path: os.Getenv(blockchainPath),
	}, rpcClient)
	if nil != err {
		return ctx.InternalServerError(err)
	}
	blockChn.SendToAddress(payload.Address,amount)
	ledger := accounting.GetLedger()
	ledger.Book(transfer.CreditorIBAN, transfer.DebtorIBAN, amount.StringFixed(2), "EUR", transfer.RemittanceInformation, time.Now().Format("2016-01-02"))
	ctx.Response.Header.Add("ResourceId", resId)

	return ctx.Created()
}
