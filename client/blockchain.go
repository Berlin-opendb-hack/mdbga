package client

import (
	"bytes"
	"fmt"
	"golang.org/x/net/context"
	"net/http"
	"net/url"
)

// GetBlockchainTransfersBlockchainPath computes a request path to the GetBlockchainTransfers action of blockchain.
func GetBlockchainTransfersBlockchainPath() string {
	return fmt.Sprintf("/blockchain-transfers")
}

// GetBlockchainTransfersBlockchain makes a request to the GetBlockchainTransfers action endpoint of the blockchain resource
func (c *Client) GetBlockchainTransfersBlockchain(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewGetBlockchainTransfersBlockchainRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewGetBlockchainTransfersBlockchainRequest create the request corresponding to the GetBlockchainTransfers action endpoint of the blockchain resource.
func (c *Client) NewGetBlockchainTransfersBlockchainRequest(ctx context.Context, path string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	if c.TokenSigner != nil {
		c.TokenSigner.Sign(req)
	}
	return req, nil
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
	// Payment reference
	Identifier string `form:"identifier" json:"identifier" xml:"identifier"`
}

// PostBlockchainTransferBlockchainPath computes a request path to the PostBlockchainTransfer action of blockchain.
func PostBlockchainTransferBlockchainPath() string {
	return fmt.Sprintf("/blockchain-transfers")
}

// PostBlockchainTransferBlockchain makes a request to the PostBlockchainTransfer action endpoint of the blockchain resource
func (c *Client) PostBlockchainTransferBlockchain(ctx context.Context, path string, payload *PostBlockchainTransferBlockchainPayload, contentType string) (*http.Response, error) {
	req, err := c.NewPostBlockchainTransferBlockchainRequest(ctx, path, payload, contentType)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewPostBlockchainTransferBlockchainRequest create the request corresponding to the PostBlockchainTransfer action endpoint of the blockchain resource.
func (c *Client) NewPostBlockchainTransferBlockchainRequest(ctx context.Context, path string, payload *PostBlockchainTransferBlockchainPayload, contentType string) (*http.Request, error) {
	var body bytes.Buffer
	if contentType == "" {
		contentType = "*/*" // Use default encoder
	}
	err := c.Encoder.Encode(payload, &body, contentType)
	if err != nil {
		return nil, fmt.Errorf("failed to encode body: %s", err)
	}
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("POST", u.String(), &body)
	if err != nil {
		return nil, err
	}
	header := req.Header
	if contentType != "*/*" {
		header.Set("Content-Type", contentType)
	}
	if c.TokenSigner != nil {
		c.TokenSigner.Sign(req)
	}
	return req, nil
}
