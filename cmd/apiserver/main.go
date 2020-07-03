package main

import (
    "flag"
    "github.com/BurntSushi/toml"
    "log"

    "github.com/burkovski/http-rest-api/internal/app/apiserver"
)

var (
    configPath string
)

func init() {
    flag.StringVar(&configPath, "config-path", "configs/apiserver.toml", "path to config file")
}

func main() {
    flag.Parse()

    config := apiserver.NewConfig()
    if _, err := toml.DecodeFile(configPath, config); err != nil {
        log.Fatal(err)
    }

    server := apiserver.New(config)
    if err := server.Start(); err != nil {
        log.Fatal(err)
    }
}
