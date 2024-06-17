package config

import (
	"log"
	"os"
	"sync"

	"github.com/spf13/viper"
)

var (
	once sync.Once
	conf *Config
)

type Config struct {
	*viper.Viper
}

// GetConfig returns the configuration object.
//
// It checks if the configuration object is nil and if so, it initializes it using the Load function.
// It then returns the configuration object.
//
// Returns:
// - *Config: The configuration object.
func GetConfig() *Config {
	if conf == nil {
		once.Do(Load)
	}
	return conf
}

// Load loads the configuration from environment variables and sets default values for all environment variables.
//
// It initializes the configuration object with a new instance of viper.Viper and sets it to the global conf variable.
// The conf.AutomaticEnv() function is called to automatically read in environment variables.
// The conf.SetDefault() function is used to set default values for the "server_port" and "service_name" environment variables.
// Finally, the conf.LoadFromFile() function is called to load the configuration from a file.
func Load() {
	// Load config from env variables
	conf = &Config{viper.New()}
	conf.AutomaticEnv()

	// set defaults for all env
	// application settings
	conf.SetDefault("server_port", "8080")
	conf.SetDefault("service_name", "riskreward")

	conf.LoadFromFile()
}

// LoadFromFile loads the configuration from a file and sets default values for all environment variables.
//
// It initializes the configuration object with a new instance of viper.Viper and sets it to the global conf variable.
// The function retrieves the current working directory and appends "/../config" to it.
// It then adds the configPath to the viper config path, sets the config name to "config", and sets the config type to "yaml".
// It reads in the configuration file using the viper ReadInConfig() function.
// If there is an error reading the config file, the function returns.
// Finally, it logs the path of the config file that was loaded.
//
// Parameters:
// - conf: a pointer to a Config object representing the configuration.
//
// Return type:
// - None.
func (conf *Config) LoadFromFile() {
	v := conf.Viper
	configPath, err := os.Getwd()

	if err != nil {
		return
	}

	configPath += "/../config"

	v.AddConfigPath(configPath)
	v.SetConfigName("config")
	v.SetConfigType("yaml")

	err = v.ReadInConfig()

	if err != nil {
		return
	}

	log.Printf("Loading config from file %s", v.ConfigFileUsed())

	// Log loading config file from
}
