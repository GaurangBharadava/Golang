package router

import (
	"github.com/GaurangBharadava/mongoapi/controller"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api/movies",controller.GetAllMovies).Methods("GET")
	router.HandleFunc("/api/movie",controller.CreateMovie).Methods("POST")
	router.HandleFunc("/api/movie/{id}",controller.MarkedAsWatched).Methods("PUT")
	router.HandleFunc("/api/movie/{id}",controller.DeleteOneMovie).Methods("DELETE")
	router.HandleFunc("/api/deleteall",controller.DeleteAllMovie).Methods("DELETE")
	return router
}