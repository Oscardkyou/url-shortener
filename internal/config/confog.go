package config

type Config struct {
	IsDebug *bool
	Listen  struct {
		Type   string
		BindIP string
		Port   string
	}
}
