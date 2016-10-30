package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = API("mdbga", func() {
	Scheme("http")
	Host("localhost:8881")
	APIKeySecurity("token", func() {
		Header("Authorisation")
	})
})


var _ = Resource("blockchain", func() {
	BasePath("/blockchain-transfers")
	Security("token", func() {})

	Action("PostBlockchainTransfer", func() {
		Routing(POST("/"))
		Response(Unauthorized, func() {})
		Response(BadRequest, func() {
			Media(ErrorMedia)
		})
		Response(InternalServerError, func() {
			Media(ErrorMedia)
		})

		Response(Created, func() {
			Headers(func() {
				Header("ResourceId")
			})
		})
		Payload(TransferMedia)
	})

	Action("GetBlockchainTransfers", func() {
		Routing(GET("/"))
		Response(Unauthorized, func() {})
		Response(OK, func() {
			Media(TransferMedia)
		})
	})
})

var TransferMedia = MediaType("application/vnd.opendb.hack.transfer+json", func() {
	Attributes(func() {
		Attribute("amount", String, "Amount of the transaction, EUR")
		Attribute("address", String, "Blockchain address")
		Attribute("identifier", String, "Payment reference")
		Attribute("exchangeRate", String, "Exchange rate")
		Attribute("date", String, "Date in RFC3339 format")
	})
	Required(
		"amount",
		"address",
		"identifier",
		"exchangeRate",
	)
	View("default", func() {
		Attribute("amount")
		Attribute("identifier")
		Attribute("date")
		Attribute("exchangeRate")

	})
	View("full", func() {
		Attribute("amount")
		Attribute("address")
		Attribute("identifier")
		Attribute("exchangeRate")
		Attribute("date")
	})
})