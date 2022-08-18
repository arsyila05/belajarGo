package controllers

import (
	"belajar-golang/database"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var product database.Product
	json.NewDecoder(r.Body).Decode(&product)
	database.Instance.Create(&product)
	json.NewEncoder(w).Encode(product)
}
func checkIfProductExists(productId string) bool {
	var product database.Product
	database.Instance.First(&product, productId)
	if product.Product_ID == 0 {
		return false
	}
	return true
}
func GetProductById(w http.ResponseWriter, r *http.Request) {
	productId := mux.Vars(r)["id"]
	if checkIfProductExists(productId) == false {
		json.NewEncoder(w).Encode("Product Not Found!")
		return
	}
	var product database.Product
	database.Instance.First(&product, productId)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(product)
}
func GetProducts(w http.ResponseWriter, r *http.Request) {
	var products []database.Product
	database.Instance.Find(&products)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(products)
}
func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	productId := mux.Vars(r)["id"]
	if checkIfProductExists(productId) == false {
		json.NewEncoder(w).Encode("Product Not Found!")
		return
	}
	var product database.Product
	database.Instance.First(&product, productId)
	json.NewDecoder(r.Body).Decode(&product)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(product)
	err := database.Instance.Model(database.Product{}).Where("product_id=?", productId).Updates(&product)

	if err != nil {
		fmt.Println(err)
	}
}
func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	productId := mux.Vars(r)["id"]
	if checkIfProductExists(productId) == false {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("Product Not Found!")
		return
	}
	var product database.Product
	database.Instance.Delete(&product, productId)
	json.NewEncoder(w).Encode("Product Deleted Successfully!")
}
