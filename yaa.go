package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	yaasearch "yaa/yaasearch"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "Yaa",
		Usage: "Yaml Search for Humans",

		Commands: []*cli.Command{
			{
				Name:      "search",
				Aliases:   []string{"s"},
				UsageText: "Yaa search [options] SearchQuery",
				Flags: []cli.Flag{
					&cli.IntFlag{
						Name:    "limit",
						Aliases: []string{"l"},
						Value:   10,
						Usage:   "Number of results to display",
					},
					&cli.StringFlag{
						Name:    "export",
						Aliases: []string{"e"},
						Usage:   "Path to save yaml files",
					},
				},

				Action: searchAction,
			},
			{
				Name:    "index",
				Aliases: []string{"i"},
				Usage:   "Path to yaml folder",
				Action:  indexAction,
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		fmt.Println(err)
	}
}

func searchAction(c *cli.Context) error {

	query := c.Args().Slice()
	if len(query) == 0 {

		return cli.Exit("No query was found, use -h for hlep.", 1)
	}

	limit := c.Int("limit")

	results := yaasearch.Search(query, limit)

	if results.Hits.Len() > 0 {
		if c.IsSet("export") {

			dest_path := c.String("export")

			_, err := os.Stat(dest_path)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			// ID is set to the file path
			count := 0
			for _, hit := range results.Hits {

				err := exportFile(hit.ID, dest_path)
				if err != nil {
					fmt.Print(err)
				}
				count++
				//fmt.Printf("Exporting %s\n", hit.ID)
			}
			fmt.Println(count, "files exported.")
			return nil
		}

		fmt.Println(results)
	} else {

		fmt.Println("No Match Found")
	}

	return nil
}

func indexAction(c *cli.Context) error {

	path := c.Args().First()

	if path == "" {

		return cli.Exit("Please provide a folder to index", 1)
	}

	yaasearch.Index(path)

	return nil
}

// copy the file to destination path specified in the export option
func exportFile(srcFilePath, destPath string) error {

	filename := filepath.Base(srcFilePath)

	dest_file := filepath.Join(destPath, filename)

	srcFile, err := os.Open(srcFilePath)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	destFile, err := os.Create(dest_file)
	if err != nil {
		return err
	}
	defer destFile.Close()

	buffer := make([]byte, 8192) // 8KB buffer size (adjust as needed)

	_, err = io.CopyBuffer(destFile, srcFile, buffer)
	if err != nil {
		return err
	}

	err = destFile.Sync()
	if err != nil {
		return err
	}

	return nil
}
