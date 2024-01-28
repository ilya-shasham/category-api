package main

import (
	"categoryAPI/globals"
	"categoryAPI/handlers"
	"categoryAPI/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	var err error
	globals.Db, err = gorm.Open(sqlite.Open("db.db"))

	if err != nil {
		log.Fatalln("failed to initialize database: ", err)
	}

	globals.Db.AutoMigrate(models.Registered...)

	app := gin.Default()

	// TODO: add an update route
	app.POST("/api/v1/category/", handlers.AddCategory)
	app.DELETE("/api/v1/category/:ids", handlers.DeleteCategory)
	app.GET("/api/v1/category/all", handlers.GetAllCategories)
	app.GET("/api/v1/category/range/:from/:to", handlers.GetCategoriesRanged)
	app.GET("/api/v1/category/prefix/:contains", handlers.GetCategoriesContaining)
	app.GET("/api/v1/category/id/:id", handlers.GetCategoryById)

	server := &http.Server{
		Addr:    "localhost:9090",
		Handler: app,
	}

	server.SetKeepAlivesEnabled(false)
	server.ListenAndServe()
}
