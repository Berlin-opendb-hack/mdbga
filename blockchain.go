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
	// BlockchainController_GetBlockchainTransfers: start_implement

	// Put your logic here

	// BlockchainController_GetBlockchainTransfers: end_implement
	res := &app.OpendbHackTransfer{}
	return ctx.OK(res)
}

// PostBlockchainTransfer runs the PostBlockchainTransfer action.
func (c *BlockchainController) PostBlockchainTransfer(ctx *app.PostBlockchainTransferBlockchainContext) error {

	token := ctx.Value("token")
	if nil == token {
		return ctx.Unauthorized()
	}
	bankEndpoint := url.URL{
		Scheme: os.Getenv("BANK_SCHEME"),
		Host: os.Getenv("BANK_HOST"),
		Path: os.Getenv("BANK_PATH"),
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
	if -1 balance.Cmp(amount) {
		return ctx.BadRequest(errors.New("Amount not available in account"))
	}
	if nil != err {
		return ctx.InternalServerError(err)
	}
	transfer := bank.Transfer{
		Amount: amount,

	}
	resId, err := bankClient.CreateTransfer(token.(string), transfer)
	if nil != err {
		return ctx.InternalServerError(err)
	}
	ctx.Response.Header.Add("ResourceId", "asdf")
	return ctx.Created()
}
