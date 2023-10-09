package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"text/template"
)

// Define a struct with a Mode field
type TemplateData struct {
	Mode string
}

func main() {
	modes := map[string]string{
		"prod": "index.html",
		"dev":  "index-dev.html",
	}
	for mode, output := range modes {
		filePath := "_index.html"
		content, err := ioutil.ReadFile(filePath)
		if err != nil {
			fmt.Println("Error reading file:", err)
			os.Exit(1)
		}

		tmpl, err := template.New("example").Parse(string(content))
		if err != nil {
			fmt.Println("Error parsing template:", err)
			os.Exit(1)
		}

		dataDev := TemplateData{Mode: mode}
		var resultDev bytes.Buffer
		err = tmpl.Execute(&resultDev, dataDev)
		if err != nil {
			fmt.Println("Error executing template:", err)
			os.Exit(1)
		}

		err = ioutil.WriteFile(output, resultDev.Bytes(), 0644)
		if err != nil {
			fmt.Println("Error writing to file:", err)
			os.Exit(1)
		}

		fmt.Printf("%s has been created successfully. \n", output)
	}

}
