package model

/// Odoo model product_template
type ProductTemplate struct {
	table       string `value:"product.template"`
	Id          int64  `json:"id"`
	Name        string `json:"name"`
	Barcode     string `json:"barcode"`
	DefaultCode string `json:"default_code"`
}

func (pt ProductTemplate) TableName() string {
	return TableName(pt)
}

func (pt ProductTemplate) Fields() []string {
	tInterface := Fields(pt)
	out := make([]string, len(tInterface))
	for _, v := range tInterface {
		out = append(out, v.(string))
	}
	return out
}
