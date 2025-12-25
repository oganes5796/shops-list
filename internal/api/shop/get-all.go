package shop

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (im *Implementation) GetAll(c *gin.Context) {
	lists, err := im.shopService.GetAll(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "api:GetAll:Service.GetAll",
		})
		return
	}

	c.JSON(http.StatusOK, lists)
}
