package router

import (
	"net/http"

	"my-go-proxy/internal/handler"
)

// New arma y devuelve el router HTTP con las rutas de pagos registradas.
func New(h *handler.PaymentHandler) http.Handler {
	mux := http.NewServeMux()

	// TODO 1: Registrar rutas sobre mux, por ejemplo:
	// mux.HandleFunc("GET /payments", h.List)
	// mux.HandleFunc("POST /payments", h.Create)
	// mux.HandleFunc("GET /payments/{id}", h.GetByID)

	// TODO 2 (opcional): Envolver mux con un middleware simple (ej. logging).
	return mux
}