package model

/// Odoo model product_template
type ProductCategory struct {
	table    string `value:"product.category"`
	Id       int64  `json:"id"`
	Name     string `json:"name"`
	ParentId int64  `json:"parent_id"`
}

func (pt ProductCategory) Fields() []string {
	tInterface := Fields(pt)
	out := make([]string, len(tInterface))
	for _, v := range tInterface {
		out = append(out, v.(string))
	}
	return out
}
