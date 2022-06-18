package config

import (
	"fmt"
	"os"
	"strconv"
)

type Config struct {
	RECEIVER_IP   string // IP address or hostname of the receiver
	HTTP_PORT     string // Port that the API listens on
	ENABLE_REMOTE bool   // Enable the web-based remote
	EISCP_LOGGING bool   // Enable logging of eISCP commands
}

var Conf *Config

// https://stackoverflow.com/questions/40326540/how-to-assign-default-value-if-env-var-is-empty
func getEnv(key, fallback string, required bool) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	if required {
		fmt.Println("Required variable", key, "is not set. You must set this environment variable for the application to function.")
		os.Exit(1)
	}

	return fallback
}

// Configure environment variables and set defaults
func init() {
	Conf = new(Config)

	// Set values, and assign a default one if it is not passed
	Conf.RECEIVER_IP = getEnv("RECEIVER_IP", "0.0.0.0", true)
	Conf.HTTP_PORT = getEnv("HTTP_PORT", "8080", false)
	Conf.EISCP_LOGGING, _ = strconv.ParseBool(
		getEnv("ISCP_LOGGING", "true", false),
	)
	Conf.ENABLE_REMOTE, _ = strconv.ParseBool(
		getEnv("ENABLE_REMOTE", "false", false),
	)
}
