package shopify

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/sellerapp-com/shopifyapi/shopify/httpclient"
	"github.com/sellerapp-com/shopifyapi/shopify/order"
	"github.com/sellerapp-com/shopifyapi/shopify/product"
	"golang.org/x/time/rate"
)

type Client struct {
	Products *product.Products
	Orders   *order.Orders
}

func NewClient(httpClient http.Client, params httpclient.ClientParams, opts httpclient.ClientOpts, limiter *rate.Limiter) (*Client, error) {
	//Secret Manager Logic goes here
	sClient := opts.SecretManager()
	ctx := context.Background()
	key := fmt.Sprintf("ADS_%s", params.AccountId())
	val, err := sClient.GetSecret(ctx, key)
	if err != nil {
		return nil, err
	}
	var token TokenRefreshResponse
	err = json.Unmarshal([]byte(val), &token)
	if err != nil {
		return nil, err
	}

	//ends here
	client := httpclient.NewShopifyClient(httpClient, params.ProfileId(), token.RefreshToken, token.AccessToken, limiter)
	products := product.NewProducts(client, opts.Logger())
	orders := order.NewOrders(client, opts.Logger())

	return &Client{
		Products: products,
		Orders:   orders,
	}, nil
}
