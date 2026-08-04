package products

import (
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
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
	products, error := h.service.ListProducts(r.Context())

	if error != nil {
		log.Println(error)
		http.Error(w, error.Error(), http.StatusInternalServerError)
		return
	}

	json.Write(w, http.StatusOK, products)
}

func (h *Handler) FindProductById(w http.ResponseWriter, r *http.Request) {
	productIDStr := chi.URLParam(r, "id")
	productID, err := strconv.ParseInt(productIDStr, 10, 64)

	if err != nil {
		http.Error(w, "invalid product id", http.StatusBadRequest)
		return
	}

	product, err := h.service.FindProductById(r.Context(), productID)

	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.Write(w, http.StatusOK, product)
}
