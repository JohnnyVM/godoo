package model

/// Odoo model product_template
type ProductTemplate struct {
	table                   string     `value:"product.template"`
	Id                      int64      `json:"id"`
	Active                  bool       `json:"active"`
	SaleOk                  bool       `json:"sale_ok"`
	PurchaseOk              bool       `json:"purchase_ok"`
	Name                    string     `json:"name"`
	CategId                 Many2One   `json:"categ_id,omitempty"`
	Barcode                 OdooString `json:"barcode,omitempty"`
	DefaultCode             OdooString `json:"default_code,omitempty"`
	Type                    string     `json:"type,omitempty"` // posible values consu, service, product
	TaxesId                 []int64    `json:"taxes_id,omitempty"`
	SupplierTaxesId         []int64    `json:"supplier_taxes_id,omitempty"`
	PublicCategIds          []int64    `json:"public_categ_ids,omitempty"`
	StandardPrice           float64    `json:"standard_price"` // cost
	Image1920               Binary     `json:"image_1920,omitempty"`
	ProductTemplateImageIds []int64    `json:"product_template_image_ids,omitempty"`
}
