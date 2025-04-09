package main

import (
	"encoding/base64"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

// arrayVarFlags store values for -var flag/s
type arrayVarFlags []string

// String is an implementation of the flag.Value interface
func (f *arrayVarFlags) String() string {
	return fmt.Sprintf("%v", *f)
}

// Set is an implementation of the flag.Value interface
func (f *arrayVarFlags) Set(value string) error {
	*f = append(*f, value)
	return nil
}

// parseVarsValues parse -var flags keys and values
func parseVarsValues(list []string) (map[string]string, error) {
	result := make(map[string]string)

	for _, item := range list {

		var cVarName, cVarValue string

		parts := strings.Split(item, "=")
		if len(parts) != 2 {
			return nil, fmt.Errorf("Invalid input value: %v", item)
		}

		cVarName = parts[0]
		cVarValue = parts[1]

		if strings.HasPrefix(cVarValue, "b64:") {
			valueWithoutPrefix := strings.TrimPrefix(cVarValue, "b64:")
			cVarValue = base64.StdEncoding.EncodeToString([]byte(valueWithoutPrefix))
		}

		result[cVarName] = cVarValue
	}

	return result, nil
}

// loadTemplate load go template from file
func loadTemplate(file string) (*template.Template, error) {
	path, err := filepath.Abs(file)
	if err != nil {
		return nil, err
	}

	tmpl, err := template.ParseFiles(path)

	return tmpl, nil
}

// fillTemplate fill template file
func fillTemplate(input string, output string, vars arrayVarFlags) error {
	templateValues, err := parseVarsValues(vars)
	if err != nil {
		return err
	}

	outputFile, _ := filepath.Abs(output)
	fileWriter, err := os.Create(outputFile)
	if err != nil {
		return err
	}

	tmpl, err := loadTemplate(input)
	if err != nil {
		return err
	}

	err = tmpl.Execute(fileWriter, templateValues)
	if err != nil {
		return err
	}

	fileWriter.Close()
	return nil
}

func main() {
	var cliVarFlags arrayVarFlags

	flag.Var(&cliVarFlags, "var", "Thempate vars. ex: -var title='hello' -var body='my friend'")
	input := flag.String("input", "", "Input file")
	output := flag.String("output", "", "Output file")
	v := flag.Bool("version", false, "Show version")
	flag.Parse()

	if *v {
		fmt.Printf("templater version: %v\n", version)
		os.Exit(0)
	}

	if len(*input) == 0 || len(*output) == 0 {
		flag.Usage()
		os.Exit(1)
	}

	err := fillTemplate(*input, *output, cliVarFlags)
	if err != nil {
		fmt.Printf("Error: %s", err.Error())
		os.Exit(1)
	}

	fmt.Println("Templater finished successfully!")
}
