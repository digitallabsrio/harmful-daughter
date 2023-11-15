package main

import (
    "fmt"
    "log"

    "github.com/spf13/viper"
    _ "github.com/spf13/viper/remote"
)

func main() {
    err := viper.AddRemoteProvider("etcd", "http://127.0.0.1:4001", "/config/config.toml")
    if err != nil {
        log.Fatal(err)
    }

    viper.SetConfigType("toml")
    err = viper.ReadRemoteConfig()
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println(viper.AllSettings())
}