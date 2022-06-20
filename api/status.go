package api

import (
	"github.com/colinmurphy1/onkyo-remote/eiscp"
	"github.com/colinmurphy1/onkyo-remote/lib"
	"github.com/gin-gonic/gin"
)

func GetStatus(c *gin.Context) {
	lib.Response(c, 200, "OK", eiscp.Conn.Status)
}
