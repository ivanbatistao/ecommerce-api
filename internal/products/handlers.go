package products

import (
	"log"
	"net/http"

	"github.com/ivanbatistao/ecommerce-api/internal/json"
)

type Handler struct {
	service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) ListProducts(w http.ResponseWriter, r *http.Request) {
	error := h.service.ListProducts(r.Context()) // products result is missing from ListProducts

	if error != nil {
		log.Println(error)
		http.Error(w, error.Error(), http.StatusInternalServerError)
		return
	}

	products := struct {
		Products []string `json:"products"`
	}{}

	json.Write(w, http.StatusOK, products)
}
