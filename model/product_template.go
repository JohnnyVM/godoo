package model

/// Odoo model product_template
type ProductTemplate struct {
	table                   string     `value:"product.template"`
	Id                      int64      `json:"id"`
	Active                  bool       `json:"active"`
	SaleOk                  bool       `json:"sale_ok"`
	PurchaseOk              bool       `json:"purchase_ok"`
	Name                    string     `json:"name"`
	CategId                 Many2One   `json:"categ_id"`
	Barcode                 OdooString `json:"barcode"`
	DefaultCode             OdooString `json:"default_code"`
	Type                    string     `json:"type"` // posible values consu, service, product
	TaxesId                 []int64    `json:"taxes_id"`
	SupplierTaxesId         []int64    `json:"supplier_taxes_id"`
	PublicCategIds          []int64    `json:"public_categ_ids"`
	StandardPrice           float64    `json:"standard_price"` // cost
	Image1920               Binary     `json:"image_1920"`
	ProductTemplateImageIds []int64    `json:"product_template_image_ids"`
}
