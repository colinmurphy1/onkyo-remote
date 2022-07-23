package api

import (
	"fmt"
	"net/http"

	"github.com/colinmurphy1/onkyo-remote/eiscp"
	"github.com/colinmurphy1/onkyo-remote/lib"
	"github.com/gin-gonic/gin"
)

// Tunes to the specified preset
func SetTunerPreset(c *gin.Context) {
	presetId := c.Param("preset")

	// Preset must be two chars in length
	if len(presetId) != 2 {
		lib.Response(c, http.StatusBadRequest, "Preset ID be exactly two characters in length", nil)
		return
	}

	// Determine if preset is configured
	if _, ok := eiscp.Conn.Status.Tuner.PresetList[presetId]; !ok {
		lib.Response(c, http.StatusBadRequest, "Invalid preset specified", nil)
		return
	}

	// Send PRS command to receiver to change preset
	if err := eiscp.Conn.SendCmd("PRS" + presetId); err != nil {
		lib.Response(c, http.StatusInternalServerError, "Could not tune to preset", err)
		return
	}

	lib.Response(c, http.StatusOK, fmt.Sprintf("Tuned to preset %s", presetId), nil)
}
