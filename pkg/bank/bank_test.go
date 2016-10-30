package bank_test

import (
	"testing"
	"net/url"
	"github.com/Berlin-opendb-hack/mdbga/pkg/bank"
	"fmt"
	"net/http/httptest"
	"github.com/golang/mock/gomock"
	"github.com/Berlin-opendb-hack/mdbga/pkg/bank/mock"
)

func TestBankClient_GetAccounts(t *testing.T) {
	endpoint := url.URL{
		Scheme: "http",
		Host: "localhost:8880",
		Path: "",
	}
	mockCtrl := gomock.NewController(t)

	client := mock_bank.NewMockRequestor(mockCtrl)
	response := httptest.NewRecorder()
	response.WriteString(`[{"iban":"DE10000000000000000306","balance":100.95,"productDescription":"persönliches Konto"}]`)
	response.Code = 200

	client.EXPECT().Do(gomock.Any()).Times(1).Return(response.Result(), nil)

	bC := bank.NewBankClient(client, endpoint)
	accounts, err := bC.GetAccounts("asdf")
	if nil != err {
		t.Error(err.Error())
	}
	if (1 != len(accounts)) {
		t.Fail()
		fmt.Printf("%+v\n", accounts)
	}
}
