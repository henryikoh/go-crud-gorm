// This is where all things concercing data would be placed
package repositery

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/henryikoh/book-management-go/pkg/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// interface for calling the data. You can add or remove as much interfaces as you like
type MovieQuery interface {
	GetBooks() models.Books
	GetBookByID(id string) (models.Book, error)
	CreateBook(book *models.Book)
	UpdateBook(book models.Book, id string) models.Book
	DeleteBook(id string) models.Book
}

type movieQuery struct {
	userId  string
	db      *gorm.DB
	context context.Context
}
type message struct {
	message string
	status  int64
}
type Errormessage struct {
	message string
	status  int64
}

func (error *Errormessage) Error() string {
	return error.message
}

// there should be different types of query
func NewMovieQuery(db *gorm.DB) MovieQuery {
	return &movieQuery{
		db: db,
	}
}

func (m *movieQuery) GetBooks() models.Books {
	var books models.Books
	result := m.db.Find(&books)
	if result.Error != nil {
		panic("failed to connect database")
	}
	return books
}
func (m *movieQuery) GetBookByID(id string) (models.Book, error) {

	var book models.Book
	result := m.db.Where("id = ?", id).First(&book)

	// fmt.Println(result.Error)

	// check error ErrRecordNotFound
	if x := errors.Is(result.Error, gorm.ErrRecordNotFound); x == true {
		error := &Errormessage{
			message: "book not found",
			status:  http.StatusNotFound,
		}

		return book, error
	} else {
		return book, nil
	}

}
func (m *movieQuery) CreateBook(book *models.Book) {
	var result = m.db.Create(book)
	if result.Error != nil {
		panic("failed to connect database")
	}
}
func (m *movieQuery) UpdateBook(book models.Book, id string) models.Book {
	m.db.Model(&book).Where("id = ?", id).Updates(book)

	return book
}
func (m *movieQuery) DeleteBook(id string) models.Book {
	var book models.Book
	m.db.Clauses(clause.Returning{}).Where("id = ?", id).Delete(&book)
	fmt.Print(book)
	return book

	// DELETE FROM users WHERE id = 10;
}
