package handlers

import (
	"log"
	"net/http"
	"github.com/filipebafica/rest_golang/src/data"
)

type Products struct {
	l	*log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) getProducts(rw http.ResponseWriter, r *http.Request) {
	productList := data.GetProducts()
	err := productList.ToJson(rw)
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}
}

func (p *Products) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	// GET Method
	if r.Method == http.MethodGet {
		p.getProducts(rw, r)
		return
	}
	// catch all
	rw.WriteHeader(http.StatusMethodNotAllowed)
}
