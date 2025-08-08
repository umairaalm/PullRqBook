package models

type Book struct {
	ID     uint   `json:"id" gorm:"primaryKey"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Year   int    `json:"year"`
}

type AddNewBook struct {
	Title  string `json:"title"`
	Author string `json:"author"`
	Year   int    `json:"year"`
}
