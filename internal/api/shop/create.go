package shop

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/oganes5796/shops-list/internal/model"
)

func (im *Implementation) Create(c *gin.Context) {
	var info *model.ShopInfo
	if err := c.BindJSON(&info); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "internal error",
		})
		return
	}

	id, err := im.shopService.Create(c.Request.Context(), info)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "api:Create:Service.Create",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id": id,
	})
}
