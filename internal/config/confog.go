package config

import "sync"

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
