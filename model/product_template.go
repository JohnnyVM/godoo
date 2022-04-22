package model

/// Odoo model product_template
type ProductTemplate struct {
	table           string     `value:"product.template"`
	Id              int64      `json:"id"`
	Active          bool       `json:"active"`
	SaleOk          bool       `json:"sale_ok"`
	PurchaseOk      bool       `json:"purchase_ok"`
	Name            string     `json:"name"`
	CategId         int64      `json:"categ_id"`
	Barcode         OdooString `json:"barcode"`
	DefaultCode     OdooString `json:"default_code"`
	Type            string     `json:"type"` // posible values consu, service, product
	TaxesId         []int64    `json:"taxes_id"`
	SupplierTaxesId []int64    `json:"supplier_taxes_id"`
	PublicCategIds  []int64    `json:"public_categ_ids"`
	StandardPrice   float64    `json:"standad_price"` // cost
}

func (pt ProductTemplate) Fields() []string {
	tInterface := Fields(pt)
	out := make([]string, len(tInterface))
	for _, v := range tInterface {
		out = append(out, v.(string))
	}
	return out
}
