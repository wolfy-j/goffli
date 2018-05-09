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
	"context"
	"errors"
	"github.com/google/go-github/github"
	"github.com/spf13/cobra"
	"github.com/wolfy-j/goffli/utils"
	"regexp"
)

func init() {
	rootCmd.AddCommand(&cobra.Command{
		Use:   "get [url] [name]",
		Short: "Install/update goffli script using public GitHub Gist url",
		Args:  cobra.ExactArgs(2),
		Run:   getHandler,
	})
}

func getHandler(cmd *cobra.Command, args []string) {
	utils.Printf("<white+hb>Downloading</reset> <cyan>%s</reset> ", args[0])

	gistID, err := findID(args[0])
	if err != nil {
		utils.Printf("- <red>%s</reset>\n", err)
		return
	}

	client := github.NewClient(nil)
	gist, _, err := client.Gists.Get(context.Background(), gistID)
	if err != nil {
		utils.Printf("- <red>error</reset> <red+hb>%s</reset>\n", err)
		return
	}

	var body string
	for _, file := range gist.Files {
		if *file.Language == "Lua" {
			body = *file.Content
			utils.Printf(
				"- <white+hb>found</reset> <cyan>%s</reset> by <yellow+hb>%s</reset>\n",
				*file.Filename,
				*gist.Owner.Login,
			)

			break
		}
	}

	if body == "" {
		utils.Printf("- <red>error</reset> <red+hb>%s</reset>\n", "unable to find lua script")
		return
	}

	if script := registry.Get(args[1]); script != nil {
		utils.Printf(
			"<white+hb>Updating script</reset> <yellow+hb>%s</reset> <green+hb>%s</reset> ",
			script.Name,
			script.Tag("version", "undefined"),
		)

		if err := registry.Register(args[1], body); err != nil {
			utils.Printf("- <red>error</reset> <red+hb>%s</reset>\n", err)
		}

		utils.Printf(
			"- <white+hb>new version</reset> <green+hb>%s</reset>\n",
			registry.Get(args[1]).Tag("version", "undefined"),
		)
	} else {
		registry.Register(args[1], body)

		if script := registry.Get(args[1]); script != nil {
			utils.Printf(
				"<white+hb>New script</reset> <yellow+hb>%s</reset> <green+hb>%s</reset> %s\n",
				script.Name,
				script.Tag("version", "undefined"),
				script.Tag("description", ""),
			)
		}
	}
}

func findID(url string) (string, error) {
	r, err := regexp.Compile(`gist\.github\.com/([^/]+)/(\S+)`)
	if err != nil {
		panic(err)
	}

	if !r.MatchString(url) {
		return "", errors.New("invalid valid Gist Url")
	}

	matches := r.FindStringSubmatch(url)
	return matches[2], nil
}
