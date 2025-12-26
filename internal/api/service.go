package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/oganes5796/shops-list/internal/service/serv"
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
	{
		shops := api.Group("/shops")
		{
			shops.POST("/", im.Create)
			shops.GET("/", im.GetAll)
			shops.GET("/:id", im.GetByID)
			shops.PUT("/:id", im.Update)
			shops.DELETE("/:id", im.Delete)
		}
	}

	return router
}
