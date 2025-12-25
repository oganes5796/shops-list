package shop

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (im *Implementation) GetByID(c *gin.Context) {
	idUrl := c.Param("id")
	if idUrl == "" || idUrl == " " {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "id parameter is required",
		})
		return
	}

	id, _ := strconv.ParseInt(idUrl, 10, 64)
	if id <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id must be positive"})
		return
	}

	shop, err := im.shopService.GetByID(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "api:GetByID:shopService.GetByID: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, shop)
}
