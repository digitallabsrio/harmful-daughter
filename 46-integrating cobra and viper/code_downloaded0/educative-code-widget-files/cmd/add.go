/*
Copyright © 2020 Gigi

*/
package cmd

import (
	"fmt"
    "strconv"
    "calc-app/pkg/calc"
	"github.com/spf13/cobra"
    "github.com/spf13/viper"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
    Use:   "add",
    Short: "Add two integers",
    Long:  `Add two integers a and b; result = a + b`,
    Args:  cobra.ExactArgs(2),
    Run: func(cmd *cobra.Command, args []string) {
        var a, b int
        var err error
        a, err = strconv.Atoi(args[0])
        if err != nil {
            panic("Arguments to `add` must be integers")
        }
        b, err = strconv.Atoi(args[1])
        if err != nil {
            panic("Arguments to `add` must be integers")
        }
        
        result := calc.Add(a, b, viper.GetBool("check"))
        fmt.Println(result)
    },
}

func init() {
    addCmd.Flags().Bool(
        "check",
        false,
        "check controls if overflow/underflow check is performed")
    err := viper.BindPFlag("check", addCmd.Flags().Lookup("check"))
    if err != nil {
        panic("Unable to bind flag")
    }

    rootCmd.AddCommand(addCmd)
}
