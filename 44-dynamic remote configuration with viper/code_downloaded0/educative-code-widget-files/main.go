package main

import (
    "fmt"
    "log"
    "reflect"

    "github.com/spf13/viper"
    _ "github.com/spf13/viper/remote"
    "time"
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

    currConfig := viper.AllSettings()
    fmt.Println(currConfig)

    for {
        time.Sleep(time.Second * 3)
        err = viper.WatchRemoteConfig()
        if err != nil {
            log.Fatal(err)
        }

        err = viper.ReadRemoteConfig()
        if reflect.DeepEqual(currConfig, viper.AllSettings()) {
            continue
        }

        currConfig = viper.AllSettings()
        fmt.Println(currConfig)
    }
}