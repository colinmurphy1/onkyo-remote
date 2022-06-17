package api

import (
	"github.com/colinmurphy1/onkyo-remote/eiscp"
	"github.com/colinmurphy1/onkyo-remote/help"
	"github.com/gin-gonic/gin"
)

func GetStatus(c *gin.Context) {
	help.Response(c, 200, "OK", eiscp.Conn.Status)
}
