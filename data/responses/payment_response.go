package responses

import "time"

type PaymentResponse struct {
	Id        string    `json:"id"`
	PayBy     string    `json:"pay_by"`
	PayDate   time.Time `json:"pay_date"`
	Amount    int64     `json:"amount"`
	TxnRef    string    `json:"txn_ref"`
	OrderInfo string    `json:"order_info"`
}

type PaymentPageResponse = BasePaginatedResponse[PaymentResponse]
