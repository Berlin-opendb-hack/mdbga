//************************************************************************//
// API "mdbga": Application Contexts
//
// Generated with goagen v1.0.0, command line:
// $ goagen
// --design=github.com/Berlin-opendb-hack/mdbga/design
// --out=$(GOPATH)/src/github.com/Berlin-opendb-hack/mdbga
// --version=v1.0.0
//
// The content of this file is auto-generated, DO NOT MODIFY
//************************************************************************//

package app

import (
	"github.com/goadesign/goa"
	"golang.org/x/net/context"
)

// GetBlockchainTransfersBlockchainContext provides the blockchain GetBlockchainTransfers action context.
type GetBlockchainTransfersBlockchainContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
}

// NewGetBlockchainTransfersBlockchainContext parses the incoming request URL and body, performs validations and creates the
// context used by the blockchain controller GetBlockchainTransfers action.
func NewGetBlockchainTransfersBlockchainContext(ctx context.Context, service *goa.Service) (*GetBlockchainTransfersBlockchainContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	rctx := GetBlockchainTransfersBlockchainContext{Context: ctx, ResponseData: resp, RequestData: req}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *GetBlockchainTransfersBlockchainContext) OK(r OpendbHackTransferCollection) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.opendb.hack.transfer+json; type=collection")
	if r == nil {
		r = OpendbHackTransferCollection{}
	}
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// OKFull sends a HTTP response with status code 200.
func (ctx *GetBlockchainTransfersBlockchainContext) OKFull(r OpendbHackTransferFullCollection) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.opendb.hack.transfer+json; type=collection")
	if r == nil {
		r = OpendbHackTransferFullCollection{}
	}
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// Unauthorized sends a HTTP response with status code 401.
func (ctx *GetBlockchainTransfersBlockchainContext) Unauthorized() error {
	ctx.ResponseData.WriteHeader(401)
	return nil
}

// InternalServerError sends a HTTP response with status code 500.
func (ctx *GetBlockchainTransfersBlockchainContext) InternalServerError(r error) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.goa.error")
	return ctx.ResponseData.Service.Send(ctx.Context, 500, r)
}

// PostBlockchainTransferBlockchainContext provides the blockchain PostBlockchainTransfer action context.
type PostBlockchainTransferBlockchainContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	Payload *PostBlockchainTransferBlockchainPayload
}

// NewPostBlockchainTransferBlockchainContext parses the incoming request URL and body, performs validations and creates the
// context used by the blockchain controller PostBlockchainTransfer action.
func NewPostBlockchainTransferBlockchainContext(ctx context.Context, service *goa.Service) (*PostBlockchainTransferBlockchainContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	rctx := PostBlockchainTransferBlockchainContext{Context: ctx, ResponseData: resp, RequestData: req}
	return &rctx, err
}

// postBlockchainTransferBlockchainPayload is the blockchain PostBlockchainTransfer action payload.
type postBlockchainTransferBlockchainPayload struct {
	// Blockchain address
	Address *string `form:"address,omitempty" json:"address,omitempty" xml:"address,omitempty"`
	// Amount of the transaction, EUR
	Amount *string `form:"amount,omitempty" json:"amount,omitempty" xml:"amount,omitempty"`
	// Date in RFC3339 format
	Date *string `form:"date,omitempty" json:"date,omitempty" xml:"date,omitempty"`
	// Exchange rate
	ExchangeRate *string `form:"exchangeRate,omitempty" json:"exchangeRate,omitempty" xml:"exchangeRate,omitempty"`
	// Paid fee for transfer
	Fee *string `form:"fee,omitempty" json:"fee,omitempty" xml:"fee,omitempty"`
	// Payment reference
	Identifier *string `form:"identifier,omitempty" json:"identifier,omitempty" xml:"identifier,omitempty"`
}

// Validate runs the validation rules defined in the design.
func (payload *postBlockchainTransferBlockchainPayload) Validate() (err error) {
	if payload.Amount == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "amount"))
	}
	if payload.Address == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "address"))
	}
	if payload.Identifier == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "identifier"))
	}
	if payload.ExchangeRate == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "exchangeRate"))
	}

	return
}

// Publicize creates PostBlockchainTransferBlockchainPayload from postBlockchainTransferBlockchainPayload
func (payload *postBlockchainTransferBlockchainPayload) Publicize() *PostBlockchainTransferBlockchainPayload {
	var pub PostBlockchainTransferBlockchainPayload
	if payload.Address != nil {
		pub.Address = *payload.Address
	}
	if payload.Amount != nil {
		pub.Amount = *payload.Amount
	}
	if payload.Date != nil {
		pub.Date = payload.Date
	}
	if payload.ExchangeRate != nil {
		pub.ExchangeRate = *payload.ExchangeRate
	}
	if payload.Fee != nil {
		pub.Fee = payload.Fee
	}
	if payload.Identifier != nil {
		pub.Identifier = *payload.Identifier
	}
	return &pub
}

// PostBlockchainTransferBlockchainPayload is the blockchain PostBlockchainTransfer action payload.
type PostBlockchainTransferBlockchainPayload struct {
	// Blockchain address
	Address string `form:"address" json:"address" xml:"address"`
	// Amount of the transaction, EUR
	Amount string `form:"amount" json:"amount" xml:"amount"`
	// Date in RFC3339 format
	Date *string `form:"date,omitempty" json:"date,omitempty" xml:"date,omitempty"`
	// Exchange rate
	ExchangeRate string `form:"exchangeRate" json:"exchangeRate" xml:"exchangeRate"`
	// Paid fee for transfer
	Fee *string `form:"fee,omitempty" json:"fee,omitempty" xml:"fee,omitempty"`
	// Payment reference
	Identifier string `form:"identifier" json:"identifier" xml:"identifier"`
}

// Validate runs the validation rules defined in the design.
func (payload *PostBlockchainTransferBlockchainPayload) Validate() (err error) {
	if payload.Amount == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "amount"))
	}
	if payload.Address == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "address"))
	}
	if payload.Identifier == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "identifier"))
	}
	if payload.ExchangeRate == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "exchangeRate"))
	}

	return
}

// Created sends a HTTP response with status code 201.
func (ctx *PostBlockchainTransferBlockchainContext) Created() error {
	ctx.ResponseData.WriteHeader(201)
	return nil
}

// BadRequest sends a HTTP response with status code 400.
func (ctx *PostBlockchainTransferBlockchainContext) BadRequest(r error) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.goa.error")
	return ctx.ResponseData.Service.Send(ctx.Context, 400, r)
}

// Unauthorized sends a HTTP response with status code 401.
func (ctx *PostBlockchainTransferBlockchainContext) Unauthorized() error {
	ctx.ResponseData.WriteHeader(401)
	return nil
}

// InternalServerError sends a HTTP response with status code 500.
func (ctx *PostBlockchainTransferBlockchainContext) InternalServerError(r error) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.goa.error")
	return ctx.ResponseData.Service.Send(ctx.Context, 500, r)
}
