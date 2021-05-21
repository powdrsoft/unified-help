package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"strconv"
)

var calcCmd = &cobra.Command{
	Use:   "calc",
	Aliases: []string{"calculator"},
	Short: "A basic calculator",
	Long: `A basic calculator, 
Supports only + - * /
Does not support ( and ).`,

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(calc(args))
	},
}

func init() {
	rootCmd.AddCommand(calcCmd)
}

func calc(args []string) float64 {
	result := toNum(args[0])
	for i := 0; i < len(args)-1; i++ {
		switch args[i] {
		case "*":
			result = toNum(args[i-1]) * toNum(args[i+1])
			i++
		case "/":
			result = toNum(args[i-1]) / toNum(args[i+1])
			i++
		case "+":
			result = result + toNum(args[i+1])
			i++
		case "-":
			result = result - toNum(args[i+1])
			i++
		}
	}
	return result
}

func toNum(str string) float64 {
	f, err := strconv.ParseFloat(str, 64)
	if err != nil {
		fmt.Printf("%s is not a number\n", str)
		os.Exit(2)
	}
	return f
}
