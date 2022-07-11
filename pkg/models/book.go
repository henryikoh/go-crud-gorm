package models

import "gorm.io/gorm"

// basic structure of a book
type Book struct {
	gorm.Model
	ID       int     `json:"ID" gorm:"primaryKey; AUTO_INCREMENT"`
	Isbn     string  `json:"Isbn"`
	Title    string  `json:"title"`
	Director *Author `json:"author" gorm:"embedded"`
}

// Products is a collection of Product
type Books []*Book

// http respone error if book now found
