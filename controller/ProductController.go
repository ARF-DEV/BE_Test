package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/ARF-DEV/BE_Test/controller/helpers"
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
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)

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
	err = database.DB.Preload("ProductImages").Find(&product, requestBody.ID).Error
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	responseBody, err := json.Marshal(product)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(responseBody)

}

func DeleteProductByID(w http.ResponseWriter, r *http.Request) {
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
	err = database.DB.Preload("ProductImages").Find(&product, requestBody.ID).Error
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	database.DB.Select("ProductImages").Delete(&models.Product{ID: requestBody.ID})

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	database.DB.Delete(&product)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	responseBody, _ := json.Marshal(product)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(responseBody)
}

func DeleteAllProduct(w http.ResponseWriter, r *http.Request) {
	var products []models.Product

	database.DB.Preload("ProductImages").Find(&products)

	database.DB.Select("ProductImages").Delete(&products)

	jsonResponse, err := json.Marshal(products)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	newImages, err := helpers.UploadProductImage(r)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	body := r.FormValue("data")
	var requestBodyMap map[string]interface{}
	json.Unmarshal([]byte(body), &requestBodyMap)

	var requestedCategory models.Category
	database.DB.First(&requestedCategory, &models.Category{Name: requestBodyMap["CategoryName"].(string)})

	var newProduct models.Product
	json.Unmarshal([]byte(body), &newProduct)

	newProduct.Category = requestedCategory
	newProduct.ProductImages = newImages
	database.DB.Create(&newProduct)

	responseBody, _ := json.Marshal(newProduct)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(responseBody)

}
func UpdateProductByID(w http.ResponseWriter, r *http.Request) {
	newImages, err := helpers.UploadProductImage(r)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	body := r.FormValue("data")

	var requestBodyMap map[string]interface{}
	json.Unmarshal([]byte(body), &requestBodyMap)
	var oldCategory models.Category
	if requestBodyMap["CategoryName"] != nil {
		database.DB.First(&oldCategory, &models.Category{Name: requestBodyMap["CategoryName"].(string)})
	}

	var updateProduct models.Product
	json.Unmarshal([]byte(body), &updateProduct)

	var oldProduct models.Product
	err = database.DB.Preload("ProductImages").First(&oldProduct, &models.Product{ID: updateProduct.ID}).Error
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	var isUpdated = false
	if updateProduct.Name != "" && updateProduct.Name != oldProduct.Name {
		oldProduct.Name = updateProduct.Name
		isUpdated = true
	}
	if updateProduct.Description != "" && updateProduct.Description != oldProduct.Description {
		oldProduct.Description = updateProduct.Description
		isUpdated = true
	}
	if updateProduct.Stock != false && updateProduct.Stock != oldProduct.Stock {
		oldProduct.Stock = updateProduct.Stock
		isUpdated = true
	}
	if updateProduct.Price != 0 && updateProduct.Price != oldProduct.Price {
		oldProduct.Price = updateProduct.Price
		isUpdated = true
	}
	if updateProduct.Price_type != 0 && updateProduct.Price_type != oldProduct.Price_type {
		oldProduct.Price_type = updateProduct.Price_type
		isUpdated = true
	}
	if updateProduct.Status != false && updateProduct.Status != oldProduct.Status {
		oldProduct.Status = updateProduct.Status
		isUpdated = true
	}

	if len(newImages) != 0 {
		var curPath, _ = os.Getwd()
		fmt.Println(len(oldProduct.ProductImages))
		for _, image := range oldProduct.ProductImages {
			var path = strings.Trim(image.ImagePath, r.Host)
			fmt.Println(filepath.Join(curPath, path))
			os.Remove(filepath.Join(curPath, path))
		}
		database.DB.Where("product_id = ?", oldProduct.ID).Delete(&models.ProductImage{})
		oldProduct.ProductImages = newImages
	}

	if (oldCategory == models.Category{}) && (updateProduct.Category != oldProduct.Category) {
		oldProduct.Category = updateProduct.Category
	}

	if isUpdated {
		oldProduct.UpdatedAt = updateProduct.UpdatedAt
	}

	database.DB.Save(&oldProduct)

	responseBody, _ := json.Marshal(oldProduct)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(responseBody)
}
