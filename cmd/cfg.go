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

// cfgCmd represents the cfg command
var cfgCmd = &cobra.Command{
	Use:   "cfg",
	Short: "Create configuration files.",
	Long: `Create mysql and yml configuration files. 
Also created config folder if doesnt exist.`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// params, err := cmd.Flags().GetString("Template")
		// if err != nil {

		// fmt.Println(err)
		// } else {
		// switch params {
		switch args[0] {
		case "mysql":
			services.Srvs.CreateFile("configs", args[0], "mysql")

		case "default":
			services.Srvs.CreateFile("configs", args[0], "yml")
		}
		// }
		fmt.Println("config created")
	},
}

func init() {
	mkCmd.AddCommand(cfgCmd)
}
