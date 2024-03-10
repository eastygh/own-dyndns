package config

type DB struct {
	Type string `json:"type" yaml:"type"`
	Url  string `json:"url" yaml:"url"`
}
