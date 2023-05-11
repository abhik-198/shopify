package order

import (
	"context"
	"net/http"

	"github.com/sellerapp-com/shopifyapi/shopify/httpclient"
)

type Orders struct {
	client httpclient.HttpClient
	logger httpclient.Logger
}
type Order struct {
	Id                            int64
	AdminGraphqlApiId             string
	AppId                         int64
	BuyerAcceptsMarketing         bool
	Confirmed                     bool
	CreatedAt                     string
	Currency                      string
	CurrentSubtotalPrice          string
	CancelReason                  interface{}
	CancelledAt                   interface{}
	CartToken                     interface{}
	CheckoutId                    interface{}
	CheckoutToken                 interface{}
	ClientDetails                 interface{}
	ClosedAt                      interface{}
	Company                       interface{}
	ContactEmail                  interface{}
	CurrentSubtotalPriceSet       CurrentSubtotalPriceSet
	CurrentTotalAdditionalFeesSet interface{}
	CurrentTotalDiscounts         string
	CurrentTotalDiscountsSet      CurrentTotalDiscountsSet
	CurrentTotalDutiesSet         interface{}
	CurrentTotalPrice             string
	CurrentTotalPriceSet          CurrentTotalPriceSet
}
type CurrentSubtotalPriceSet struct {
	ShopMoney        ShopMoney
	PresentmentMoney PresentmentMoney
}
type CurrentTotalDiscountsSet struct {
	ShopMoney        ShopMoney
	PresentmentMoney PresentmentMoney
}
type CurrentTotalPriceSet struct {
	ShopMoney        ShopMoney
	PresentmentMoney PresentmentMoney
}
type ShopMoney struct {
	Amount       string
	CurrencyCode string
}
type PresentmentMoney struct {
	Amount       string
	CurrencyCode string
}

func NewOrders(client httpclient.HttpClient, logger httpclient.Logger) *Orders {
	return &Orders{
		client: client,
		logger: logger,
	}
}

func (p *Orders) List(ctx context.Context) ([]*Order, *http.Response, error) {
	req, err := http.NewRequest("GET", "/admin/products.json", nil)
	if err != nil {
		return nil, nil, err
	}

	var orderWrapper struct {
		Orders []*Order `json:"orders"`
	}
	resp, err := p.client.Do(req)
	if err != nil {
		return nil, resp, err
	}

	return orderWrapper.Orders, resp, nil
}
