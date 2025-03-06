package main

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"os"
	"regexp"
)

type Location struct {
	Name               string         `yaml:"name"`
	ChangePathRegex    *regexp.Regexp `yaml:"change_path_regex"`
	ChangePathCommand  string         `yaml:"change_path_command"`
	CurrentPathRegex   *regexp.Regexp `yaml:"current_path_regex"`
	CurrentPathCommand string         `yaml:"current_path_command"`
	CreatorName        string         `yaml:"creator_name"`
	LayoutName         string         `yaml:"layout_name"`
}

type Spacer struct {
	Locations []Location `yaml:"locations"`
}

type Config struct {
	Spacer Spacer `yaml:"spacer"`
}

func ReadConfig(filename string) Config {
	raw, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("unable to read config: ", filename)
		os.Exit(1)
	}

	config := Config{}
	if err := yaml.Unmarshal(raw, &config); err != nil {
		fmt.Println("unable to parse config: ", filename)
		os.Exit(1)
	}

	return config
}
