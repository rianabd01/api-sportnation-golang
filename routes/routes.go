package routes

import (
	"sportnation/controllers"

	"github.com/gorilla/mux"
)

func SetupRoutes() *mux.Router {
    router := mux.NewRouter()

    router.HandleFunc("/hello", controllers.HelloHandler).Methods("GET")
    router.HandleFunc("/products", controllers.GetAllProducts).Methods("GET")

    return router
}
