package config

type Config struct {
	IsDebug *bool
	Listen  struct {
		Type   string
		BindIP string
		Port   string
	}
}

var instance *Config 
var once sync.Once 

func GetConfig() *Config 
	once sync.Once 

func 
