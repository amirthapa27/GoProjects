package routes

import (
	"github.com/gorilla/mux"

	"mysqlAPI/pkg/controllers"
)

//crerating a function called RegisterBookStoreRoutes
// var RegisterBookStoreRoutes = func(router *mux.Router) {
// 	router.HandleFunc()
// }

func RegisterBookStoreRoutes(router *mux.Router) {
	router.HandleFunc("/book/", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/book/", controllers.GetBooks).Methods("GET")
	router.HandleFunc("/book/{bookId}", controllers.GetBook).Methods("GET")
	router.HandleFunc("/book/{bookId}", controllers.UpdateBook).Methods("PUT")
	router.HandleFunc("/book/{bookId}", controllers.DeleteBook).Methods("DELETE")
}
