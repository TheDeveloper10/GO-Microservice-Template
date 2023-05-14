package config

type MasterConfig struct {
	Example bool       `json:"example"`
	HTTP    HTTPConfig `json:"http"`
}

type HTTPConfig struct {
	Address string `json:"address"`
}
