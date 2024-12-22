package util

import (
	"encoding/json"
	"os"
)

type Config struct {
	SessionCookie string `json:"session_cookie"`
}

func NewConfig(p string) (*Config, error) {
	f, err := os.Open(p)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var conf Config
	decoder := json.NewDecoder(f)
	err = decoder.Decode(&conf)
	if err != nil {
		return nil, err
	}
	return &conf, nil
}
