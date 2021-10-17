package main

import (
	"log"
	"net/http"

	"github.com/go-go/examples/apiWithMux/controllers"
	"github.com/gorilla/mux"
)

func main() {
	// init router
	r := mux.NewRouter()

	controllers.InitMock()

	// route handlers / endpoints
	r.HandleFunc("/api/books", controllers.GetBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", controllers.GetBook).Methods("GET")
	r.HandleFunc("/api/books", controllers.CreateBook).Methods("POST")
	r.HandleFunc("/api/books/{id}", controllers.UpdateBook).Methods("PUT")
	r.HandleFunc("/api/books/{id}", controllers.DeleteBooks).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", r))
}
