package shop

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/oganes5796/shops-list/internal/service"
)

type Implementation struct {
	shopService service.ShopService
}

func NewImplementation(shopService service.ShopService) *Implementation {
	return &Implementation{
		shopService: shopService,
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
			shops.GET("/")
			shops.GET("/:id", im.GetByID)
			shops.PUT("/:id")
			shops.DELETE("/:id", im.Delete)
		}
	}

	return router
}
