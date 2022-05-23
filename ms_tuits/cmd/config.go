package cmd

//Config - Struct with all the configurations
type Config struct {
	Environment string `arg:"env:ENVIRONMENT"`
	ServerConfig
}

//ServerConfig - Struct with configurations related to Server
type ServerConfig struct {
	Port string `arg:"env:SERVER_PORT"`
	Name string `arg:"env:SERVER_NAME"`
}

//DefaultConfiguration - returns a configuration object to be used in the service
func DefaultConfiguration() *Config {
	return &Config{
		Environment: "dev",
		ServerConfig: ServerConfig{
			Name: "tuits-service",
			Port: "8000",
		},
	}
}
