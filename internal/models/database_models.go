package models

type Zone struct {
	ID     int `db:"id"`
	Width  int `db:"width"`
	Height int `db:"height"`
}

type Banner struct {
	ID               int    `db:"id"`
	Date             string `db:"date"`
	BackgroundColor  string `db:"background_color"`
	BackgroundImage  string `db:"background_image"`
	ButtonText       string `db:"button_text"`
	ButtonColor      string `db:"button_color"`
	ButtonBackground string `db:"button_background"`
	Title            string `db:"title"`
	Description      string `db:"description"`
	TextAlign        string `db:"text_align"`
	Link             string `db:"link"`
	LinkIsExternal   bool   `db:"link_is_external"`
}
