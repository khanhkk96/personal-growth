package requests

type BaseRequest struct {
	Query *string `json:"q"`
	Page  int     `json:"page" default:"1"`
	Limit int     `json:"limit" default:"10"`
	Order string  `json:"order" default:"created_at"`
}
