package access

import "github.com/elanticrypt0/go4it"

type AccessConfig struct {
	IsEnabled bool   `json:"is_enabled" toml:"is_enabled"`
	BaseURL   string `json:"base_url" toml:"base_url"`
}

func LoadConfig(config *AccessConfig) {
	go4it.ReadAndParseToml("./config/access.toml", &config)
}
