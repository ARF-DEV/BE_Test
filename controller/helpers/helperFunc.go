package helpers

import (
	"encoding/base64"
	"errors"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/ARF-DEV/BE_Test/models"
)

func UploadProductImage(r *http.Request) ([]models.ProductImage, error) {
	var newImages []models.ProductImage
	err := r.ParseMultipartForm(100 << 20)
	if err != nil {

		return []models.ProductImage{}, err
	}
	formData := r.MultipartForm
	images := formData.File["files"]
	for i := range images {

		file, err := images[i].Open()
		defer file.Close()

		if err != nil {
			return []models.ProductImage{}, err
		}
		curPath, _ := os.Getwd()
		extension := filepath.Ext(images[i].Filename)
		name := images[i].Filename[0 : len(images[i].Filename)-len(extension)]

		filePath := "/static/images/" + base64.StdEncoding.EncodeToString([]byte(name)) + extension
		out, err := os.Create(filepath.Join(curPath, filePath))
		if err != nil {
			return []models.ProductImage{}, errors.New("unable to create the file for writing. Check your write access privilege")
		}
		defer out.Close()

		_, err = io.Copy(out, file)

		if err != nil {
			return []models.ProductImage{}, err
		}
		var newProductImage = models.ProductImage{ImagePath: r.Host + filePath}
		newImages = append(newImages, newProductImage)
	}

	return newImages, nil
}

func SaveImage(r *http.Request) ([]string, error) {
	err := r.ParseMultipartForm(100 << 20)
	if err != nil {

		return nil, err
	}
	formData := r.MultipartForm
	images := formData.File["files"]
	var imagesPath []string
	for i := range images {

		file, err := images[i].Open()
		defer file.Close()

		if err != nil {
			return nil, err
		}
		curPath, _ := os.Getwd()
		extension := filepath.Ext(images[i].Filename)
		name := images[i].Filename[0 : len(images[i].Filename)-len(extension)]

		filePath := "/static/images/" + base64.StdEncoding.EncodeToString([]byte(name)) + extension
		out, err := os.Create(filepath.Join(curPath, filePath))
		if err != nil {
			return nil, errors.New("unable to create the file for writing. Check your write access privilege")
		}
		defer out.Close()

		_, err = io.Copy(out, file)

		if err != nil {
			return nil, err
		}
		var newImagePath = r.Host + filePath
		imagesPath = append(imagesPath, newImagePath)
	}

	return imagesPath, nil
}
