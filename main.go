package main

import (
	"ar_backend/database"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

const (
	webport = "8080"
)

func main() {
	db, err := database.ConnectDB()
	if err == nil {
		err = database.CreateTable(db)
		if err == nil {
			router := gin.Default()
			router.Run("localhost:" + webport)
		}
	}
}
