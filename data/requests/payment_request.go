package requests

type PaymentRequest struct {
	Amount      int64  `json:"amount" validate:"min=1000"`
	Description string `json:"description"`
}

type PaymentResultRequest struct {
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
