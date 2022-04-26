package data

import (
	"time"
	"encoding/json"
	"io"
)

// Product defines the struct for an API product
type Product struct {
	ID			int		`json:"id"`
	Name		string	`json:"name"`
	Description	string	`json:"description"`
	Price		float32	`json:"price"`
	SKU			string	`json:"sku"`
	CreatedOn	string	`json:"-"`
	UpdatedOn	string	`json:"-"`
	DeletedOn	string	`json:"-"`
}

type Products []*Product

func (p *Products) ToJson(w io.Writer) error {
	encoder := json.NewEncoder(w)
	return encoder.Encode(p)
}

func GetProducts() Products {
	return productList
}

var productList = []*Product {
	&Product {
		ID:				1,
		Name:			"Latte",
		Description:	"Frothy milky coffee",
		Price:			2.45,
		SKU:			"abc323",
		CreatedOn:		time.Now().UTC().String(),
		UpdatedOn:		time.Now().UTC().String(),
	},
	&Product {
		ID:				2,
		Name:			"Espresso",
		Description:	"Short and strong coffee without milk",
		Price:			1.99,
		SKU:			"fjd34",
		CreatedOn:		time.Now().UTC().String(),
		UpdatedOn:		time.Now().UTC().String(),
	},
}