package api

import (
	"net/http"

	"github.com/colinmurphy1/onkyo-remote/eiscp"
	"github.com/colinmurphy1/onkyo-remote/lib"
	"github.com/gin-gonic/gin"
)

// Tunes to the specified preset
func SetTunerPreset(c *gin.Context) {
	presetId := c.Param("preset")

	// Determine if preset is valid (read xml data put into presets struct)

	// Send PRS command to receiver to change preset
	if err := eiscp.Conn.SendCmd("PRS" + presetId); err != nil {
		lib.Response(c,
			http.StatusInternalServerError,
			"Could not tune to preset",
			err,
		)
	}

	lib.Response(c, http.StatusOK, "OK", presetId)
}
