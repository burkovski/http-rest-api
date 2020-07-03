package apiserver

import "github.com/burkovski/http-rest-api/internal/store"

type Config struct {
    BindAddr string `toml:"bind_addr"`
    LogLevel string `toml:"log_level"`
    Store    *store.Config
}

func NewConfig() *Config {
    return &Config{
        BindAddr: "localhost:8080",
        LogLevel: "debug",
        Store:    store.NewConfig(),
    }
}
