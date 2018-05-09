// Copyright © 2018 Wolfy-J <wolfy.jd@gmail.com>
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package cmd

import (
	"github.com/spf13/cobra"
	"github.com/wolfy-j/goffli/ffmpeg"
	"github.com/wolfy-j/goffli/scripts"
	"github.com/wolfy-j/goffli/utils"
)

const (
	vendor      = "wolfy-j"
	application = "goffli"
)

var registry *scripts.Registry

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:           "goffli",
	SilenceErrors: true,
	Short:         utils.Sprintf("\n%s\n", makeBanner()),
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		utils.Printf("Error: <red>%s</reset>\n", err)
	}
}

func init() {
	rootCmd.PersistentFlags().BoolVarP(&utils.Verbose, "verbose", "v", false, "verbose output")

	// script registry relies on local application data
	registry = scripts.NewRegistry(vendor, application)

	// virtual Goffli commands based on copy of runCmd
	for _, script := range registry.GetAll() {
		cmd := *runCmd
		cmd.Use = script.Name
		cmd.Short = utils.Sprintf("<green+hb>%s</reset>", script.Tag("description", script.Path))

		rootCmd.AddCommand(&cmd)
	}
}

func makeBanner() *utils.Banner {
	bn := utils.NewBanner(70)

	bn.Add("<yellow+hb>Goffli</reset> <green+hb>is programmable ffmpeg wrapper</reset>")
	bn.Add("<green+hb>with ability to share snippets using <yellow+hb>GitHub Gist</reset><green+hb>.</reset>")

	if v, err := ffmpeg.Version(); err != nil {
		bn.Add("<white+hb>ffmpeg version</reset> <red>%s</reset>", err)
	} else {
		bn.Add("<white+hb>ffmpeg version</reset> <cyan+hb>%s</reset>", v)
	}

	return bn
}
