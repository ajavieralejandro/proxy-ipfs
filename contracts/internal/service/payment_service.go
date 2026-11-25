// archivo: internal/service/payment_service.go
package service

import (
	"errors"

	"my-go-proxy/internal/model"
	"my-go-proxy/internal/store"
)

var EmptyId = errors.New("Empty id")
var EmptyAmountWei = errors.New("Empty Amount Wei")


// PaymentResponse expone los campos de Payment más el estado calculado.
type PaymentResponse struct {
	model.Payment
	Status string `json:"status"`
}

// PaymentController / PaymentService coordina operaciones sobre pagos.
type PaymentService struct {
	store *store.PaymentStore
}

// NewPaymentService crea una instancia del servicio con su store.
func NewPaymentService(s *store.PaymentStore) *PaymentService {
	// TODO 1: Devuelve el puntero a PaymentService guardando el store recibido.
	toR := &PaymentService{
		store: s,
	}
	return toR
}

// toResponse convierte un *model.Payment en PaymentResponse incluyendo Status().
func toResponse(p *model.Payment) PaymentResponse {
	return PaymentResponse{
		Payment: *p,
		Status:  p.Status(),
	}
}

// List devuelve todos los pagos con su estado calculado.
func (s *PaymentService) List() []PaymentResponse {
	items := s.store.List()
	out := make([]PaymentResponse, 0, len(items))

	for _, payment := range items {
		item := payment
		response := toResponse(&item)
		out = append(out, response)
	}

	return out
}	

// GetByID busca un pago por ID y devuelve su respuesta con Status().
func (s *PaymentService) GetByID(id string) (PaymentResponse, error) {

	response,error := s.store.GetByID(id)
	if(error!=nil)
		return PaymentResponse{}, error
	return toResponse(response),nil
	
}

// Create valida y registra un nuevo pago en el store.
func (s *PaymentService) Create(p model.Payment) (PaymentResponse, error) {
	// TODO 5:
	// 1) Valida que p.ID != "" y p.AmountWei != "" (si falta alguno, devuelve errors.New(...)).
	// 2) Llama a s.store.Create(p). Si da error, devuélvelo.
	// 3) Devuelve toResponse(created), nil.

	if(p.ID == "" ){
		return PaymentResponse{}, errors.EmptyId
	}
	if(p.AmountWei ==""){
			return PaymentResponse{}, errors.EmptyAmountWei

	}
	return toResponse(&p)
	

}