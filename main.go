package main 

import (
	"github.com/SunnyChandraa/go-restapi-gin/controllers/productController"
	"github.com/SunnyChandraa/go-restapi-gin/models"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default();
	models.ConnectDatabase()

	r.GET("api/products", productController.Index)
	r.GET("api/products/:id", productController.Show)
	r.POST("api/products", productController.Create)
	r.PUT("api/products/update/:id", productController.Update)
	r.DELETE("api/products/delete", productController.Delete)

	r.Run()
}