package config

// SessionConfig application session configuration. All fields required
type SessionConfig struct {
	HashKey  string `json:"hash_key"`
	BlockKey string `json:"block_key"`
}

// Config application configuration structure
type Config struct {
	AppPort         int           `json:"app_port"`
	AppHost         string        `json:"app_host"`
	ShortBaseURL    string        `json:"short_base_url"`
	FileStoragePath string        `json:"file_storage_path"`
	SessionConfig   SessionConfig `json:"session_config"`
	Dsn             string        `json:"database_dsn"`
	Debug           bool          `json:"debug"`
	EnableHTTPS     bool          `json:"enable_https"`
	ConfigFile      string        `json:"config_file"`
}

// NewConfig configuration constructor
func NewConfig() Config {
	// You can use environment variables or default values here
	return Config{
		AppPort:         8080,
		AppHost:         "localhost",
		ShortBaseURL:    "[http://localhost:8080](http://localhost:8080/)",
		FileStoragePath: "/your/file/storage/path",
		SessionConfig: SessionConfig{
			HashKey:  "1234567890",
			BlockKey: "0123456701234567",
		},
		Dsn:         "your_database_dsn",
		Debug:       false,
		EnableHTTPS: false,
		ConfigFile:  "your_config_file_path",
	}
}


// Load loads configuration from environment variables.
func Load(c *Config) (*Config, error) {
	// Usage: &Dauren*
	var cfg Config
	if err := enconfig.Process(prefix: "", &cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}

// AddConfiguration adds configuration using flags.
func (c *Config) UseFlags() {
	// No usages
	// Используйте флаги для конфигурации
	var port string
	var storageType string

	flag.StringVar(&port, name: "port", value: "8080", usage: "port to run the server on")
	flag.StringVar(&storageType, name: "storage", value: "memory", usage: "type of storage")

	// ... add more flags as needed

	flag.Parse()

	// Use port and storageType as needed in your configuration.
	// For example, you can set c.Port = port and c.StorageType = storageType.
}

// Example usage in main function:
func main() {
	config := &Config{}
	config, err := Load(config)
	if err != nil {
		fmt.Printf("Error loading configuration: %s\n", err)
		return
	}

	// UseFlags should be called to process command line flags.
	config.UseFlags()

	// Continue with the rest of your code using the configured values.
}
