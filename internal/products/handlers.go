package products

import (
	"log"
	"net/http"

	"github.com/Shenith404/go-ecom/internal/json"
)

type handler struct {
	service Service
} 

func NewHandler(s Service) *handler {
	return &handler{
		service: s,
	}
}

func (h * handler) ListProducts(w http.ResponseWriter, r *http.Request) {
	//call the service
	err := h.service.ListProducts(r.Context())
	if err != nil {
		log.Printf("error listing products: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	products := []string{"Product 1", "Product 2", "Product 3"}

	json.Write(w,http.StatusOK,products)

}