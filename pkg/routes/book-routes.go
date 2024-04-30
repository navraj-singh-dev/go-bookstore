package routes

import (
	"github.com/gorilla/mux"
	"github.com/navraj-singh-dev/go-bookstore/pkg/controllers"
)

func RegisterBookRoutes(router *mux.Router) {
	router.HandleFunc("/book", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/book", controllers.GetAllBooks).Methods("GET")
	router.HandleFunc("/book/{bookId}", controllers.GetBookById).Methods("GET")
	router.HandleFunc("/book/{bookId}", controllers.UpdateBookById).Methods("PUT")
	router.HandleFunc("/book/{bookId}", controllers.DeleteById).Methods("DELETE")

}
