package model

/// Odoo model product_template
type ProductPublicCategory struct {
	table    string `value:"product.public.category"`
	Id       int64  `json:"id"`
	Name     string `json:"name"`
	ParentId int64  `json:"parent_id"`
	Sequence int64  `json:"sequence"`
}

func (pt ProductPublicCategory) Fields() []string {
	tInterface := Fields(pt)
	out := make([]string, len(tInterface))
	for _, v := range tInterface {
		out = append(out, v.(string))
	}
	return out
}
