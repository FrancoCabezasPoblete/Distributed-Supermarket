package main

import (
	"log"
	"strconv"

	"tarea1/golang-api/controller"
	products "tarea1/golang-api/protobuf"
	"tarea1/golang-api/storage"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

const (
	SERVER_PORT = 5000
	GRPC_ADDR   = "10.10.11.43:9000"
)

var SERVER_HOST = ":" + strconv.Itoa(SERVER_PORT)

func main() {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(GRPC_ADDR, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %s", err)
	}

	client := products.NewProductsClient(conn)

	DB := storage.ConnectDatabase()
	InitAPI(client)
	conn.Close()
	DB.Close()
}

func InitAPI(client products.ProductsClient) {
	router := gin.Default()

	router.POST("/api/clientes/iniciar_sesion", controller.LogInClient)

	router.GET("/api/productos", controller.GetProducts)
	router.POST("/api/productos", controller.AddProduct)

	router.GET("/api/productos/:id", controller.GetProduct)
	router.DELETE("/api/productos/:id", controller.DeleteProduct)
	router.PUT("/api/productos/:id", controller.UpdateProduct)

	router.POST("/api/compras", controller.AddShopping)
	router.GET("/api/estadisticas", controller.GetStatistics)

	router.POST("/api/proveedor", controller.RequestProduct(client))

	router.Run(SERVER_HOST)
}
