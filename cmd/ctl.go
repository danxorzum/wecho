/*
Copyright © 2021 Miguel Angel A. Navarro <migue.danxor@gmail.com>

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

// ctlCmd represents the ctl command
var ctlCmd = &cobra.Command{
	Use:   "ctl",
	Short: "Created echo controller",
	Long: `This shit helps you to created echos controllers. 
You can make crud template or empty one.`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		params, err := cmd.Flags().GetString("Template")
		if err != nil {

			fmt.Println(err)
		} else {
			switch params {
			case "crud":
				services.Srvs.CreateFile("controllers", args[0], params)

			case "default":
				services.Srvs.CreateFile("controllers", args[0], "default")
			}
		}
		fmt.Println("controller created")
	},
}

func init() {
	ctlCmd.Flags().StringP("Template", "t", "default", "Select the controller templeate.")
	mkCmd.AddCommand(ctlCmd)
}
