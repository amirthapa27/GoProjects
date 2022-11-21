package entity

type Post struct {
	ID    int64  `json:"id" gorm:"primaryKey"`
	Title string `json:"title"`
	Text  string `json:"text"`
}
