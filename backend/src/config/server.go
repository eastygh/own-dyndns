package config

type Server struct {
	Port     int    `json:"port" yaml:"port"`
	BasePath string `json:"base-path" yaml:"base-path"`
}
