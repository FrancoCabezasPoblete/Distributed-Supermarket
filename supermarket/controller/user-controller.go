package controller

import (
	"fmt"
	"net/http"
	"tarea1/golang-api/entity"
	"tarea1/golang-api/storage"

	"github.com/gin-gonic/gin"
)

func LogInClient(c *gin.Context) {
	var client entity.Client

	if err := c.BindJSON(&client); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		fmt.Println(err)
	} else {
		isValid := storage.ValidateSession(client)

		c.IndentedJSON(http.StatusOK, gin.H{
			"acceso_valido": isValid,
		})
	}
}
