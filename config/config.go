package config

import (
	"flag"
	"fmt"
	"os"
)

type Config struct {
	Receiver         configReceiver `yaml:"receiver"`      // IP address or hostname of the receiver
	HTTPPort         string         `yaml:"http_port"`     // Port that the API listens on
	EnableRemote     bool           `yaml:"web_remote"`    // Enable the web-based remote
	EnableDebugTools bool           `yaml:"debug_tools"`   // Enable the debugging endpoints
	MaxVolume        uint           `yaml:"max_volume"`    // Maximum volume level
	Logging          configLogging  `yaml:"logging"`       // Logging settings
	Inputs           []configInputs `yaml:"inputs"`        // Custom inputs
	HiddenInputs     []string       `yaml:"hidden_inputs"` // Hidden inputs
}

type configReceiver struct {
	Address string `yaml:"address"`
	Port    int    `yaml:"port"`
}

type configLogging struct {
	Eiscp bool `yaml:"eiscp"` // eiscp logging
	HTTP  bool `yaml:"http"`  // http logging
}

type configInputs struct {
	Hex  string `yaml:"hex"`
	Name string `yaml:"name"`
}

// Initialize Conf struct
var Conf *Config = new(Config)

func init() {
	// Command-line arguments
	configFile := flag.String("config", "", "Path to yaml config file")
	flag.Parse()

	// Require that -config be passed, and show usage
	if *configFile == "" {
		flag.Usage()
		os.Exit(1)
	}

	// Read and parse yaml configuration file
	if err := Conf.parseConfig(*configFile); err != nil {
		fmt.Println("Error loading configuration:\n", err)
		os.Exit(1)
	}
}
