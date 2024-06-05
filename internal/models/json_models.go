package models

type BannerJson struct {
	ID               int    `json:"id"`
	Date             string `json:"date"`
	BackgroundColor  string `json:"background_color"`
	BackgroundImage  string `json:"background_image"`
	ButtonText       string `json:"button_text"`
	ButtonColor      string `json:"button_color"`
	ButtonBackground string `json:"button_background"`
	Title            string `json:"title"`
	Description      string `json:"description"`
	TextAlign        string `json:"text_align"`
	Link             string `json:"link"`
	LinkIsExternal   bool   `json:"link_is_external"`
}

type ZoneBanner struct {
	ZoneID  int          `json:"zoneId"`
	Width   int          `json:"width"`
	Height  int          `json:"height"`
	Banners []BannerJson `json:"banners"`
}
