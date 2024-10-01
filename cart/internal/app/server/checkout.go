package server

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"gitlab.ozon.dev/1mikle1/homework/cart/internal/domain"
)

type CheckoutRequest struct {
	UserId domain.UID `json:"user" validate:"min=1"`
}

type CheckoutResponse struct {
	OrderId domain.OID `json:"orderID" validate:"min=1"`
}

const PathCheckout = "DELETE /cart/checkout"

func (s *CartServer) ExtractCheckoutRequest(r *http.Request) (checkoutRequest *CheckoutRequest, err error) {
	checkoutRequest = &CheckoutRequest{}
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body, checkoutRequest)
	if err != nil {
		return nil, err
	}

	// Валидация запроса
	_, err = s.validator.Validate(*checkoutRequest)
	if err != nil {
		return nil, err
	}

	return checkoutRequest, nil
}

// "DELETE /user/{user_id}/cart"
func (s *CartServer) Checkout(w http.ResponseWriter, r *http.Request) {

	checkoutRequest, err := s.ExtractCheckoutRequest(r)

	if err != nil {
		serverErr := ServerError{
			Status: http.StatusBadRequest,
			Text:   err.Error(),
			Path:   PathDeleteCart,
		}
		writeErrorResponse(w, &serverErr)
		return
	}

	err = s.SendCheckoutResponse(w, r, checkoutRequest)
	if err != nil {
		serverErr := ServerError{
			Path:   PathDeleteCart,
			Status: http.StatusInternalServerError,
			Text:   err.Error(),
		}
		writeErrorResponse(w, &serverErr)
		return
	}

}

func (s *CartServer) SendCheckoutResponse(w http.ResponseWriter, r *http.Request, checkoutRequest *CheckoutRequest) (err error) {
	id, err := s.cartService.Checkout(r.Context(), checkoutRequest.UserId)
	if err != nil {
		return err
	}

	result := &CheckoutResponse{OrderId: id}

	rawResponse, err := json.Marshal(result)
	if err != nil {
		return err
	}

	fmt.Fprint(w, string(rawResponse))
	return nil
}
