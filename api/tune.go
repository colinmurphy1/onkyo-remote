package api

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/colinmurphy1/onkyo-remote/eiscp"
	"github.com/colinmurphy1/onkyo-remote/lib"
	"github.com/gin-gonic/gin"
)

// Tunes to the specified preset
func SetTunerPreset(c *gin.Context) {
	presetId := c.Param("preset")

	// Preset must be two chars in length
	if len(presetId) != 2 {
		lib.Response(c, http.StatusBadRequest, "Preset ID must be exactly 2 characters in length", nil)
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

// Tunes to the specified frequency
func SetTunerFrequency(c *gin.Context) {
	freq := c.Param("freq")

	// Remove "." from frequency
	freq = strings.Replace(freq, ".", "", 1)

	if len(freq) != 5 {
		lib.Response(c, http.StatusBadRequest, "Frequency must be exactly 5 characters in length", nil)
		return
	}

	// Frequency must be an number (period is ok, as it was removed previously)
	if _, err := strconv.Atoi(freq); err != nil {
		lib.Response(c, http.StatusBadRequest, "Specified frequency is not an integer", nil)
		return
	}

	// Attempt to tune to frequency
	// We do not need to determine if the frequency is an AM or FM frequency; the receiver will automatically determine this
	if err := eiscp.Conn.SendCmd(fmt.Sprintf("TUN%s", freq)); err != nil {
		lib.Response(c, http.StatusOK, "Error tuning to frequency", nil)
		return
	}

	// TODO: Implement a feature to look for N/A responses in the watcher so we can display an error message if an invalid frequency is passed

	// Return 200 OK response
	lib.Response(c, http.StatusOK, "OK", nil)
}
