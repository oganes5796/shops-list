package api

import (
	"errors"
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/oganes5796/shops-list/internal/orders/model"
	"github.com/oganes5796/shops-list/internal/orders/service/orders"
)

func (im *Implementation) Create(c *gin.Context) {
	var req model.CreateOrderRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request body",
		})
		return
	}

	if len(req.Cart) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "cart is empty",
		})
		return
	}

	id, err := im.services.OrdersService.Create(c.Request.Context(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to create order",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id": id.String(),
	})
}

func (im *Implementation) GetByID(c *gin.Context) {
	idParam := c.Param("id")

	orderID, err := uuid.Parse(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid order id",
		})
		return
	}

	order, err := im.services.OrdersService.GetByID(c.Request.Context(), orderID)
	if err != nil {
		if errors.Is(err, orders.ErrOrderNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "order not found",
			})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "internal error",
		})
		return
	}

	c.JSON(http.StatusOK, order)
}

func (im *Implementation) Update(c *gin.Context) {
	idParam := c.Param("id")

	orderID, err := uuid.Parse(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid order id",
		})
		return
	}

	var req model.UpdateOrderRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request body",
		})
		return
	}

	err = im.services.OrdersService.Update(c.Request.Context(), orderID, &req)
	if err != nil {
		if errors.Is(err, orders.ErrOrderNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "order not found",
			})
			return
		}

		if errors.Is(err, orders.ErrInvalidStatusTransition) {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "invalid order status transition",
			})
			return
		}
		slog.Error("update: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "internal error",
		})
		return
	}
	c.Status(http.StatusNoContent)
}
