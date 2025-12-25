package shop

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	servShop "github.com/oganes5796/shops-list/internal/service/shop"
)

func (im *Implementation) GetByID(c *gin.Context) {

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil || id <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid id",
		})
		return
	}

	shop, err := im.shopService.GetByID(c.Request.Context(), id)
	if err != nil {
		if errors.Is(err, servShop.ErrShopNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "shop not found",
			})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "internal error",
		})
		return
	}

	c.JSON(http.StatusOK, shop)
}
