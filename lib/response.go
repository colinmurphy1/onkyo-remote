package lib

import "github.com/gin-gonic/gin"

// API response, specify response code and data
func Response(c *gin.Context, httpCode int, message string, data interface{}) {
	c.IndentedJSON(httpCode, gin.H{
		"message": message,
		"data":    data,
	})
}
