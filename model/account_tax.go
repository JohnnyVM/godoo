package model

/// Odoo model account_tax
type AccountTax struct {
	table      string  `value:"account.tax"`
	Id         int64   `json:"id"`
	Active     bool    `json:"active"`
	Name       string  `json:"name"`
	TypeTaxUse string  `json:"type_tax_use"` // posible values sale, purchase, none
	Amount     float64 `json:"amount"`       // posible values sale, purchase, none
}
