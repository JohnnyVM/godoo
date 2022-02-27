package model

/// Odoo model product_template
type ResPartner struct {
	table string `value:"res.partner"`
	Id    int64  `json:"id"`
	Name  string `json:"name"`
}

func (rp ResPartner) TableName() string {
	return TableName(rp)
}

func (rp ResPartner) Fields() []string {
	tInterface := Fields(rp)
	out := make([]string, len(tInterface))
	for _, v := range tInterface {
		out = append(out, v.(string))
	}
	return out
}
