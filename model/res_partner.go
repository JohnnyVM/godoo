package model

/// Odoo model product_template
type ResPartner struct {
	table  string `value:"res.partner"`
	Id     int64  `json:"id"`
	Active bool   `json:"active"`
	Name   string `json:"name"`
}
