package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

// func commandCommand(name string, args []string, commands *[]Command) error {
// cwd, err := os.Getwd()
// if err != nil {
// 	fmt.Println(err)
// 	os.Exit(1)
// }
//
// match := Match(name, cwd, commands)
// if match == nil {
// 	return nil
// }
//
// 	return ExecuteCommandInteractive(match.Command, args)
// }
//
// func listCommands(commands *[]Command) {
// 	for _, c := range *commands {
// 		fmt.Println(c.Name)
// 	}
// }

func commandPath(locations *[]Location) error {
	cwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	home := os.Getenv("HOME")
	var simplePath = cwd
	if strings.HasPrefix(simplePath, home) {
		simplePath = strings.Replace(simplePath, home, "~", 1)
	}

	match := MatchCurrentPath(cwd, locations)
	if match == nil {
		fmt.Println(simplePath)
		return nil
	}
	if match.CurrentPathCommand != "" {
		out, err := ExecuteCommandQuietlyCaptureOutput(match.CurrentPathCommand, []string{})
		if err != nil {
			return err
		}

		fmt.Print(out)
		return nil
	}

	fmt.Println(simplePath)
	return nil
}

func main() {
	configFilename := flag.String("config", os.Getenv("SPACER_CONFIG"), "config file (yaml), or set SPACER_CONFIG")
	flag.Parse()

	config := ReadConfig(*configFilename)

	switch flag.Arg(0) {
	case "path":
		fs := flag.NewFlagSet("path", flag.ExitOnError)
		fs.Parse(flag.Args())
		ExitIfNonZero(commandPath(&config.Spacer.Locations))
	default:
		if flag.NArg() > 0 {
			fmt.Println("unknown command:", flag.Arg(0))
		}
		fmt.Println("supported commands: path")
		os.Exit(1)

	}

	// commandError := commandCommand(fs.Arg(0), fs.Args()[1:], &config.Commander.Commands)
	// ExitIfNonZero(commandError)
}
