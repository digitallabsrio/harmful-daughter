package main

import (
    "fmt"
    "github.com/spf13/viper"
)

func main() {
    viper.Set("bool", true)
    viper.Set("int", 777)

    fmt.Println(viper.GetBool("bool"))
    fmt.Println(viper.GetInt("int"))
}