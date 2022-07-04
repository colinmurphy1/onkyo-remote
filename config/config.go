package config

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Receiver     configReceiver `yaml:"receiver"`   // IP address or hostname of the receiver
	HTTPPort     string         `yaml:"http_port"`  // Port that the API listens on
	EnableRemote bool           `yaml:"web_remote"` // Enable the web-based remote
	EnableRaw    bool           `yaml:"api_raw"`    // Enable the /api/raw/:command endpoint
	MaxVolume    uint           `yaml:"max_volume"` // Maximum volume level
	Logging      configLogging  `yaml:"logging"`    // Logging settings
	Inputs       []configInputs `yaml:"inputs"`     // Custom inputs
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

var Conf *Config

// Configure environment variables and set defaults
func init() {
	// Command-line arguments
	configFile := flag.String("config", "", "Path to yaml config file")
	flag.Parse()

	// Require that -config be passed, and show usage
	if *configFile == "" {
		flag.Usage()
		os.Exit(1)
	}

	// Initialize Config struct
	Conf = new(Config)

	// Read yaml file
	content, err := ioutil.ReadFile(*configFile)
	if err != nil {
		fmt.Println("Could not open configuration file:", err)
		os.Exit(1)
	}

	// Load yaml config into Config struct
	err = yaml.Unmarshal(content, &Conf)

	if err != nil {
		fmt.Println("Could not parse configuration file:", err)
		os.Exit(1)
	}
}
