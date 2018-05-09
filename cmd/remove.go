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
	"github.com/wolfy-j/goffli/utils"
)

func init() {
	rootCmd.AddCommand(&cobra.Command{
		Use:   "remove [script]",
		Short: "Remove goffli script",
		Args:  cobra.ExactArgs(1),
		Run:   removeHandler,
	})
}

func removeHandler(cmd *cobra.Command, args []string) {
	utils.Printf("<white+hb>Removing script<reset> <cyan>%s<reset> ", args[0])

	if err := registry.Remove(args[0]); err != nil {
		utils.Printf("- <red>error:<reset> <red+hb>%s<reset>\n", err)
	} else {
		utils.Printf("- <green>done<reset>\n")
	}
}
