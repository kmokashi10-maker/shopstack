package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Product struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
	Stock int     `json:"stock"`
}

// temporary product data
var products = []Product{
	{ID: 1, Name: "Laptop", Price: 55000, Stock: 10},
	{ID: 2, Name: "Keyboard", Price: 1500, Stock: 25},
	{ID: 3, Name: "Mouse", Price: 800, Stock: 30},
}

func main() {
	//home route
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Welcome to Shopstack API",
		})
	})

	// get display pruducts
	router.GET("/products", func(c *gin.Context) {
		c.JSON(http.StatusOK, products)
	})

	//product by id
	router.GET("/products/:id", func(c *gin.Context) {
		id := c.Param("id")
		for _, product := range products {
			if fmt.Sprint(product.ID) == id {

				c.JSON(http.StatusOK, product)
			}
		}
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Product not found",
		})
	})
	//POST : creat or adding new product
	router.POST("products", func(c *gin.Context) {
		var newProduct Product

		if err := c.ShouldBindJSON(&newProduct); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}

		newProduct.ID = len(products) + 1
		products = append(products, newProduct)

		c.JSON(http.StatusCreated, newProduct)
	})

	router.Run(":8080")
}
