package shopify

type TokenRefreshResponse struct {
	AccessToken  string `json:"access_token"`
	ExpiresIn    int64  `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	TokenType    string `json:"token_type"`
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

// {
//     "orders": [
//         {
//             "id": 5383941914937,
//             "admin_graphql_api_id": "gid://shopify/Order/5383941914937",
//             "app_id": 40550465537,
//             "browser_ip": null,
//             "buyer_accepts_marketing": false,
//             "cancel_reason": null,
//             "cancelled_at": null,
//             "cart_token": null,
//             "checkout_id": null,
//             "checkout_token": null,
//             "client_details": null,
//             "closed_at": null,
//             "company": null,
//             "confirmed": true,
//             "contact_email": null,
//             "created_at": "2023-05-05T08:17:19-04:00",
//             "currency": "EUR",
//             "current_subtotal_price": "224.97",
//             "current_subtotal_price_set": {
//                 "shop_money": {
//                     "amount": "224.97",
//                     "currency_code": "EUR"
//                 },
//                 "presentment_money": {
//                     "amount": "224.97",
//                     "currency_code": "EUR"
//                 }
//             },
//             "current_total_additional_fees_set": null,
//             "current_total_discounts": "0.00",
//             "current_total_discounts_set": {
//                 "shop_money": {
//                     "amount": "0.00",
//                     "currency_code": "EUR"
//                 },
//                 "presentment_money": {
//                     "amount": "0.00",
//                     "currency_code": "EUR"
//                 }
//             },
//             "current_total_duties_set": null,
//             "current_total_price": "238.47",
//             "current_total_price_set": {
//                 "shop_money": {
//                     "amount": "238.47",
//                     "currency_code": "EUR"
//                 },
//                 "presentment_money": {
//                     "amount": "238.47",
//                     "currency_code": "EUR"
//                 }
//             },
//             "current_total_tax": "13.50",
//             "current_total_tax_set": {
//                 "shop_money": {
//                     "amount": "13.50",
//                     "currency_code": "EUR"
//                 },
//                 "presentment_money": {
//                     "amount": "13.50",
//                     "currency_code": "EUR"
//                 }
//             },
//             "customer_locale": null,
//             "device_id": null,
//             "discount_codes": [],
//             "email": "",
//             "estimated_taxes": false,
//             "financial_status": "paid",
//             "fulfillment_status": null,
//             "landing_site": null,
//             "landing_site_ref": null,
//             "location_id": null,
//             "merchant_of_record_app_id": null,
//             "name": "#1001",
//             "note": null,
//             "note_attributes": [],
//             "number": 1,
//             "order_number": 1001,
//             "order_status_url": "https://ondc-sellerapp.myshopify.com/76190122297/orders/3d0603cc98452416bf1a329522479bb5/authenticate?key=b84d1e09092547379ab0d85c01711733",
//             "original_total_additional_fees_set": null,
//             "original_total_duties_set": null,
//             "payment_gateway_names": [
//                 ""
//             ],
//             "phone": null,
//             "presentment_currency": "EUR",
//             "processed_at": "2023-05-05T08:17:19-04:00",
//             "reference": null,
//             "referring_site": null,
//             "source_identifier": null,
//             "source_name": "40550465537",
//             "source_url": null,
//             "subtotal_price": "224.97",
//             "subtotal_price_set": {
//                 "shop_money": {
//                     "amount": "224.97",
//                     "currency_code": "EUR"
//                 },
//                 "presentment_money": {
//                     "amount": "224.97",
//                     "currency_code": "EUR"
//                 }
//             },
//             "tags": "",
//             "tax_lines": [
//                 {
//                     "price": "13.50",
//                     "rate": 0.06,
//                     "title": "State tax",
//                     "price_set": {
//                         "shop_money": {
//                             "amount": "13.50",
//                             "currency_code": "EUR"
//                         },
//                         "presentment_money": {
//                             "amount": "13.50",
//                             "currency_code": "EUR"
//                         }
//                     },
//                     "channel_liable": null
//                 }
//             ],
//             "taxes_included": false,
//             "test": false,
//             "token": "3d0603cc98452416bf1a329522479bb5",
//             "total_discounts": "0.00",
//             "total_discounts_set": {
//                 "shop_money": {
//                     "amount": "0.00",
//                     "currency_code": "EUR"
//                 },
//                 "presentment_money": {
//                     "amount": "0.00",
//                     "currency_code": "EUR"
//                 }
//             },
//             "total_line_items_price": "224.97",
//             "total_line_items_price_set": {
//                 "shop_money": {
//                     "amount": "224.97",
//                     "currency_code": "EUR"
//                 },
//                 "presentment_money": {
//                     "amount": "224.97",
//                     "currency_code": "EUR"
//                 }
//             },
//             "total_outstanding": "0.00",
//             "total_price": "238.47",
//             "total_price_set": {
//                 "shop_money": {
//                     "amount": "238.47",
//                     "currency_code": "EUR"
//                 },
//                 "presentment_money": {
//                     "amount": "238.47",
//                     "currency_code": "EUR"
//                 }
//             },
//             "total_shipping_price_set": {
//                 "shop_money": {
//                     "amount": "0.00",
//                     "currency_code": "EUR"
//                 },
//                 "presentment_money": {
//                     "amount": "0.00",
//                     "currency_code": "EUR"
//                 }
//             },
//             "total_tax": "13.50",
//             "total_tax_set": {
//                 "shop_money": {
//                     "amount": "13.50",
//                     "currency_code": "EUR"
//                 },
//                 "presentment_money": {
//                     "amount": "13.50",
//                     "currency_code": "EUR"
//                 }
//             },
//             "total_tip_received": "0.00",
//             "total_weight": 0,
//             "updated_at": "2023-05-05T08:17:20-04:00",
//             "user_id": null,
//             "billing_address": null,
//             "customer": null,
//             "discount_applications": [],
//             "fulfillments": [],
//             "line_items": [
//                 {
//                     "id": 13896730837305,
//                     "admin_graphql_api_id": "gid://shopify/LineItem/13896730837305",
//                     "fulfillable_quantity": 3,
//                     "fulfillment_service": "manual",
//                     "fulfillment_status": null,
//                     "gift_card": false,
//                     "grams": 1300,
//                     "name": "Big Brown Bear Boots",
//                     "price": "74.99",
//                     "price_set": {
//                         "shop_money": {
//                             "amount": "74.99",
//                             "currency_code": "EUR"
//                         },
//                         "presentment_money": {
//                             "amount": "74.99",
//                             "currency_code": "EUR"
//                         }
//                     },
//                     "product_exists": false,
//                     "product_id": null,
//                     "properties": [],
//                     "quantity": 3,
//                     "requires_shipping": true,
//                     "sku": null,
//                     "taxable": true,
//                     "title": "Big Brown Bear Boots",
//                     "total_discount": "0.00",
//                     "total_discount_set": {
//                         "shop_money": {
//                             "amount": "0.00",
//                             "currency_code": "EUR"
//                         },
//                         "presentment_money": {
//                             "amount": "0.00",
//                             "currency_code": "EUR"
//                         }
//                     },
//                     "variant_id": null,
//                     "variant_inventory_management": null,
//                     "variant_title": null,
//                     "vendor": null,
//                     "tax_lines": [
//                         {
//                             "channel_liable": null,
//                             "price": "13.50",
//                             "price_set": {
//                                 "shop_money": {
//                                     "amount": "13.50",
//                                     "currency_code": "EUR"
//                                 },
//                                 "presentment_money": {
//                                     "amount": "13.50",
//                                     "currency_code": "EUR"
//                                 }
//                             },
//                             "rate": 0.06,
//                             "title": "State tax"
//                         }
//                     ],
//                     "duties": [],
//                     "discount_allocations": []
//                 }
//             ],
//             "payment_terms": null,
//             "refunds": [],
//             "shipping_address": null,
//             "shipping_lines": []
//         }
//     ]
// }
