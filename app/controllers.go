//************************************************************************//
// API "mdbga": Application Controllers
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
	"net/http"
)

// initService sets up the service encoders, decoders and mux.
func initService(service *goa.Service) {
	// Setup encoders and decoders
	service.Encoder.Register(goa.NewJSONEncoder, "application/json")
	service.Encoder.Register(goa.NewGobEncoder, "application/gob", "application/x-gob")
	service.Encoder.Register(goa.NewXMLEncoder, "application/xml")
	service.Decoder.Register(goa.NewJSONDecoder, "application/json")
	service.Decoder.Register(goa.NewGobDecoder, "application/gob", "application/x-gob")
	service.Decoder.Register(goa.NewXMLDecoder, "application/xml")

	// Setup default encoder and decoder
	service.Encoder.Register(goa.NewJSONEncoder, "*/*")
	service.Decoder.Register(goa.NewJSONDecoder, "*/*")
}

// BlockchainController is the controller interface for the Blockchain actions.
type BlockchainController interface {
	goa.Muxer
	GetBlockchainTransfers(*GetBlockchainTransfersBlockchainContext) error
	PostBlockchainTransfer(*PostBlockchainTransferBlockchainContext) error
}

// MountBlockchainController "mounts" a Blockchain resource controller on the given service.
func MountBlockchainController(service *goa.Service, ctrl BlockchainController) {
	initService(service)
	var h goa.Handler

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewGetBlockchainTransfersBlockchainContext(ctx, service)
		if err != nil {
			return err
		}
		return ctrl.GetBlockchainTransfers(rctx)
	}
	h = handleSecurity("token", h)
	service.Mux.Handle("GET", "/blockchain-transfers", ctrl.MuxHandler("GetBlockchainTransfers", h, nil))
	service.LogInfo("mount", "ctrl", "Blockchain", "action", "GetBlockchainTransfers", "route", "GET /blockchain-transfers", "security", "token")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewPostBlockchainTransferBlockchainContext(ctx, service)
		if err != nil {
			return err
		}
		// Build the payload
		if rawPayload := goa.ContextRequest(ctx).Payload; rawPayload != nil {
			rctx.Payload = rawPayload.(*PostBlockchainTransferBlockchainPayload)
		} else {
			return goa.MissingPayloadError()
		}
		return ctrl.PostBlockchainTransfer(rctx)
	}
	h = handleSecurity("token", h)
	service.Mux.Handle("POST", "/blockchain-transfers", ctrl.MuxHandler("PostBlockchainTransfer", h, unmarshalPostBlockchainTransferBlockchainPayload))
	service.LogInfo("mount", "ctrl", "Blockchain", "action", "PostBlockchainTransfer", "route", "POST /blockchain-transfers", "security", "token")
}

// unmarshalPostBlockchainTransferBlockchainPayload unmarshals the request body into the context request data Payload field.
func unmarshalPostBlockchainTransferBlockchainPayload(ctx context.Context, service *goa.Service, req *http.Request) error {
	payload := &postBlockchainTransferBlockchainPayload{}
	if err := service.DecodeRequest(req, payload); err != nil {
		return err
	}
	if err := payload.Validate(); err != nil {
		// Initialize payload with private data structure so it can be logged
		goa.ContextRequest(ctx).Payload = payload
		return err
	}
	goa.ContextRequest(ctx).Payload = payload.Publicize()
	return nil
}
