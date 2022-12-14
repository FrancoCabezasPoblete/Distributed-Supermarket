package controller

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"tarea1/golang-api/entity"
	"tarea1/golang-api/storage"

	"github.com/gin-gonic/gin"
)

func AddShopping(c *gin.Context) {
	var shopping entity.Shopping

	if err := c.BindJSON(&shopping); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		fmt.Println(err)
	} else {
		lastId := storage.AddShopping(shopping)
		deliveryId := rand.Intn(8*10e6) + 10e6

		values, _ := json.Marshal(entity.Delivery{
			IdDespacho: deliveryId,
			IdCompra:   lastId,
		})

		responseBody := bytes.NewBuffer(values)

		URL := "http://10.10.11.204:5000/api/supermercado/despacho"
		resp, _ := http.Post(URL, "application/json", responseBody)

		resp.Body.Close()

		c.IndentedJSON(http.StatusOK, gin.H{
			"id_despacho": deliveryId,
			"id_compra":   lastId,
		})
	}
}
