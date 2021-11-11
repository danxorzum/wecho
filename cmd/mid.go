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

// midCmd represents the mid command
var midCmd = &cobra.Command{
	Use:   "mid",
	Short: "Create middlewares.",
	Long: `Create middlewares files an middleware folder from template.
use -t to select the template,
default is "default".
crud to create CRUD middleware.
	`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		params, err := cmd.Flags().GetString("Template")
		if err != nil {

			fmt.Println(err)
		} else {
			switch params {
			case "crud":
				services.Srvs.CreateFile("middlewares", args[0], "default")

			case "default":
				services.Srvs.CreateFile("middlewares", args[0], "default")
			}
		}
		fmt.Println("controller created")
	},
}

func init() {
	midCmd.Flags().StringP("Template", "t", "default", "Select the middleware templeate.")
	mkCmd.AddCommand(midCmd)
}
