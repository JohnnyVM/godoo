package model

/// Odoo model product_template
type ProductCategory struct {
	table    string `value:"product.category"`
	Id       int64  `json:"id"`
	Name     string `json:"name"`
	ParentId int64  `json:"parent_id"`
}
