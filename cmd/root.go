/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

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
	"os"
	"path/filepath"

	"github.com/spf13/cobra"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "wecho",
	Short: "Cli to develop echo apps faster.",
	Long: `Wecho is a CLI tool to make things easier.
You can:
Initialize project.
Create files.
Create flowas(in development).
	`,
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	cobra.OnInitialize(initConfig)
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	home, err := homedir.Dir()
	fmt.Println(home)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Search config in home directory with name ".cobra" (without extension).
		viper.AddConfigPath("./config")
		viper.SetConfigType("yaml")
		viper.SetConfigName(".cobra")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}

	path, _ := filepath.Abs("")

	viper.SetDefault("cmddir", path)
}

// func initConfig() {
// 	home, err := homedir.Dir()
// 	fmt.Println(home)
// 	if err != nil {
// 		fmt.Println(err)
// 		os.Exit(1)
// 	}

// 	viper.AddConfigPath(home)
// 	viper.AddConfigPath("./config")
// 	viper.SetConfigName(".wecho")

// 	viper.AutomaticEnv()
// 	file := home + "/.wecho"

// 	viper.WriteConfig()
// 	if err := viper.ReadInConfig(); err != nil {
// 		log.Fatalf("Error reading config file, %s", err)
// 	}

// 	if viper.GetString("default.templatefolder") == "" {
// 		fmt.Println("falta templatefolder en el archivo " + file)
// 		os.Exit(1)
// 	} else {
// 		viper.SetDefault("homedir", home+"/"+viper.GetString("default.templatefolder"))
// 	}
// }
