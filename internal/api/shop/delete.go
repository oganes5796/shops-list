package shop

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (im *Implementation) Delete(c *gin.Context) {
	idUrl := c.Param("id")
	id, err := strconv.ParseInt(idUrl, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "api:Delete:ParseInt",
		})
		return
	}

	err = im.shopService.Delete(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "api:Delete:shopService:Delete",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
}
