package cfg

import "os"

type Config struct {
	ApiKey string `json:"api_key"`
}

var cfg Config

func GetConfig() Config {
	return cfg
}

func init() {
	var res bool
	cfg.ApiKey, res = os.LookupEnv("TG_KEY")
	if !res {
		panic("Variable TG_KEY is not defined.")
	}
}
