// Package configs provides configuration management functionality
package configs

import "github.com/spf13/viper"

// config holds the global configuration instance
var config *Config

// option defines the configuration options structure
type option struct {
	configFolders []string // List of folders to search for config files
	configFile    string   // Name of the config file without extension
	configType    string   // Type/format of the config file (e.g., yaml, json)
}

// Init initializes the configuration system with optional settings.
// It performs the following steps:
// 1. Creates default options for config folders, file, and type
// 2. Applies any provided option functions to customize the configuration
// 3. Sets up Viper with the configured paths and settings
// 4. Reads and unmarshals the configuration into a global Config struct
//
// Parameters:
//   - opts ...Option: Variadic parameter of Option functions to customize configuration
//
// Returns:
//   - error: Returns an error if configuration reading or unmarshaling fails
//
// Usage:
//
//	err := Init(
//	    WithConfigFile("custom"),
//	    WithConfigType("yaml"),
//	)
func Init(opts ...Option) error {
	// Initialize with default options
	opt := option{
		configFolders: getDefaultConfigFolders(),
		configFile:    getDefaultConfigFile(),
		configType:    getDefaultConfigType(),
	}

	// Apply custom options if provided
	for _, optsFunc := range opts {
		optsFunc(&opt)
	}

	// Configure Viper with the specified paths
	for _, folder := range opt.configFolders {
		viper.AddConfigPath(folder)
	}

	// Set up Viper configuration
	viper.SetConfigName(opt.configFile)
	viper.SetConfigType(opt.configType)
	viper.AutomaticEnv()

	// Initialize the global config
	config = &Config{}

	// Read the configuration file
	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	// Unmarshal the configuration into the global config struct
	return viper.Unmarshal(config)
}

// Option is a function type that modifies the option struct
type Option func(*option)

// getDefaultConfigFolders returns the default configuration folders
func getDefaultConfigFolders() []string {
	return []string{"./configs"}
}

// getDefaultConfigFile returns the default configuration file name
func getDefaultConfigFile() string {
	return "config"
}

// getDefaultConfigType returns the default configuration file type
func getDefaultConfigType() string {
	return "yaml"
}

// WithConfigFolders creates an Option to set custom configuration folders
func WithConfigFolders(configFolders []string) Option {
	return func(o *option) {
		o.configFolders = configFolders
	}
}

// WithConfigFile creates an Option to set a custom configuration file name
func WithConfigFile(configFile string) Option {
	return func(o *option) {
		o.configFile = configFile
	}
}

// WithConfigType creates an Option to set a custom configuration file type
func WithConfigType(configType string) Option {
	return func(o *option) {
		o.configType = configType
	}
}

// GetConfig returns the current configuration instance
// If config is nil, returns an empty Config struct
func GetConfig() *Config {
	if config == nil {
		return &Config{}
	}
	return config
}
