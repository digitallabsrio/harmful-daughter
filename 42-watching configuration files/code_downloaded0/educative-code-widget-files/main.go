package main

import (
    "fmt"
    "github.com/fsnotify/fsnotify"
    "github.com/spf13/viper"
)

func main() {
    viper.AddConfigPath(".")
    viper.SetConfigFile("config.toml")

    err := viper.ReadInConfig()
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println(viper.AllSettings())

    viper.WatchConfig()
    viper.OnConfigChange(func(e fsnotify.Event) {
        err := viper.ReadInConfig()
        if err != nil {
            fmt.Println(err)
        }

        fmt.Println(viper.AllSettings())
    })

    // Infinite loop
    select {}
}