package model

/// Odoo model product_template
type ProductPublicCategory struct {
	table    string `value:"product.public.category"`
	Id       int64  `json:"id"`
	Name     string `json:"name"`
	ParentId int64  `json:"parent_id,omitempty"`
	Sequence int64  `json:"sequence"`
}
