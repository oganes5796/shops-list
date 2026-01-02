package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/oganes5796/shops-list/internal/orders/model"
)

func (im *Implementation) GetByID(c *gin.Context) {
	idOrder := c.Param("id")

	order, err := im.services.OrdersService.GetByID(c.Request.Context(), idOrder)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, order)
}

func (im *Implementation) Create(c *gin.Context) {
	var orderInfo model.OrderInfo
	if err := c.ShouldBindJSON(&orderInfo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	idOrder, err := im.services.OrdersService.Create(c.Request.Context(), &orderInfo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"id": idOrder})
}

func (im *Implementation) Update(c *gin.Context) {
	idOrder := c.Param("id")

	var orderInfo model.OrderInfo
	if err := c.ShouldBindJSON(&orderInfo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := im.services.OrdersService.Update(c.Request.Context(), idOrder, &orderInfo); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}
