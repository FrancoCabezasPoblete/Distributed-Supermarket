package controller

import (
	"net/http"
	"strconv"

	"tarea1/golang-api/entity"
	products "tarea1/golang-api/protobuf"
	"tarea1/golang-api/storage"

	"github.com/gin-gonic/gin"
)

func GetProducts(c *gin.Context) {
	products := storage.GetProducts()
	c.IndentedJSON(http.StatusOK, products)

	/*if products == nil || len(products) == 0 {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.IndentedJSON(http.StatusOK, products)
	}*/
}

func GetProduct(c *gin.Context) {
	productId, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		panic(err)
	}

	product := storage.GetProduct(productId)

	if product == nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.IndentedJSON(http.StatusOK, product)
	}
}

func AddProduct(c *gin.Context) {
	var product entity.Product

	if err := c.BindJSON(&product); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	} else {
		latestId := storage.AddProduct(product)
		c.IndentedJSON(http.StatusCreated, gin.H{
			"id_producto": latestId,
		})
	}
}

func DeleteProduct(c *gin.Context) {
	productId, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		panic(err)
	}

	latestId := storage.DeleteProduct(productId)
	c.IndentedJSON(http.StatusOK, gin.H{
		"id_producto": latestId,
	})
}

func UpdateProduct(c *gin.Context) {
	productId, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		panic(err)
	}

	product := storage.GetProduct(productId)

	if err := c.BindJSON(&product); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	} else {
		latestId := storage.UpdateProduct(product)
		c.IndentedJSON(http.StatusCreated, gin.H{
			"id_producto": latestId,
		})
	}
}

func RequestProduct(client products.ProductsClient) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		var product entity.Product

		if err := c.BindJSON(&product); err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
		} else {

			req := &products.SendProductRequest{IdProducto: int32(product.IdProduct), Nombre: product.Name, CantidadDisponible: int32(product.Available)}

			if res, err := client.SendProduct(c, req); err != nil {
				c.AbortWithStatus(http.StatusBadRequest)
			} else {
				product.Available = int(res.CantidadDisponible)
				latestId := storage.UpdateProduct(&product)
				c.IndentedJSON(http.StatusCreated, gin.H{
					"id_producto": latestId,
				})
			}
		}
	}

	return gin.HandlerFunc(fn)
}
