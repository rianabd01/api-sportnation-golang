package config

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func ConnectDB() (*sql.DB, error) {
    db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/sportnation_app")
    if err != nil {
        fmt.Println("Unable to connect!")
        return nil, err
    }
    return db, nil
}
