package main

import (
	"ar_backend/database"
	"ar_backend/route"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

const (
	webport = "8083"
)

func main() {
	db, err := database.ConnectDB()
	if err == nil {
		err = database.CreateTable(db)
		if err == nil {
			router := gin.Default()
			router.GET("/infos", getInfos)
			router.GET("/info/:code", getInfo)
			router.POST("/info", insertInfo)
			router.Run("localhost:" + webport)
		}
	}
}

func getInfos(c *gin.Context) {
	db, err := database.ConnectDB()
	if err == nil {
		defer db.Close()

		infos := route.GetInfos(db)
		if infos == nil || len(infos) == 0 {
			c.AbortWithStatus(http.StatusNotFound)
		} else {
			c.IndentedJSON(http.StatusOK, infos)
		}
	}
}

func getInfo(c *gin.Context) {
	code := c.Param("code")

	db, err := database.ConnectDB()
	if err == nil {
		defer db.Close()

		info := route.GetInfo(db, code)
		if info == nil {
			c.AbortWithStatus(http.StatusNotFound)
		} else {
			c.IndentedJSON(http.StatusOK, info)
		}

	}
}

func insertInfo(c *gin.Context) {
	db, err := database.ConnectDB()
	if err == nil {
		defer db.Close()

		var info route.Info
		if err := c.BindJSON(&info); err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
		} else {
			route.InsertInfo(db, info)
			c.IndentedJSON(http.StatusCreated, info)
		}
	}
}
