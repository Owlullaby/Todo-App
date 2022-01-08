package router

import (
	"todo_app/middleware"
	"github.com/gorilla/mux"
	"net/http"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	router.PathPrefix("/template").Handler(http.StripPrefix("/template/",http.FileServer(http.Dir("./template/"))))
	// router.HandleFunc("/todo/{id}", middleware.GetTodo).Methods("GET")
	router.HandleFunc("/", middleware.GetAllTodo).Methods("GET")
	router.HandleFunc("/addtodo", middleware.CreateTodo).Methods("POST")
	router.HandleFunc("/todo/{id}", middleware.UpdateTodo)
	router.HandleFunc("/deletetodo/{id}", middleware.DeleteTodo)

	return router
}
