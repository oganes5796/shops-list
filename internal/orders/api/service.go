package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/oganes5796/shops-list/internal/orders/service/serv"
)

type Implementation struct {
	services *serv.Serv
}

func NewImplementation(services *serv.Serv) *Implementation {
	return &Implementation{
		services: services,
	}
}

func (im *Implementation) InitRoutes() *gin.Engine {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "hello",
		})
	})

	api := router.Group("/api")

	orders := api.Group("/orders")
	{
		orders.GET("/:id", im.GetByID)

		orders.POST("/", im.Create)
		orders.PATCH("/:id", im.Update)
	}

	return router
}
