package Models

import "gorm.io/gorm"

type URL struct {

	gorm.Model
	OriginalURL string `json:"original_url`
	ShortCode string `json:"shortcode`
}