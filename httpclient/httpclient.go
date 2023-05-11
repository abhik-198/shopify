package httpclient

import (
	"context"
	"net/http"
)

type HttpClient interface {
	Scheme() string
	Host() string
	Do(*http.Request) (*http.Response, error)
}

type ClientParams interface {
	ProfileId() uint64
	AccessToken() string
	RefreshToken() string
	AccountId() string
}

type ClientOpts interface {
	Logger() Logger
	SecretManager() SecretManager
}

type Logger interface {
	Error(msg string)
	Debug(msg string)
	Info(msg string)
}

type SecretManager interface {
	GetSecret(ctx context.Context, key string) (val string, err error)
}
