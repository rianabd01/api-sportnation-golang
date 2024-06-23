package repositories

import (
	"database/sql"
	"sportnation/models"
)

func GetAllProducts(db *sql.DB) ([]models.Product, error) {
    rows, err := db.Query("SELECT productId, name, quantity FROM products")
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    products := []models.Product{}
    for rows.Next() {
        var product models.Product
        if err := rows.Scan(&product.ID, &product.Name, &product.Quantity); err != nil {
            return nil, err
        }
        products = append(products, product)
    }
    return products, nil
}
