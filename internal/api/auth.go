package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/oganes5796/shops-list/internal/model"
	"github.com/oganes5796/shops-list/internal/service/auth"
)

func (im *Implementation) Register(c *gin.Context) {
	var req struct {
		Username string     `json:"username" binding:"required"`
		Role     model.Role `json:"role" binding:"required"`
	}

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	input := &model.UserInfo{
		Username: req.Username,
		Role:     req.Role,
	}

	id, err := im.services.AuthService.Register(c.Request.Context(), input)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, gin.H{"id": id})
}

func (im *Implementation) Login(c *gin.Context) {
	var req struct {
		Username string `json:"username" binding:"required"`
	}

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := im.services.AuthService.GetUserByUsername(c.Request.Context(), req.Username)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": auth.ErrUsernameNotFound})
	}

	c.JSON(http.StatusOK, gin.H{
		"id":       user.ID,
		"username": user.Info.Username,
		"role":     user.Info.Role,
	})
}
