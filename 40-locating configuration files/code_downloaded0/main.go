package main

import (
    "fmt"
    "github.com/spf13/viper"
)

func main() {
    viper.AddConfigPath(".")
    viper.AddConfigPath("~/test-config")
    viper.AddConfigPath("/no-such-dir")
    viper.AddConfigPath("/tmp/test-config")

    viper.SetConfigFile("Configfile")
    viper.SetConfigType("toml")

    fmt.Println("Config path and files set!")
}