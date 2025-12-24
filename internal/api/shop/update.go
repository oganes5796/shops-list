package shop

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/oganes5796/shops-list/internal/model"
)

func (im *Implementation) Update(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	var info model.ShopInfo
	if err := c.BindJSON(&info); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid body"})
		return
	}

	err = im.shopService.Update(c.Request.Context(), id, &info)
	if err != nil {
		if errors.Is(err, model.ErrShopNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "shop not found"})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal error"})
		return
	}

	c.Status(http.StatusNoContent)
}
