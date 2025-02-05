package models

type BookIssuance struct {
	ID         uint   `json:"id" gorm:"primaryKey"`
	BookTitle  string `json:"book_title"`
	IssuedTo   string `json:"issued_to"`
	IssuedDate string `json:"issued_date"`
	ReturnDate string `json:"return_date"`
}
