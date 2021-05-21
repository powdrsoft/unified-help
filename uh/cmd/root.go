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
	Short: "Unified Help. Source: https://github.com/powdrsoft/unified-help",
	Long: `Unifying help files written in MarkDown, from multiple locations/repositories.
Source: https://github.com/powdrsoft/unified-help`,
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
		if strings.ToLower(a.Name()) == strings.ToLower(command) {
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
		Use:  strings.ToLower(name),
		Short: strings.Replace(path, usr.HomeDir, "~", 1),
		Long:  `Source: ` + path,
		Run: func(cmd *cobra.Command, args []string) {
			//fmt.Println(cmd.Name())
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
		fmt.Printf("Config: %s\n\n", viper.ConfigFileUsed())

		mdFiles := helpers.GetMDFiles(viper.GetStringSlice("notes")...)

		for key, value := range mdFiles {
			//fmt.Printf("creating %s", f.Name)
			rootCmd.AddCommand(createFromConfig(key, value))
		}
	}
}
