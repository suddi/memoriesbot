package config

// AppConfig - structure of app configuration
type AppConfig struct {
	Name    string
	Version string
}

// AuthConfig - structure of auth configuration
type AuthConfig struct {
	GooglePhotos GooglePhotosAuthConfig
}

// AwsConfig - structure of AWS configuration
type AwsConfig struct {
	Region string
}

// GooglePhotosAuthConfig - structure of Google Photos auth configuration
type GooglePhotosAuthConfig struct {
	ClientID     string
	ClientSecret string
}

// Config - structure of configuration
type Config struct {
	App  AppConfig
	Auth AuthConfig
	Aws  AwsConfig
}

// Get - get config from passed environment variables
func Get() *Config {
	return &Config{
		App: AppConfig{
			Name:    "memoriesbot",
			Version: getEnv("VERSION", "1.0.0"),
		},
		Auth: AuthConfig{
			GooglePhotos: GooglePhotosAuthConfig{
				ClientID:     getEnv("CLIENT_ID", ""),
				ClientSecret: getEnv("CLIENT_SECRET", ""),
			},
		},
		Aws: AwsConfig{
			Region: getEnv("AWS_REGION", "ap-northeast-1"),
		},
	}
}
