package configs

// Config represents the main configuration structure for the application.
// It holds all the configuration settings organized in sub-structures.
// The configuration can be loaded from various sources using mapstructure.
type (
	Config struct {
		Service  Service  `mapstructure:"service"`
		Database Database `mapstructure:"database"`
		Spotify  Spotify  `mapstructure:"spotify"`
	}

	Service struct {
		Port      string `mapstructure:"port"`
		SecretJWT string `mapstructure:"secretJWT"`
	}

	Database struct {
		DataSourcesName string `mapstructure:"dataSourcesName"`
	}
)

type Spotify struct {
	ClientID     string `mapstructure:"clientID"`
	ClientSecret string `mapstructure:"clientSecret"`
	TokenURL     string `mapstructure:"tokenURL"`
}
