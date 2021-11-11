/*
Copyright © 2021 Miguel Angel A. Navarro <EMAIL ADDRESS>

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

	"github.com/spf13/cobra"
)

// flowCmd represents the flow command
var flowCmd = &cobra.Command{
	Use:   "flow",
	Short: "Creat entired echo flow",
	Long: `Create an entired echo flow, routes whit controller and models etc...
Currently in develop`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("working in this shit")
	},
}

func init() {
	rootCmd.AddCommand(flowCmd)
}
