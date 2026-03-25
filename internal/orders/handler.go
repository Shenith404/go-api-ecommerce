package orders

import (
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

func (h *handler) PlaceOrder(w http.ResponseWriter, r *http.Request) {
	json.Write(w,http.StatusCreated,nil)
}