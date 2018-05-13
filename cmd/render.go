// Copyright © 2018 Guillaume Jacquet <gjacquet@upgrade.com>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"github.com/gjacquet/solarize/colors"
	"github.com/gjacquet/solarize/render"
	"github.com/spf13/cobra"
)

// renderCmd represents the render command
var renderCmd = &cobra.Command{
	Use:   "render <template> <output file>",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		templatePath := args[0]
		outputPath := args[1]

		tone, _ := cmd.Flags().GetString("tone")

		var palette colors.Palette
		if tone == colors.Dark {
			palette = colors.DarkPalette()
		} else {
			palette = colors.LightPalette()
		}

		renderer, err := render.FromFile(templatePath)
		if err != nil {
			panic(err)
		}

		err = renderer.RenderToFile(palette, outputPath)
		if err != nil {
			panic(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(renderCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// renderCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	renderCmd.Flags().StringP("tone", "t", "dark", "Set the tone. Valid values are either dark or light.")
}
