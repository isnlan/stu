package main

import (
	"encoding/json"
	"fmt"
	"github.com/xeipuuv/gojsonschema"
)

func main() {

	schema := `{
    "$schema": "http://json-schema.org/draft-04/schema#",
    "type": "object",
    "properties": {
        "name": {
            "type": "string"
        },
        "email": {
            "type": "string"
        },
        "age": {
            "type": "integer",
            "minimum": 0,
            "exclusiveMinimum": false
        },
        "telephone": {
            "type": "string",
            "pattern": "^(\\([0-9]{3}\\))?[0-9]{3}-[0-9]{4}$"
        }
    },
    "required": [
        "name",
        "email"
    ],
    "additionalProperties": false
}`
	schemaLoader := gojsonschema.NewStringLoader(schema)


	doc := `{
    "name": "Sherlock Holmes",
    "email": "sherlock@gmail.com",
    "age": 164
}`
	documentLoader := gojsonschema.NewStringLoader(doc)


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

	fmt.Println("------")
	sch, err := gojsonschema.NewSchema(schemaLoader)
	if err != nil {
		panic(err)
	}
	fmt.Println(sch)

	var js JsonSchema
	err = json.Unmarshal([]byte(schema), &js)
	if err != nil {
		panic(err)
	}

	fmt.Println(js)
}

type JsonSchema struct {
	Required []string `json:"required"`
}
