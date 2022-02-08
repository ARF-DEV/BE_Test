package main

import (
	"net/http"

	"github.com/ARF-DEV/BE_Test/controller"
	"github.com/ARF-DEV/BE_Test/database"
	"github.com/gorilla/mux"
)

func HandleRequest() {
	router := mux.NewRouter()
	router.HandleFunc("/api/product/all", controller.GetAllProducts)
	router.HandleFunc("/api/product/id", controller.GetProductByID)

	//dunno what happen here (especially the http.StripPrefix)
	router.PathPrefix("/file/").Handler(http.StripPrefix("/file/", http.FileServer(http.Dir("./static"))))

	http.ListenAndServe(":5000", router)
}

func main() {
	database.InitDB()
	HandleRequest()

}
