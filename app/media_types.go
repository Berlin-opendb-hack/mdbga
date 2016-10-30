//************************************************************************//
// API "mdbga": Application Media Types
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

import "github.com/goadesign/goa"

// OpendbHackTransfer media type (default view)
//
// Identifier: application/vnd.opendb.hack.transfer+json; view=default
type OpendbHackTransfer struct {
	// Amount of the transaction, EUR
	Amount string `form:"amount" json:"amount" xml:"amount"`
	// Date in RFC3339 format
	Date *string `form:"date,omitempty" json:"date,omitempty" xml:"date,omitempty"`
	// Exchange rate
	ExchangeRate string `form:"exchangeRate" json:"exchangeRate" xml:"exchangeRate"`
	// Payment reference
	Identifier string `form:"identifier" json:"identifier" xml:"identifier"`
}

// Validate validates the OpendbHackTransfer media type instance.
func (mt *OpendbHackTransfer) Validate() (err error) {
	if mt.Amount == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "amount"))
	}
	if mt.Identifier == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "identifier"))
	}
	if mt.ExchangeRate == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "exchangeRate"))
	}

	return
}

// OpendbHackTransfer media type (full view)
//
// Identifier: application/vnd.opendb.hack.transfer+json; view=full
type OpendbHackTransferFull struct {
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

// Validate validates the OpendbHackTransferFull media type instance.
func (mt *OpendbHackTransferFull) Validate() (err error) {
	if mt.Amount == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "amount"))
	}
	if mt.Address == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "address"))
	}
	if mt.Identifier == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "identifier"))
	}
	if mt.ExchangeRate == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "exchangeRate"))
	}

	return
}

// OpendbHackTransferCollection is the media type for an array of OpendbHackTransfer (default view)
//
// Identifier: application/vnd.opendb.hack.transfer+json; type=collection; view=default
type OpendbHackTransferCollection []*OpendbHackTransfer

// Validate validates the OpendbHackTransferCollection media type instance.
func (mt OpendbHackTransferCollection) Validate() (err error) {
	for _, e := range mt {
		if e.Amount == "" {
			err = goa.MergeErrors(err, goa.MissingAttributeError(`response[*]`, "amount"))
		}
		if e.Identifier == "" {
			err = goa.MergeErrors(err, goa.MissingAttributeError(`response[*]`, "identifier"))
		}
		if e.ExchangeRate == "" {
			err = goa.MergeErrors(err, goa.MissingAttributeError(`response[*]`, "exchangeRate"))
		}

	}
	return
}

// OpendbHackTransferCollection is the media type for an array of OpendbHackTransfer (full view)
//
// Identifier: application/vnd.opendb.hack.transfer+json; type=collection; view=full
type OpendbHackTransferFullCollection []*OpendbHackTransferFull

// Validate validates the OpendbHackTransferFullCollection media type instance.
func (mt OpendbHackTransferFullCollection) Validate() (err error) {
	for _, e := range mt {
		if e.Amount == "" {
			err = goa.MergeErrors(err, goa.MissingAttributeError(`response[*]`, "amount"))
		}
		if e.Address == "" {
			err = goa.MergeErrors(err, goa.MissingAttributeError(`response[*]`, "address"))
		}
		if e.Identifier == "" {
			err = goa.MergeErrors(err, goa.MissingAttributeError(`response[*]`, "identifier"))
		}
		if e.ExchangeRate == "" {
			err = goa.MergeErrors(err, goa.MissingAttributeError(`response[*]`, "exchangeRate"))
		}

	}
	return
}
