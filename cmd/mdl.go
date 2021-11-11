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

	"github.com/danxorzum/wecho/services"
	"github.com/spf13/cobra"
)

// mdlCmd represents the mdl command
var mdlCmd = &cobra.Command{
	Use:   "mdl",
	Short: "Create model files",
	Long: `Create model fiels and model folder.
use -t flag to select template:
crud	: Create crud template.
default	: Empty template`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		params, err := cmd.Flags().GetString("Template")
		if err != nil {

			fmt.Println(err)
		} else {
			switch params {
			case "crud":
				services.Srvs.CreateFile("models", args[0], params)

			case "default":
				services.Srvs.CreateFile("models", args[0], "default")
			}
		}
		fmt.Println("model created")
	},
}

func init() {
	mdlCmd.Flags().StringP("Template", "t", "default", "Select the Model templeate.")
	mkCmd.AddCommand(mdlCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// mdlCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// mdlCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
