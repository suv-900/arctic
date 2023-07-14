package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/suv-900/netflix-clone/controllers"
	"github.com/suv-900/netflix-clone/routes"
)

func main() {
	router := mux.NewRouter()
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://locahost:3000"},
		AllowedMethods: []string{http.MethodGet, http.MethodPost,
			http.MethodPut, http.MethodPut},
		AllowCredentials: true,
	})

	//fs:=http.FileServer(http.Dir("./static"))
	//router.Handle("/",fs)
	handler := c.Handler(router)
	controllers.ConnectAndMigrateDB()
	routes.HandleRoutes(router)
	fmt.Println("server started on port: 8000")
	log.Fatal(http.ListenAndServe(":8000", handler))
}
