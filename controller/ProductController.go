package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/ARF-DEV/BE_Test/database"
	"github.com/ARF-DEV/BE_Test/models"
)

func GetAllProducts(w http.ResponseWriter, r *http.Request) {
	var products []models.Product

	database.DB.Preload("ProductImages").Find(&products)

	jsonResponse, err := json.Marshal(products)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Print(string(jsonResponse[0]))

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(products)

}
func GetProductByID(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var requestBody models.Product
	err = json.Unmarshal(body, &requestBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var product models.Product
	database.DB.Preload("ProductImages").Find(&product, requestBody.ID)

	responseBody, err := json.Marshal(product)
	fmt.Println(string(responseBody))
	if err != nil {
		fmt.Print(err.Error())
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(responseBody)

}
