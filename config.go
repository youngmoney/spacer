package main

import (
	"errors"
	"fmt"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
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

type Creator struct {
	Name    string `yaml:"name"`
	Command string `yaml:"command"`
}

type PaneDirection int64

const (
	UP PaneDirection = iota + 1
	DOWN
	LEFT
	RIGHT
)

func (p *PaneDirection) UnmarshalYAML(n *yaml.Node) error {
	v := strings.ToLower(n.Value)
	switch v {
	case "up":
		*p = UP
	case "down":
		*p = DOWN
	case "left":
		*p = LEFT
	case "right":
		*p = RIGHT
	default:
		return errors.New("unknown PaneDirection " + v)
	}
	return nil
}

type Pane struct {
	LocationName string        `yaml:"location_name"`
	Command      string        `yaml:"command"`
	Direction    PaneDirection `yaml:"direction"`
	Percent      int           `yaml:"percent"`
	Children     []Pane        `yaml:"children"`
}

type Layout struct {
	Name         string `yaml:"name"`
	LocationName string `yaml:"location_name"`
	Command      string `yaml:"command"`
	Children     []Pane `yaml:"children"`
}

type Spacer struct {
	Locations []Location `yaml:"locations"`
	Creators  []Creator  `yaml:"creators"`
	Layouts   []Layout   `yaml:"layouts"`
}

type Config struct {
	Spacer Spacer `yaml:"spacer"`
}

func ReadConfig(filename string) Config {
	raw, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Fprintln(os.Stderr, "unable to read config:", filename, err)
		os.Exit(1)
	}

	config := Config{}
	if err := yaml.Unmarshal(raw, &config); err != nil {
		fmt.Fprintln(os.Stderr, "unable to parse config:", filename, err)
		os.Exit(1)
	}

	return config
}
