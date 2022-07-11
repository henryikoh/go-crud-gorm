package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/henryikoh/book-management-go/pkg/models"
	"github.com/henryikoh/book-management-go/pkg/repositery"
)

type Handlers struct {
	// db would contain the current database ( its currently being used to instansiate the Movies repo )
	db repositery.DAO
}

type message struct {
	Message string `json:"message"`
	Status  int64  `json:"status"`
}

func JSONError(w http.ResponseWriter, err message) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(int(err.Status))
	json.NewEncoder(w).Encode(err)
}

func InitHandlers(doa repositery.DAO) *Handlers {
	// http.NewServeMux().router.
	s := &Handlers{
		db: doa,
	}
	return s
}

// helpers are fucntions that can help handle repeared tasks like writting a json object to the reposone writer
func (h *Handlers) ToJSON(w http.ResponseWriter, mov interface{}) error {
	w.Header().Set("content-type", "application/json")
	e := json.NewEncoder(w)
	return e.Encode(mov)
}
func (h *Handlers) DecodeBody(w http.ResponseWriter, r *http.Request) models.Book {
	var movie models.Book
	err := json.NewDecoder(r.Body).Decode(&movie)
	if err != nil {
		panic("could not parse the body to json")
	}
	return movie
}

// uses the post method to create a book
func (h *Handlers) CreateBook() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		book := h.DecodeBody(w, r)

		fmt.Printf("%v", book)
		h.db.NewMovieQuery().CreateBook(&book)

		h.ToJSON(w, book)
	}
}

// uses the get method to get all books
func (h *Handlers) GetBooks() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		books := h.db.NewMovieQuery().GetBooks()
		h.ToJSON(w, books)
	}
}

// uses the get method to get a book by ID
func (h *Handlers) GetBookById() http.HandlerFunc {
	// nothing here for now but this space can be used to prepare and dependencies

	return func(w http.ResponseWriter, r *http.Request) {

		// uses mux to fetch the ID from the URL
		vars := mux.Vars(r)
		bookId := vars["id"]

		// send the ID to the book queryinterface
		book, err := h.db.NewMovieQuery().GetBookByID(bookId)

		// fmt.Println(err)

		// if there is an error send error back to client
		if err != nil {
			// compost the http error
			messages := message{
				Message: err.Error(),
				Status:  http.StatusNotFound,
			}
			JSONError(w, messages)

		} else {
			// if not erro send the result back to client
			h.ToJSON(w, book)
		}

	}
}

// uses the put method to update a book by ID
func (h *Handlers) UpdateBookById() http.HandlerFunc {
	// nothing here for now but this space can be used to prepare and dependencies

	return func(w http.ResponseWriter, r *http.Request) {

		// uses mux to fetch the ID from the URL
		vars := mux.Vars(r)
		bookId := vars["id"]

		// get the content of the body
		book := h.DecodeBody(w, r)

		// send the ID to the book queryinterface
		result := h.db.NewMovieQuery().UpdateBook(book, bookId)

		fmt.Println(result)

		// if there is an error send error back to client
		// if err != nil {
		// 	// compost the http error
		// 	messages := message{
		// 		Message: err.Error(),
		// 		Status:  http.StatusNotFound,
		// 	}
		// 	JSONError(w, messages)

		// } else {
		// 	// if not erro send the result back to client
		// 	h.ToJSON(w, book)
		// }

	}
}

// uses the del method to delete a book by ID
func (h *Handlers) DeleteBookById() http.HandlerFunc {
	// nothing here for now but this space can be used to prepare and dependencies

	return func(w http.ResponseWriter, r *http.Request) {

		// uses mux to fetch the ID from the URL
		vars := mux.Vars(r)
		bookId := vars["id"]

		// send the ID to the book deleteinterface
		result := h.db.NewMovieQuery().DeleteBook(bookId)

		fmt.Println(result)

		// if there is an error send error back to client
		// if err != nil {
		// 	// compost the http error
		// 	messages := message{
		// 		Message: err.Error(),
		// 		Status:  http.StatusNotFound,
		// 	}
		// 	JSONError(w, messages)

		// } else {
		// 	// if not erro send the result back to client
		// 	h.ToJSON(w, book)
		// }

	}
}
