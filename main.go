package main

import (
	"gocommerce/configs"
	"gocommerce/handlers"
	"gocommerce/middlewares"
	"gocommerce/migrations"

	// "gocommerce/seeders"

	"net/http"
	_ "net/http/pprof" // Import pprof to enable the pprof handlers

	"github.com/gin-gonic/gin"
)

func main() {

	db, err := configs.InitDB()
	if err != nil {
		panic(err)
	}
	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}
	defer sqlDB.Close()

	migrations.Migrate(db)
	// seeders.Seed(db)

	server := gin.Default()
	router := server
	//PRODUCTS
	router.GET("/products", middlewares.AuthMiddleware(), handlers.ListProducts(db))
	router.GET("/products/:id", handlers.GetProduct(db))
	router.POST("/products", handlers.CreateProduct(db))
	router.PUT("/products/:id", handlers.UpdateProduct(db))
	router.DELETE("/products/:id", handlers.DeleteProduct(db))

	// Rute CRUD Product Category
	router.GET("/product-categories", handlers.ListProductCategories(db))
	router.GET("/product-categories/:id", handlers.GetProductCategory(db))
	router.POST("/product-categories", handlers.CreateProductCategory(db))
	router.PUT("/product-categories/:id", handlers.UpdateProductCategory(db))
	router.DELETE("/product-categories/:id", handlers.DeleteProductCategory(db))

	// Transaction
	router.POST("/transactions", handlers.CreateTransaction(db))
	router.GET("/transactions/:id", handlers.GetTransactionWithItems(db))

	router.POST("/login", handlers.Login(db))
	router.POST("/register", handlers.Register(db))

	router.GET("/debug/pprof/*pprof", gin.WrapH(http.DefaultServeMux))
	server.Run(":8000")
}
