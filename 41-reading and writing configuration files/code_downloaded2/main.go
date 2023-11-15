package main

import (
    "fmt"
    "io/ioutil"
    "github.com/spf13/viper"
)

func main() {
    err := ioutil.WriteFile("config.toml", []byte{}, 0644)
    if err != nil {
        fmt.Println(err)
    }
    err = ioutil.WriteFile("config.json", []byte{}, 0644)
    if err != nil {
        fmt.Println(err)
    }
    
    viper.AddConfigPath(".")
    viper.Set("config-1", map[string]int{"a": 1, "b": 2})
    err = viper.WriteConfig()
    if err != nil {
        fmt.Println(err)
    }

    bytes, err := ioutil.ReadFile("config.toml")
    fmt.Println(string(bytes))

    bytes, err = ioutil.ReadFile("config.json")
    fmt.Println(string(bytes))
} 