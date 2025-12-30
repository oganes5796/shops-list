package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/oganes5796/shops-list/internal/model"
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

	auth := router.Group("/auth")
	{
		auth.POST("/register", im.Register)
		auth.POST("/login", im.Login)
	}

	api := router.Group("/api")

	api.Use(AuthMiddleware(im.services.AuthService))

	shops := api.Group("/shops")
	{
		shops.GET("/", RoleMiddleware(model.RoleUser, model.RoleManager), im.GetAll)
		shops.GET("/:id", RoleMiddleware(model.RoleUser, model.RoleManager), im.GetByID)

		shops.POST("/", RoleMiddleware(model.RoleManager), im.Create)
		shops.PUT("/:id", RoleMiddleware(model.RoleManager), im.Update)
		shops.DELETE("/:id", RoleMiddleware(model.RoleManager), im.Delete)
	}

	return router
}
