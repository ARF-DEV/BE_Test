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
	router.HandleFunc("/api/product/all", controller.GetAllProducts).Methods("GET")
	router.HandleFunc("/api/product/id", controller.GetProductByID).Methods("GET")
	router.HandleFunc("/api/product/delete/all", controller.DeleteAllProduct).Methods("DELETE")
	router.HandleFunc("/api/product/delete/id", controller.DeleteProductByID).Methods("DELETE")
	router.HandleFunc("/api/product/create", controller.CreateProduct).Methods("POST")
	router.HandleFunc("/api/product/update/id", controller.UpdateProductByID).Methods("PUT")
	router.HandleFunc("/api/product/category_id", controller.GetProductByCategoryID).Methods("Get")

	router.HandleFunc("/api/category/all", controller.GetAllCategory).Methods("GET")
	router.HandleFunc("/api/category/id", controller.GetCategoryByID).Methods("GET")
	router.HandleFunc("/api/category/delete/all", controller.DeleteAllCategory).Methods("DELETE")
	router.HandleFunc("/api/category/delete/id", controller.DeleteCategoryByID).Methods("DELETE")
	router.HandleFunc("/api/category/create", controller.CreateCategory).Methods("POST")
	router.HandleFunc("/api/category/update/id", controller.UpdateCategoryByID).Methods("PUT")

	//dunno what happen here (especially the http.StripPrefix)

	http.ListenAndServe(":5000", router)
}

func main() {
	database.InitDB()
	HandleRequest()

}
