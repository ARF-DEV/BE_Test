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

func CreateCategory(w http.ResponseWriter, r *http.Request) {
	newImages, err := helpers.SaveImage(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	body := r.FormValue("data")

	var newCategory models.Category
	err = json.Unmarshal([]byte(body), &newCategory)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	newCategory.Image = newImages[0]
	err = database.DB.Create(&newCategory).Error
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	responseBody, err := json.Marshal(newCategory)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(responseBody)
}
func GetAllCategory(w http.ResponseWriter, r *http.Request) {
	var categories []models.Category
	err := database.DB.Find(&categories).Error

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var responseBody, _ = json.Marshal(categories)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(responseBody)
}
func GetCategoryByID(w http.ResponseWriter, r *http.Request) {

	var body, err = ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var targetCategory models.Category
	err = json.Unmarshal(body, &targetCategory)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var categoryFound models.Category
	err = database.DB.Find(&categoryFound, targetCategory.ID).Error
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	responseBody, err := json.Marshal(categoryFound)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(responseBody)
}
func DeleteAllCategory(w http.ResponseWriter, r *http.Request) {
	var categories []models.Category

	var err = database.DB.Find(&categories).Error
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var ids []int
	for _, item := range categories {
		ids = append(ids, item.ID)
	}

	database.DB.Model(&models.Product{}).Where("category_id = ?", ids).Update("category_id", nil)
	err = database.DB.Delete(&categories).Error
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	responseBody, err := json.Marshal(categories)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(responseBody)

}
func DeleteCategoryByID(w http.ResponseWriter, r *http.Request) {
	var body, err = ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var targetCategory models.Category
	err = json.Unmarshal(body, &targetCategory)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var categoryFound models.Category
	err = database.DB.Find(&categoryFound, targetCategory.ID).Error
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	database.DB.Model(&models.Product{}).Where("category_id = ?", targetCategory).Update("category_id", nil)

	err = database.DB.Delete(&categoryFound, targetCategory.ID).Error
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	responseBody, err := json.Marshal(categoryFound)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(responseBody)
}
func UpdateCategoryByID(w http.ResponseWriter, r *http.Request) {
	newImages, err := helpers.SaveImage(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	body := r.FormValue("data")

	var targetCategory models.Category
	err = json.Unmarshal([]byte(body), &targetCategory)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Println(body)
	var oldCategory models.Category
	err = database.DB.First(&oldCategory, targetCategory.ID).Error
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var isUpdated = false
	if oldCategory.Name != targetCategory.Name {
		oldCategory.Name = targetCategory.Name
	}
	if oldCategory.Status != targetCategory.Status {
		oldCategory.Status = targetCategory.Status
	}

	if len(newImages) != 0 {
		var curPath, _ = os.Getwd()
		fmt.Println(len(oldCategory.Image))
		var path = strings.Trim(oldCategory.Image, r.Host)
		fmt.Println(filepath.Join(curPath, path))
		os.Remove(filepath.Join(curPath, path))

		oldCategory.Image = newImages[0]
		isUpdated = true
	}
	if isUpdated {
		oldCategory.UpdatedAt = targetCategory.CreatedAt
	}

	database.DB.Save(&oldCategory)

	responseBody, err := json.Marshal(oldCategory)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(responseBody)
}
