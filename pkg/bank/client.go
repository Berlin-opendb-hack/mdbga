package bank

import (
	"net/url"
	"net/http"
	"fmt"
	"github.com/pkg/errors"
	"io/ioutil"
	"encoding/json"
)

type CashAccount struct {
	Iban  string `json:"iban"`
	Balance float64 `json:"balance"`
	ProductDescription  string `json:"productDescription"`
}

type Transfer struct {
	Amount string `json:"amount"`
	CreditorIBAN string `json:"creditorIBAN"`
	CreditorBIC string `json:"creditorBIC"`
	CreditorName string `json:"creditorName"`
	DebtorIBAN string `json:"debtorIBAN"`
	DebtorBIC string `json:"debtorBIC"`
	DebtorName string `json:"debtorName"`
	Currency string `json:"currency"`
	RemittanceInformation string `json:"remittanceInformation"`
}

type Requestor interface {
	Do(req *http.Request) (*http.Response, error)
}
type BankClient struct {
	httpClient Requestor
	endpoint   url.URL
}

func NewBankClient(client Requestor, endpoint url.URL) *BankClient {
	return &BankClient{
		httpClient:client,
		endpoint:endpoint,
	}
}

func (b *BankClient) CreateTransfer(token string, transfer Transfer) (string, error) {
	request, err := b.buildRequest(token, "POST", "/transactions")
	if nil != err {
		return "", err
	}
	res, err := b.httpClient.Do(request)
	if (nil != err) {
		return "", err
	}
	if (res.StatusCode == 201) {
		return res.Header.Get("ResourceId"), nil
	}
	return "", errors.Errorf("Unsupported status code %s", res.Status)
}

func (b *BankClient) GetAccounts(token string) ([]CashAccount, error) {
	request, err := b.buildRequest(token, "GET", "/cashAccounts")
	if nil != err {
		return []CashAccount{}, err
	}
	res, err := b.httpClient.Do(request)
	if nil != err {
		return []CashAccount{}, err
	}
	if (res.StatusCode == 200) {
		cashAccounts := []CashAccount{}
		responseBody, err := ioutil.ReadAll(res.Body)
		if nil != err {
			return []CashAccount{}, err
		}
		json.Unmarshal(responseBody, &cashAccounts)

		return cashAccounts, nil
	}
	return []CashAccount{}, errors.Errorf("Unsupported status code %s", res.Status)


}

func (b *BankClient) buildRequest(token string, method string, path string) (*http.Request, error) {
	endpoint := url.URL{
		Scheme: b.endpoint.Scheme,
		Host: b.endpoint.Host,
		Path: b.endpoint.Path + path,

	}
	request, err := http.NewRequest(method,endpoint.String(), nil)
	if nil != err {
		return nil, err
	}

	request.Header.Add("X-Token", fmt.Sprintf("%s", token))

	return request, nil

}