package main

import (
    "fmt"
    "github.com/spf13/viper"
)

func main() {
    v1 := viper.New()
    v2 := viper.New()

    v1.SetConfigFile("config.json")
    v2.SetConfigFile("config.toml")

    err := v1.ReadInConfig()
    if err != nil {
        fmt.Println(err)
    }

    err = v2.ReadInConfig()
    if err != nil {
        fmt.Println(err)
    }

    fmt.Println(v1.AllSettings())
    fmt.Println(v2.AllSettings())
}