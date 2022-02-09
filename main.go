package main

import (
	"net/http"
	"os"

	"github.com/ARF-DEV/BE_Test/controller"
	"github.com/ARF-DEV/BE_Test/database"
	"github.com/gorilla/mux"
)

func HandleRequest() {
	dir := "/static/"
	os.Mkdir(dir, 0777)
	router := mux.NewRouter()
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
	router.HandleFunc("/api/product/all", controller.GetAllProducts)
	router.HandleFunc("/api/product/id", controller.GetProductByID)
	router.HandleFunc("/api/product/delete/all", controller.DeleteAllProduct)
	router.HandleFunc("/api/product/delete/id", controller.DeleteProductByID)
	router.HandleFunc("/api/product/create", controller.CreateProduct).Methods("POST")
	router.HandleFunc("/api/product/update/id", controller.UpdateProductByID).Methods("PUT")

	//dunno what happen here (especially the http.StripPrefix)

	http.ListenAndServe(":5000", router)
}

func main() {
	database.InitDB()
	HandleRequest()

}
