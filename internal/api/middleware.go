package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/oganes5796/shops-list/internal/model"
	"github.com/oganes5796/shops-list/internal/service"
)

func AuthMiddleware(authService service.AuthService) gin.HandlerFunc {
	return func(c *gin.Context) {
		username := c.GetHeader("User")
		if username == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "User header is missing"})
			return
		}

		user, err := authService.GetUserByUsername(c.Request.Context(), username)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid user"})
			return
		}

		c.Set("user", user)
		c.Set("role", user.Info.Role)

		c.Next()
	}
}

func RoleMiddleware(allowed ...model.Role) gin.HandlerFunc {
	return func(c *gin.Context) {
		role, exists := c.Get("role")
		if !exists {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "role not found in context"})
			return
		}

		for _, allowedRole := range allowed {
			if role == allowedRole {
				c.Next()
				return
			}
		}

		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "role not allowed"})
	}
}
