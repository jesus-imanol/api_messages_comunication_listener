package services

import (
	"net/http"
	"strings"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func RoleMiddleware(secretKey string, expectedRoles []string) gin.HandlerFunc {
    return func(c *gin.Context) {
        token := c.GetHeader("Authorization")
        if token == "" {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header missing"})
            c.Abort()
            return
        }
        token = strings.TrimPrefix(token, "Bearer ")
        parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
            return []byte(secretKey), nil
        })
        if err != nil || !parsedToken.Valid {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
            c.Abort()
            return
        }

        claims, ok := parsedToken.Claims.(jwt.MapClaims)
        if !ok {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid claims"})
            c.Abort()
            return
        }
        role, ok := claims["role"].(string)
        if !ok {
            c.JSON(http.StatusForbidden, gin.H{"error": "Access denied"})
            c.Abort()
            return
        }

        isRoleAllowed := false
        for _, expectedRole := range expectedRoles {
            if role == expectedRole {
                isRoleAllowed = true
                break
            }
        }
        
        if !isRoleAllowed {
            c.JSON(http.StatusForbidden, gin.H{"error": "Access denied"})
            c.Abort()
            return
        }

        c.Next()
    }
}