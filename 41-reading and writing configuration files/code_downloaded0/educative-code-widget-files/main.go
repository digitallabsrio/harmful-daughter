package main

import (
    "fmt"
    "github.com/spf13/viper"
)

func main() {
    viper.SetConfigFile("config.toml")

    err := viper.ReadInConfig()
    if err != nil {
        fmt.Println(err.Error())
    } else {
        fmt.Println(viper.Sub("section-1").AllSettings())
        fmt.Println(viper.Sub("section-2").AllSettings())
    }
}