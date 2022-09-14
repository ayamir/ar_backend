package main

import (
	"ar_backend/database"
	"ar_backend/route"
	"net/http"

	"github.com/gin-contrib/cors"
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
			router.Use(cors.New(cors.Config{
				AllowOrigins: []string{"http://localhost", "http://10.112.79.143"},
				AllowMethods: []string{"GET", "POST"},
				AllowHeaders: []string{"Authorization", "Content-Type", "Upgrade", "Origin",
					"Connection", "Accept-Encoding", "Accept-Language", "Host", "Access-Control-Request-Method", "Access-Control-Request-Headers"},
			}))

			router.GET("/infos", getInfos)
			router.GET("/info/:code", getInfo)

			router.POST("/info", insertInfo)
			router.POST("/info/motto", updateInfo)

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
		if err = c.BindJSON(&info); err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
		} else {
			err = route.InsertInfo(db, info)
			if err != nil {
				c.AbortWithStatus(http.StatusInternalServerError)
			} else {
				c.IndentedJSON(http.StatusCreated, info)
			}
		}
	}
}

func updateInfo(c *gin.Context) {
	db, err := database.ConnectDB()
	if err == nil {
		defer db.Close()

		var infoMotto route.InfoMotto
		if err = c.ShouldBindJSON(&infoMotto); err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
		} else {
			err = route.UpdateInfo(db, infoMotto)
			if err != nil {
				c.AbortWithStatus(http.StatusInternalServerError)
			} else {
				c.IndentedJSON(http.StatusOK, infoMotto)
			}
		}
	}
}
