package handlers

import (
	"log"
	"net/http"
	"github.com/filipebafica/src/data"
	"encoding/json"
)

type Products struct {
	l	*log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	productList := data.GetProducts()
	data, err != json.Marshall(productList)
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}

	rw.Writer(data)
}
