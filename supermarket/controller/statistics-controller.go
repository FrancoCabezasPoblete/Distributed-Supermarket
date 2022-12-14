package controller

import (
	"net/http"

	"tarea1/golang-api/entity"
	"tarea1/golang-api/storage"

	"github.com/gin-gonic/gin"
)

func GetStatistics(c *gin.Context) {
	statistics := entity.Statistics{
		MostBought:  storage.GetMostBought(),
		LeastBought: storage.GetLeastBought(),
		MostProfit:  storage.GetMostProfit(),
		LeastProfit: storage.GetLeastProfit(),
	}

	c.IndentedJSON(http.StatusOK, statistics)
}
