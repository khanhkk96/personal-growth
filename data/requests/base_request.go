package requests

type BaseRequest struct {
	Query   string `query:"q" json:"q" validate:"omitempty"`
	Page    int    `query:"page" default:"1"`
	Limit   int    `query:"limit" default:"10"`
	Order   string `query:"order" validate:"oneof=ASC DESC" default:"DESC"`
	OrderBy string `query:"order_by" json:"order_by" default:"created_at"`
}
