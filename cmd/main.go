// The only goal of the main to to be

package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/henryikoh/book-management-go/pkg/myhttp"
	"github.com/henryikoh/book-management-go/pkg/repositery"
)

func main() {

	// init Data Access Object
	dao := repositery.InitDAO()

	// pass it into serverhandler this gives all handles access to the dao
	// This also intainsiates all the handlers to the routes
	srv := myhttp.NewServer(*dao)

	server := &http.Server{

		Addr:    ":8080",
		Handler: srv,

		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	// listen and server at port :8080
	fmt.Println("starting server at port 8080")
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}

	// fmt.Println("starting server at port 8080")
	// if err := http.ListenAndServe(":8080", srv); err != nil {
	// 	log.Fatal(err)
	// }
}
