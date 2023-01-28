package src

import (
	"bufio"
	"os"

	"gopkg.in/yaml.v3"
)

// Read a file line by line.
// For reading the domain list
func ReadFileLineByLine(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, nil
}

// Read the example-files YAML file
type File struct {
	Filetype string `yaml:"filetype"`
	Filename string `yaml:"filename"`
}

// Read in the YAML configuration file
func ReadFilesConfigYaml(path string) ([]File, error) {
	var data []File
	config, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	err = yaml.Unmarshal(config, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}
