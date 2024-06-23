package controllers

import (
	"net/http"
	"sportnation/services"
	"sportnation/utils"
)

func GetAllProducts(w http.ResponseWriter, r *http.Request) {
    products, err := services.GetAllProducts()
    if err != nil {
        utils.RespondWithError(w, http.StatusInternalServerError, "Failed to get products")
        return
    }
    utils.RespondWithJSON(w, http.StatusOK, products)
}
