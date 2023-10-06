package main

import (
	"fmt"
	"os"

	yaasearch "yaa/yaasearch"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "Yaa",
		Usage: "Yaml Searach for Humans",

		Flags: []cli.Flag{
			&cli.IntFlag{
				Name:  "limit",
				Value: 10,
				Usage: "Number of results to display",
			},
		},

		Commands: []*cli.Command{
			{
				Name:    "search",
				Aliases: []string{"s"},
				Usage:   "Search for sigma rules",

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
		return cli.Exit("Please provide a search query", 1)
	}
	limit := c.Int("limit")

	results := yaasearch.Search(query, limit)
	if results != nil {

		fmt.Println(results)
		fmt.Println()

	} else {
		fmt.Println("Error: Search was not successful")
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
