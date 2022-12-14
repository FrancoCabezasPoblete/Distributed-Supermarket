package controller

import (
	"bytes"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"tarea2/delivery/entity"
	"tarea2/delivery/rabbitMQ"
	"tarea2/delivery/storage"
)

func GetEstado(c *gin.Context) {

	// get the id from the client
	IdDespacho, err := strconv.Atoi(c.Param("id"))

	if err != nil {

		panic(err)

	}

	estado := storage.AddEstado(IdDespacho)

	if estado == nil {

		c.AbortWithStatus(http.StatusNotFound)

	} else {

		c.IndentedJSON(http.StatusOK, estado)

	}

}

func PostRegister(c *gin.Context) {

	var registro entity.Registro

	// create registro with it's corresponding struct
	if err := c.BindJSON(&registro); err != nil {

		c.AbortWithStatus(http.StatusBadRequest)

	} else {

		values, _ := json.Marshal(entity.Registro{
			IdDespacho: registro.IdDespacho,
			Estado:     "RECIBIDO",
			IdCompra:   registro.IdCompra,
		})

		registroJSON := bytes.NewBuffer(values)

		rabbitMQ.RabbitMQ(registroJSON)

		c.IndentedJSON(http.StatusCreated, gin.H{
			"id_despacho": registro.IdDespacho,
		})
	}
}
