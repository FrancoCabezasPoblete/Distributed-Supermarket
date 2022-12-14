package main

import (
	"strconv"

	"github.com/gin-gonic/gin"

	"tarea2/delivery/controller"
	"tarea2/delivery/storage"
)

const SERVER_PORT = 5000

var SERVER_HOST = ":" + strconv.Itoa(SERVER_PORT)

func main() {

	DB := storage.ConnectDatabase()

	initAPI()

	DB.Close()

}

func initAPI() {

	router := gin.Default()

	router.GET("api/clientes/estado_despacho/:id", controller.GetEstado)

	router.POST("api/supermercado/despacho", controller.PostRegister)

	router.Run(SERVER_HOST)

}
