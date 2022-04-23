package model

/// Odoo model product_image
type ProductImage struct {
	table     string     `value:"product.image"`
	Id        int64      `json:"id"`
	Name      string     `json:"name"`
	VideoUrl  OdooString `json:"video_url,omitempty"`
	Image1920 Binary     `json:"image_1920,omitempty"`
}
