package data

import (
	"encoding/json"
	"fmt"
	"io"
	"time"
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

// FromJSON deserializes the JSON data to Product
func (p *Product) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(p)
}

// Products is a collenction of Product
type Products []*Product

// ToJSON serializes the contents of the collection to JSON
// NewEnconder provides better performance than json.Ummarshal as it does not
// have to buffer the output into an in memory slice of bytes
func (p *Products) ToJson(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}

// GetProducts returns a list of products
func GetProducts() Products {
	return productList
}

// AddProduct adds a product to db
func AddProduct(p *Product) {
	p.ID = getNextID()
	productList = append(productList, p)
}

func UpdateProduct(id int, p *Product) error {
	index, err := findProduct(id)
	if err != nil {
		return err
	}
	p.ID = id
	productList[index] = p
	return nil
}

var ErrProductNotFound = fmt.Errorf("Product not found")

func findProduct(id int) (int, error) {
	for i, p := range productList {
		if p.ID == id {
			return i, nil
		}
	}
	return -1, ErrProductNotFound
}

func getNextID() int {
	lp := productList[len(productList) -1]
	return lp.ID + 1
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
