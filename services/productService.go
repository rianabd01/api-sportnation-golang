package services

import (
	"sportnation/config"
	"sportnation/models"
	"sportnation/repositories"
)

func GetAllProducts() ([]models.Product, error) {
    db, err := config.ConnectDB()
    if err != nil {
        return nil, err
    }
    defer db.Close()

    return repositories.GetAllProducts(db)
}
