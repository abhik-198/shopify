package shopify

type ShopifyError interface {
	ErrorCode() string
	ErrorDescription() string
}

type CommonError struct {
	Code        string `json:"code"`
	Description string `json:"details"`
}

func (c *CommonError) ErrorCode() string {
	return c.Code
}

func (c *CommonError) ErrorDescription() string {
	return c.Description
}
