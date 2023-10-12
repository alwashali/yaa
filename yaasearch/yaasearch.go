package yamlsearch

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	bleve "github.com/blevesearch/bleve/v2"
	"github.com/blevesearch/bleve/v2/search/highlight/highlighter/ansi"
	"gopkg.in/yaml.v3"
)

var indexDir = "yaml_index"

func Index(dataDir string) error {

	// Open or create a new index
	index, err := bleve.Open(indexDir)

	if err == bleve.ErrorIndexPathDoesNotExist {

		mapping := bleve.NewIndexMapping()
		index, err = bleve.New(indexDir, mapping)
		if err != nil {
			fmt.Printf("Error creating index: %v\n", err)
			return err
		}
	} else if err != nil {
		fmt.Printf("Error opening index: %v\n", err)
		return err
	}

	stopChan := make(chan struct{})
	go showIndicatorsDots(stopChan)

	// Walk through the YAML files and index them
	err = filepath.Walk(dataDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() && (strings.HasSuffix(strings.ToLower(path), ".yml") || strings.HasSuffix(strings.ToLower(path), ".yaml")) {

			data, err := os.ReadFile(path)
			if err != nil {
				fmt.Printf("Error reading file %s: %v\n", path, err)
				return nil
			}

			// Parse the YAML data
			var yamlData map[string]interface{}

			if err := yaml.Unmarshal(data, &yamlData); err != nil {
				fmt.Printf("Error parsing YAML file %s: %v\n", path, err)
				return nil
			}

			index.Index(path, yamlData)

		}
		return nil
	})

	if err != nil {
		fmt.Printf("Error walking the directory: %v\n", err)
		return err
	}
	close(stopChan)
	println("Done!")

	index.Close()
	return nil

}

func Search(query []string, limit int) *bleve.SearchResult {
	// Search for a term within the index

	if indexExists(indexDir) {

		index, err := bleve.Open(indexDir)
		if err != nil {
			fmt.Printf("Error searching index: %v\n", err)
			return nil
		}
		defer index.Close()

		queryStr := strings.Join(query, " ")
		query := bleve.NewQueryStringQuery(queryStr)
		search := bleve.NewSearchRequest(query)
		search.Size = limit
		search.Highlight = bleve.NewHighlightWithStyle(ansi.Name)

		result, err := index.Search(search)

		if err != nil {
			fmt.Printf("Error searching index: %v\n", err)
			return nil
		}

		return result
	} else {
		fmt.Println("Index was not found")
	}

	return nil

}

func indexExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false // Folder does not exist
	}
	return false
}

func showIndicatorsDots(stopChan <-chan struct{}) {
	dots := []string{".", "..", "...", "...."}
	index := 0

	for {
		select {
		case <-stopChan:
			return
		default:
			fmt.Printf("\rIndexing%s", dots[index])
			index = (index + 1) % len(dots)
			time.Sleep(500 * time.Millisecond)
		}
	}
}
