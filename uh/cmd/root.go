/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/powdrsoft/unified-help/uh/helpers"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
	"os/user"
	"strings"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "uh",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	//fmt.Println(os.Args)

	if len(os.Args) < 2 ||
		strings.HasPrefix(os.Args[1], "-") ||
		containsCommandName(rootCmd.Commands(), os.Args[1]) {

		if err := rootCmd.Execute(); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

	} else {
		if err := helpers.Exec("tldr", os.Args[1:]...); err != nil {
			//man takes only one arg
			if err := helpers.Exec("man", os.Args[1]); err != nil {
				if err := rootCmd.Execute(); err != nil {
					//fmt.Println(err)
					os.Exit(1)
				}
			}
		}
	}
}

func containsCommandName(s []*cobra.Command, command string) bool {
	for _, a := range s {
		if a.Name() == command {
			return true
		}
	}
	return false
}

func isSubCommand(a string, b []string) bool {
	for _, v := range b {
		if a == v {
			return true
		}
	}
	return false
}

func Map(vs []string, f func(string) string) []string {
	vsm := make([]string, len(vs))
	for i, v := range vs {
		vsm[i] = f(v)
	}
	return vsm
}

func init() {
	//cobra.OnInitialize(initConfig)
	initConfig()
	fmt.Println("READING CONFIG")
	fmt.Println(viper.Get("color"))
	fmt.Println(viper.Get("notes"))

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.uh.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	//rootCmd.AddCommand(createFromConfig("intellij"))
}

func getMDFilesFromConfigNotes(configNotes []string) {

}

func createFromConfig(name string, path string) *cobra.Command {
	usr, _ := user.Current()
	var cmd = &cobra.Command{
		Use:   name,
		Short: strings.Replace(path, usr.HomeDir, "~", 1),
		Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:` + path,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(cmd.Name())
			helpers.ExecCobraCmd(path)
		},
	}
	return cmd
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".uh" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".uh")
	}
	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())

		mdFiles := helpers.GetMDFiles(viper.GetStringSlice("notes")...)

		for key, value := range mdFiles {
			//fmt.Printf("creating %s", f.Name)
			rootCmd.AddCommand(createFromConfig(key, value))
		}
	}
}
