package controllers

import (
	"Routes/src/models"
	"Routes/src/response"
	"Routes/src/service"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	var productInput models.ProductRequestVO
	if err := json.NewDecoder(r.Body).Decode(&productInput); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var product, err = service.CreateProductService(productInput.ProductName, productInput.CostPrice, productInput.SellPrice)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response.JSON(w, http.StatusOK, product)
}

func FindProductById(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	productId, err := strconv.ParseInt(parametros["productid"], 10, 64)
	if err != nil {
		response.Error(w, http.StatusBadRequest, err)
		return
	}

	product, err := service.FindProductById(productId)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
	}

	response.JSON(w, http.StatusOK, product)
}

func FindAllProducts(w http.ResponseWriter, r *http.Request) {
	products, err := service.GetAllProducts()
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
	}

	response.JSON(w, http.StatusOK, products)
}

func UpdateProductParsingId(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	productId, err := strconv.ParseInt(parametros["productid"], 10, 64)
	if err != nil {
		response.Error(w, http.StatusBadRequest, err)
		return
	}

	var productInput models.ProductRequestVO
	if err := json.NewDecoder(r.Body).Decode(&productInput); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	product, err := service.UpdateProductById(productId, productInput)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
	}

	response.JSON(w, http.StatusOK, product)
}

func DeleteProductParsingId(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	productId, err := strconv.ParseInt(parametros["productid"], 10, 64)
	if err != nil {
		response.Error(w, http.StatusBadRequest, err)
		return
	}

	if err := service.DeleteProductParsingId(productId); err != nil {
		response.Error(w, http.StatusInternalServerError, err)
	}

	response.JSON(w, http.StatusNoContent, nil)
}
