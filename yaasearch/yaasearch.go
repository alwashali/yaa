package yamlsearch

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/blevesearch/bleve"
	_ "github.com/blevesearch/bleve/config"
	"github.com/blevesearch/bleve/document"
	"github.com/go-yaml/yaml"
)

var indexDir = "yaml_index"

func Index(dataDir string) error {

	//remove the index directory if already exists
	if indexExists(indexDir) {
		err := os.Remove(indexDir)
		if err != nil {
			fmt.Println("Error:", err)
		}
	}

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

	// Walk through the YAML files and index them
	err = filepath.Walk(dataDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() && strings.HasSuffix(strings.ToLower(path), ".yml") {

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

			// Index the YAML data
			index.Index(path, yamlData)

		}
		return nil
	})

	if err != nil {
		fmt.Printf("Error walking the directory: %v\n", err)
		return err
	}

	println("Indexing finished")

	index.Close()
	return nil

}

func Search(query []string) *bleve.SearchResult {
	// Search for a term within the index

	if indexExists(indexDir) {

		index, err := bleve.Open(indexDir)
		defer index.Close()

		queryStr := strings.Join(query, " ")
		query := bleve.NewQueryStringQuery(queryStr)

		search := bleve.NewSearchRequest(query)
		search.Highlight = bleve.NewHighlightWithStyle("ansi")

		fmt.Printf("Searching for '%s'\n", queryStr)

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

func GetDocument(id string) *document.Document {
	if indexExists(indexDir) {

		index, err := bleve.Open(indexDir)

		defer index.Close()
		if err != nil {
			fmt.Println(err)
		}
		doc, err := index.Document(id)
		return doc

	} else {
		fmt.Println("Index was not found")
	}

	return nil
}
