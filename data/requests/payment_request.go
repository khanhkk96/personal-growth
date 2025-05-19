package requests

type PaymentRequest struct {
	Amount      int64  `json:"amount" validate:"min=1000"`
	Description string `json:"description"`
}
