package shop

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/oganes5796/shops-list/internal/model"
)

func (im *Implementation) Update(c *gin.Context) {
	idStr := c.Param("id")
	id, _ := strconv.ParseInt(idStr, 10, 64)

	var info model.ShopInfo
	if err := c.BindJSON(&info); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "api:update:BindJSON",
		})
		return
	}

	err := im.shopService.Update(c.Request.Context(), id, &info)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "api:update:shopService.Update",
		})
		return
	}

	c.Status(http.StatusNoContent)
}
