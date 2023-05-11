package constants

import "fmt"

const (
	SECURE_HTTP   = "https"
	CLIENT_ID     = "6e01146eec336da3bfb74ddd6620147b"
	CLIENT_SECRET = "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"
	REDIRECT_URL = "https://ondc-sellerapp.myshopify.com/oauth/callback"
)
func FormatID(id uint64) string {
	return fmt.Sprintf("%d", id)
}
const (
	Hostname ="quick-start-115efcbb.myshopify.com"
)