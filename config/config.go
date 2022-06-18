package config

import (
	"os"
	"strconv"
)

type Config struct {
	RECEIVER_IP    string // IP address or hostname of the receiver
	HTTP_PORT      string // Port that the API and control panel listens on
	ENABLE_REMOTE  bool   // Enable the web-based control panel
	ENABLE_LOGGING bool   // Enable logging of eISCP commands
}

var Conf *Config

// https://stackoverflow.com/questions/40326540/how-to-assign-default-value-if-env-var-is-empty
func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

// Configure environment variables and set defaults
func init() {
	Conf = new(Config)

	// Set values, and assign a default one if it is not passed
	Conf.RECEIVER_IP = getEnv("RECEIVER_IP", "0.0.0.0")
	Conf.HTTP_PORT = getEnv("HTTP_PORT", "8080")
	Conf.ENABLE_LOGGING, _ = strconv.ParseBool(
		getEnv("ENABLE_LOGGING", "true"),
	)
	Conf.ENABLE_REMOTE, _ = strconv.ParseBool(
		getEnv("ENABLE_REMOTE", "true"),
	)
}
