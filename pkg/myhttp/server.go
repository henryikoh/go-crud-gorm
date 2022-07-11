/*
	This page is the main server of the project. It would contain all the routers and handlers
	on this sever. They may be a need for multiple servers if this page grows too large.

	All the handler functions are attached to the server this makes it easier to share dependecies and also
	keeps evethings in one place. Like middlewares and helpers. see example below

	could also be named rest handlers or myhttp or routes
*/

// this was orginially named "handlers" but I changed the name to "myhttp" as I feel this might be more descriptive
package myhttp

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/henryikoh/book-management-go/pkg/myhttp/handlers"
	"github.com/henryikoh/book-management-go/pkg/repositery"
)

/*
	The server is a struct and would contain the dependencies its need within the struct also server varieables
	can be definded here.

*/

type Server struct {
	// Router is currently the default server mux, This could also be gorlaar mux or and servemux struct
	// This type would change but ideally this should be a type that is a router
	router *mux.Router

	// db would contain the current database ( its currently being used to instansiate the Movies repo )
	db       repositery.DAO
	handlers *handlers.Handlers
}

// New serve is used to created the server and also assign the routes the the router
func NewServer(dao repositery.DAO) *Server {
	// http.NewServeMux().router.
	s := &Server{
		// other server variables can also be set here. Check serve struck docs for more info
		// if the variable is capital it would be exported.

		// I implemented the defaul server mux which works
		router:   mux.NewRouter(),
		db:       dao,
		handlers: handlers.InitHandlers(dao),
	}

	s.routes()
	return s
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// My server can be used has an http handlers because it impliments the serverHTTP interface
	// the server is stared here

	// Once the serve is called it is rirected to your router of choice
	// I implemented with but MUX and the default router. ( its your picks )

	// s.router.ServeHTTP(w, r)

	s.router.ServeHTTP(w, r)
}

// here us a list of all routes
func (s *Server) routes() {
	// the routes a are unique because they are not handler functions directly
	// but ther are fucntions that return handler functions. this results in a
	// closure that can be used to do other cool things.

	// Here is the mux example

	s.router.HandleFunc("/book", s.handlers.CreateBook()).Methods("POST")
	// s.router.HandleFunc("/book", s.CreateBook()).Methods("POST")
	s.router.HandleFunc("/books", s.handlers.GetBooks()).Methods("GET")
	s.router.HandleFunc("/book/{id}", s.handlers.GetBookById()).Methods("GET")
	s.router.HandleFunc("/book/{id}", s.handlers.UpdateBookById()).Methods("PUT")
	s.router.HandleFunc("/book/{id}", s.handlers.DeleteBookById()).Methods("DELETE")

}

// Middleware
// Middleware in go is just a functions that takes in a handler function and return a handler function

// This functions just chceks to see if the request is a get request
func (s *Server) GetMiddleware(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Invalid method", http.StatusBadRequest)
			return
		}
		h(w, r)
	}
}

// helpers are fucntions that can help handle repeared tasks like writting a json object to the reposone writer
func (s *Server) ToJSON(w http.ResponseWriter, mov interface{}) error {
	w.Header().Set("content-type", "application/json")
	e := json.NewEncoder(w)
	return e.Encode(mov)
}
