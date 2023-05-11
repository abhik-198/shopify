package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type Product struct {
	ProductID      int64       `json:"product_id"`
	Title          string      `json:"title"`
	BodyHTML       string      `json:"body_html"`
	Vendor         string      `json:"vendor"`
	ProductType    string      `json:"product_type"`
	Handle         string      `json:"handle"`
	TemplateSuffix interface{} `json:"template_suffix"`
	PublishedScope string      `json:"published_scope"`
	Tags           string      `json:"tags"`
	Available      bool        `json:"available"`

	Variants []*ProductVariant `json:"variants"`
	Options  []*ProductOption  `json:"options"`
	Images   []*ProductImage   `json:"images"`
	Image    *ProductImage     `json:"image"`

	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	PublishedAt time.Time `json:"published_at"`
}

type ProductVariant struct {
	ID                   int64           `json:"id"`
	ProductID            int64           `json:"product_id"`
	Title                string          `json:"title"`
	Price                string          `json:"price"`
	Sku                  string          `json:"sku"`
	Position             int             `json:"position"`
	Grams                int             `json:"grams"`
	InventoryPolicy      string          `json:"inventory_policy"`
	FulfillmentService   string          `json:"fulfillment_service"`
	InventoryManagement  string          `json:"inventory_management"`
	Option1              string          `json:"option1"`
	Option2              string          `json:"option2"`
	Option3              string          `json:"option3"`
	OptionValues         []VariantOption `json:"option_values"`
	Taxable              bool            `json:"taxable"`
	Barcode              string          `json:"barcode"`
	ImageID              interface{}     `json:"image_id"`
	CompareAtPrice       string          `json:"compare_at_price"`
	Available            bool            `json:"available"`
	InventoryQuantity    int             `json:"inventory_quantity"`
	Weight               float64         `json:"weight"`
	WeightUnit           string          `json:"weight_unit"`
	OldInventoryQuantity int             `json:"old_inventory_quantity"`
	RequiresShipping     bool            `json:"requires_shipping"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type VariantOption struct {
	OptionID int64  `json:"option_id"`
	Name     string `json:"name"`
	Value    string `json:"value"`
}

type ProductOption struct {
	ID        int64    `json:"id"`
	ProductID int64    `json:"product_id"`
	Name      string   `json:"name"`
	Position  int      `json:"position"`
	Values    []string `json:"values"`
}

type ProductImage struct {
	ID         int64   `json:"id"`
	ProductID  int64   `json:"product_id"`
	Position   int     `json:"position"`
	CreatedAt  string  `json:"created_at"`
	UpdatedAt  string  `json:"updated_at"`
	Width      int64   `json:"width"`
	Height     int64   `json:"height"`
	Src        string  `json:"src"`
	VariantIds []int64 `json:"variant_ids"`
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/users", getUsers).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", r))
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	req, err := http.NewRequest("GET", "https://quick-start-115efcbb.myshopify.com/admin/api/2023-04/products.json", nil)
	if err != nil {
		return
	}
	req.Header.Add("X-Shopify-Access-Token", "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")
	var productWrapper struct {
		Products []*Product `json:"products"`
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return
	}
	fmt.Print(string(body))
	json.Unmarshal(body, &productWrapper.Products)
	fmt.Println(productWrapper.Products)
	return
}
