//go:generate goagen bootstrap -d github.com/Berlin-opendb-hack/mdbga/design

package main

import (
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware"
	"github.com/Berlin-opendb-hack/mdbga/app"
	"golang.org/x/net/context"
	"net/http"
)

const (
	bankScheme = "BANK_SCHEME"
	bankHost = "BANK_HOST"
	bankPath = "BANK_PATH"
	masterAccountIban = "MASTER_ACCOUNT_IBAN"
	masterAccountBic = "MASTER_ACCOUNT_BIC"
	masterAccountHolder = "MASTER_ACCOUNT_HOLDER"
	defalultBic = "DEFALULT_BIC"
	blockchainScheme = "BLOCKCHAIN_SCHEME"
	blockchainHost = "BLOCKCHAIN_HOST"
	blockchainPath = "BLOCKCHAIN_PATH"
)


func main() {
	// Create service
	service := goa.New("mdbga")

	// Mount middleware
	service.Use(middleware.RequestID())
	service.Use(middleware.LogRequest(true))
	service.Use(middleware.ErrorHandler(service, true))
	service.Use(middleware.Recover())
	app.UseTokenMiddleware(service, AuthMiddleWare)


	// Mount "blockchain" controller
	c := NewBlockchainController(service)
	app.MountBlockchainController(service, c)

	// Start service
	if err := service.ListenAndServe(":8881"); err != nil {
		service.LogError("startup", "err", err)
	}
}

func AuthMiddleWare(h goa.Handler) goa.Handler {
	return func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		token := req.Header.Get("Authorisation")
		if "" == token {
			return goa.ErrUnauthorized("Authorisation header required")
		}
		newctx := context.WithValue(ctx, "token", token)

		return h(newctx, rw, req)

	}
}