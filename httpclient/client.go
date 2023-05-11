package httpclient

import (
	"context"
	"net/http"
	"time"

	"github.com/sellerapp-com/shopifyapi/shopify/constants"
	"golang.org/x/oauth2"
	"golang.org/x/time/rate"
)

// AmazonClient will manage all credentials related to a particular profile
// and manage the tokens associated with the profile. It also manages creating
// new clients to make requests for the Amazon Advertising APIs so that
// parallell requests can be made to the API endpoints
type ShopifyClient struct {
	scheme   string
	profile  uint64
	clientId string
	client   http.Client
	source   oauth2.TokenSource
	host     string
	limiter  *rate.Limiter
}

func NewShopifyClient(client http.Client, profile uint64, refreshToken string, accessToken string, limiter *rate.Limiter) *ShopifyClient {
	// cfg := oauth2.Config{
	// 	ClientID:     constants.CLIENT_ID,
	// 	ClientSecret: constants.CLIENT_SECRET,
	// 	//	Endpoint:     geo.Endpoint(),
	// 	Scopes: []string{constants.FormatID(profile)},
	// }
	cfg := &oauth2.Config{
		ClientID:     constants.CLIENT_ID,
		ClientSecret: constants.CLIENT_SECRET,
		Scopes:       []string{"read_products", "write_products", "write_orders", "read_orders"},
		Endpoint: oauth2.Endpoint{
			AuthURL:  "https://ondc-sellerapp.myshopify.com/admin/oauth/authorize",
			TokenURL: "https://ondc-sellerapp.myshopify.com/admin/oauth/access_token",
		},
		RedirectURL: constants.REDIRECT_URL,
	}
	token := oauth2.Token{
		RefreshToken: refreshToken,
		AccessToken:  accessToken,
		TokenType:    constants.CLIENT_SECRET,
		Expiry:       time.Now(),
	}

	src := cfg.TokenSource(context.Background(), &token)
	// 2 hour timeout for client so that we do not have long
	// running clients which are waiting around a lot
	// also ignoring cancel function since we do not want to cancel the
	// context until it times out
	// ctx, _ := context.WithTimeout(context.Background(), 7200*time.Second)
	// client := oauth2.NewClient(ctx, src)
	// client := &http.Client{}
	return &ShopifyClient{
		scheme:   constants.SECURE_HTTP,
		profile:  profile,
		client:   client,
		source:   src,
		clientId: constants.CLIENT_ID,
		host:     constants.Hostname,
		limiter:  limiter,
	}
}

func (a *ShopifyClient) Scheme() string {
	return a.scheme
}

func (a *ShopifyClient) Host() string {
	return a.host
}

func (a *ShopifyClient) Do(req *http.Request) (*http.Response, error) {
	// NOTE: This method is not preferred and we would rather derive a client
	// from the oauth2.Config or oauth2.NewClient to get a client with built in
	// header adding capability, but there seems to be a problem with downloading
	// reports from S3 when using that type of client which is why we have decided
	// to go with this approach for now until we investigate and fix the problem.
	// refer to gitlab/advertising issue #1 for details

	req.Header.Add("X-Shopify-Access-Token", constants.CLIENT_SECRET)
	// ctx := context.Background()
	// err = a.limiter.Wait(ctx) // This is a blocking call.
	// if err != nil {
	// 	return nil, err
	// }
	return a.client.Do(req)
}
