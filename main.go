package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

func cwdOrExit() string {
	cwd, err := os.Getwd()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	return cwd
}

func commandPath(locations *[]Location) {
	cwd := cwdOrExit()

	home := os.Getenv("HOME")
	var path = cwd
	if strings.HasPrefix(path, home) {
		path = strings.Replace(path, home, "~", 1)
	}

	match := MatchCurrentPath(cwd, locations)
	if match != nil && match.CurrentPathCommand != "" {
		out, err := ExecuteCommandQuietlyCaptureOutput(match.CurrentPathCommand, []string{}, cwd)
		ExitIfNonZero(err)
		path = strings.Trim(out, "\n")
	}

	fmt.Println(path)
}

func commandChange(name string, create bool, layout bool, locations *[]Location, creators *[]Creator, layouts *[]Layout) string {
	cwd := cwdOrExit()
	match := MatchChangePath(name, cwd, locations)
	if match == nil {
		if create {
			if err := os.Chdir(commandCreate(name, locations, creators)); err != nil {
				fmt.Fprintln(os.Stderr, "unable to change to created directory:", err)
				os.Exit(1)
			}
			return commandChange(name, false, layout, locations, creators, layouts)
		} else {
			fmt.Fprintln(os.Stderr, "location named:", name, "does not exist or does not match the current directory")
			os.Exit(1)
		}
	}

	if match.ChangePathCommand == "" {
		os.Exit(1)
	}

	raw, err := ExecuteCommandQuietlyCaptureOutput(match.ChangePathCommand, []string{}, cwd)
	ExitIfNonZero(err)
	d := VerifyDirectory(raw)
	if !layout || match.LayoutName == "" {
		return d
	}
	if err := os.Chdir(d); err != nil {
		fmt.Fprintln(os.Stderr, "unable to change to directory:", d, " ", err)
		os.Exit(1)
	}
	return doLayout(match.LayoutName, []int{}, locations, layouts)
}

func commandCreate(name string, locations *[]Location, creators *[]Creator) string {
	match := MatchName(name, locations)
	if match == nil {
		fmt.Fprintln(os.Stderr, "no creatable location named:", name)
		os.Exit(1)
	}

	if match.CreatorName == "" {
		os.Exit(1)
	}

	creator := MatchCreatorName(match.CreatorName, creators)
	if creator == nil {
		fmt.Fprintln(os.Stderr, "no creator named:", name)
		os.Exit(1)
	}
	if creator.Command == "" {
		fmt.Fprintln(os.Stderr, "no command for creator:", creator.Name)
		os.Exit(1)
	}

	raw, err := ExecuteCommandInteractiveCaptureCwd(creator.Command, []string{})
	ExitIfNonZero(err)
	return VerifyDirectory(raw)
}

func commandLayout(name string, position []int, locations *[]Location, layouts *[]Layout) string {
	match := MatchName(name, locations)
	if match == nil {
		fmt.Fprintln(os.Stderr, "no location named:", name)
		os.Exit(1)
	}

	if match.LayoutName == "" {
		os.Exit(1)
	}
	return doLayout(match.LayoutName, position, locations, layouts)
}

func doLayout(name string, position []int, locations *[]Location, layouts *[]Layout) string {

	layout := MatchLayoutName(name, layouts)
	if layout == nil {
		fmt.Fprintln(os.Stderr, "no layout named:", name)
		os.Exit(1)
	}
	var children = layout.Children
	var command = layout.Command
	var location = layout.LocationName
	for _, p := range position {
		if len(children) <= p {
			fmt.Fprintln(os.Stderr, "no children at position:", PositionString(position))
			os.Exit(1)
		}
		child := children[p]
		command = child.Command
		children = child.Children
		location = child.LocationName
	}

	var cwd = cwdOrExit()

	if location != "" {
		cwd = commandChange(location, false, false, locations, &[]Creator{}, layouts)
		if err := os.Chdir(cwd); err != nil {
			fmt.Fprintln(os.Stderr, "unable to change to directory:", cwd, " ", err)
			os.Exit(1)
		}
	}

	for i, child := range children {
		if os.Getenv("SPACER_TMUX_DISABLED") == "" {
			// fmt.Println("child: ", child.Direction, PositionString(append(position, i)))
			SplitWindow(child.Direction, child.Percent, append(position, i), name)
		}
	}

	if command != "" {
		raw, err := ExecuteCommandInteractiveCaptureCwd(command, []string{})
		ExitIfNonZero(err)
		cwd = VerifyDirectory(raw)
	}
	return cwd
}

func printNames(locations *[]Location) {
	for _, l := range *locations {
		fmt.Println(l.Name)
	}
}

func WriteCwd(f string, cwd string) {
	if f == "" {
		fmt.Println("cd", cwd)
		return
	}
	if err := os.WriteFile(f, []byte(cwd+"\n"), 0666); err != nil {
		fmt.Fprintln(os.Stderr, "error writing to file:", f, "cwd:", cwd, "err:", err)
		os.Exit(1)
	}
}

func main() {
	configFilename := flag.String("config", os.Getenv("SPACER_CONFIG"), "config file (yaml), or set SPACER_CONFIG")
	cwdfile := flag.String("cwd_file", "", "a file to write the new cwd to")
	flag.Parse()

	config := ReadConfig(*configFilename)

	switch flag.Arg(0) {
	case "path":
		fs := flag.NewFlagSet("path", flag.ExitOnError)
		fs.Parse(flag.Args()[1:])
		commandPath(&config.Spacer.Locations)
	case "change":
		fs := flag.NewFlagSet("change", flag.ExitOnError)
		create := fs.Bool("create", false, "create the location if the path does not match")
		layout := fs.Bool("layout", false, "layout the location")
		fs.Parse(flag.Args()[1:])
		if fs.NArg() != 1 {
			printNames(&config.Spacer.Locations)
			os.Exit(1)
		}
		cwd := commandChange(fs.Arg(0), *create, *layout, &config.Spacer.Locations, &config.Spacer.Creators, &config.Spacer.Layouts)
		WriteCwd(*cwdfile, cwd)
	case "create":
		fs := flag.NewFlagSet("create", flag.ExitOnError)
		fs.Parse(flag.Args()[1:])
		if fs.NArg() != 1 {
			printNames(&config.Spacer.Locations)
			os.Exit(1)
		}
		cwd := commandCreate(fs.Arg(0), &config.Spacer.Locations, &config.Spacer.Creators)
		WriteCwd(*cwdfile, cwd)
	case "layout":
		fs := flag.NewFlagSet("layout", flag.ExitOnError)
		position := fs.String("position", "", "internally used for multi-pane layout")
		fs.Parse(flag.Args()[1:])
		if fs.NArg() != 1 {
			printNames(&config.Spacer.Locations)
			os.Exit(1)
		}
		p, perr := ParsePositions(*position)
		if perr != nil {
			fmt.Fprintln(os.Stderr, perr)
			os.Exit(1)
		}
		if len(p) == 0 {
			cwd := commandLayout(fs.Arg(0), p, &config.Spacer.Locations, &config.Spacer.Layouts)
			WriteCwd(*cwdfile, cwd)
		} else {
			cwd := doLayout(fs.Arg(0), p, &config.Spacer.Locations, &config.Spacer.Layouts)
			WriteCwd(*cwdfile, cwd)
		}
	default:
		if flag.NArg() > 0 {
			fmt.Fprintln(os.Stderr, "unknown command:", flag.Arg(0))
		}
		fmt.Fprintln(os.Stderr, "supported commands: path, change, create, layout")
		os.Exit(1)

	}
}
