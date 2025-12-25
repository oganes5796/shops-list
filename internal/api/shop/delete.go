package shop

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	servShop "github.com/oganes5796/shops-list/internal/service/shop"
)

func (im *Implementation) Delete(c *gin.Context) {
	idUrl := c.Param("id")
	id, err := strconv.ParseInt(idUrl, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "internal error",
		})
		return
	}

	err = im.shopService.Delete(c.Request.Context(), id)
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

	c.Status(http.StatusNoContent)
}
