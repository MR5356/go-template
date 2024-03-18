package ginutil

import "github.com/gin-gonic/gin"

func GetId(c *gin.Context) (id string) {
	return c.Query("id")
}

func GetToken(c *gin.Context) string {
	tokenString := c.GetHeader("Authorization")
	if len(tokenString) == 0 {
		tokenString, _ = c.Cookie("token")
	}

	return tokenString
}
