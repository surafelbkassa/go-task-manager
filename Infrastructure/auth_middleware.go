package Infrastructure

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(jwtSvc JWTServiceInterface, requiredRole string) gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := c.GetHeader("Authorization")
		if auth == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is required"})
			return
		}
		parts := strings.SplitN(auth, " ", 2)
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization header format must be Bearer {token}"})
			return
		}
		token := parts[1]
		userID, role, err := jwtSvc.ValidateToken(token)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			return
		}
		if requiredRole != "" && requiredRole != role {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Insufficient permissions"})
			return
		}
		c.Set("user_id", userID)
		c.Set("user_role", role)
		c.Next()
	}
}
