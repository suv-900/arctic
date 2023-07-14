package routes

import (
	"github.com/gorilla/mux"
	"github.com/suv-900/netflix-clone/controllers"
)

var HandleRoutes = func(router *mux.Router) {
	router.HandleFunc("/register", controllers.CreatNewUser).Methods("POST")
	router.HandleFunc("/login", controllers.LoginUser).Methods("POST")
	router.HandleFunc("/checkusername", controllers.SearchForSimilarUsernames).Methods("POST")
	router.HandleFunc("/checkserver", controllers.CheckServerHealth).Methods("GET")
}
