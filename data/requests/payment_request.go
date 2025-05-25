package requests

import "time"

type PaymentFilters struct {
	BaseRequest
	PayBy   string     `query:"pay_by" json:"pay_by" validate:"omitempty,oneof=momo vnpay"`
	PayFrom *time.Time `query:"pay_from" json:"pay_from" validate:"omitempty"`
	PayTo   *time.Time `query:"pay_to" json:"pay_to" validate:"omitempty"`
}

type PaymentRequest struct {
	Amount      int64  `json:"amount" validate:"min=10000"`
	Description string `json:"description" default:"Payment for order"`
}

type VNPayPaymentResultRequest struct {
	Amount            int64  `query:"vnp_Amount"`
	BankCode          string `query:"vnp_BankCode"`
	BankTransactionNo string `query:"vnp_BankTranNo"`
	CardType          string `query:"vnp_CardType"`
	OrderInfo         string `query:"vnp_OrderInfo"`
	PayDate           string `query:"vnp_PayDate"`
	ResponseCode      string `query:"vnp_ResponseCode"`
	TmnCode           string `query:"vnp_TmnCode"`
	TransactionNo     string `query:"vnp_TransactionNo"`
	TransactionStatus string `query:"vnp_TransactionStatus"`
	TxnRef            string `query:"vnp_TxnRef"`
}

type MomoPaymentResultRequest struct {
	PartnerCode       string `query:"partnerCode"`
	OrderId           string `query:"orderId"`
	Amount            int64  `query:"amount"`
	OrderInfo         string `query:"orderInfo"`
	OrderType         string `query:"orderType"`
	TransactionNo     string `query:"transId"`
	ResponseCode      string `query:"resultCode"`
	TransactionStatus string `query:"message"`
	Type              string `query:"payType"`
	PayDate           string `query:"responseTime"`
	// Signature string `query:"signature"`
	// BankCode          string `query:"bankCode"`
	// BankTransactionNo string `query:"bankTransId"`
}
