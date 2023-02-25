package routes

import (
	"log"
	"net/http"

	"go-desenvolvendo-api-rest/controllers"
	"go-desenvolvendo-api-rest/middleware"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func HandlleRequest() {
	r := mux.NewRouter()
	r.Use(middleware.ContentTypeMidleware)
	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/api/personalidades", controllers.TodasPersonalidades).Methods("Get")
	r.HandleFunc("/api/personalidades/{id}", controllers.RetornaUmaPersonalidade).Methods("Get")
	r.HandleFunc("/api/personalidades", controllers.CriaPersonalidade).Methods("Post")
	r.HandleFunc("/api/personalidades/{id}", controllers.DeletaPersonalidade).Methods("Delete")
	r.HandleFunc("/api/personalidades/{id}", controllers.EditaPersonalidade).Methods("Put")
	log.Fatal(http.ListenAndServe(":8000", handlers.CORS(handlers.AllowedOrigins([]string{"*"}))(r)))
}
