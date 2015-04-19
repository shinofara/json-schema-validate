package main

import (
	"fmt"
	"github.com/xeipuuv/gojsonschema"
	"path/filepath"
)

func main() {

	schema_path, _ := filepath.Abs("schema.json")
	data_path, _ := filepath.Abs("data.json")

	schemaLoader := gojsonschema.NewReferenceLoader("file://" + schema_path)
	documentLoader := gojsonschema.NewReferenceLoader("file://" + data_path)

	result, err := gojsonschema.Validate(schemaLoader, documentLoader)
	if err != nil {
		panic(err.Error())
	}

	if result.Valid() {
		fmt.Printf("The document is valid\n")
	} else {
		fmt.Printf("The document is not valid. see errors :\n")
		for _, desc := range result.Errors() {
			fmt.Printf("- %s\n", desc)
		}
	}

}
