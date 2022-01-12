package models

type Book struct {
	ID     uint   `json:"id" gorm:"Primary_key"`
	Title  string `json:"title"`
	Author string `json:"author"`
}
