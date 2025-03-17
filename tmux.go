package main

import (
	"os"
	"strconv"
)

func directionFlags(d PaneDirection) []string {
	switch d {
	case UP:
		return []string{"-v", "-b"}
	case DOWN:
		return []string{"-v"}
	case LEFT:
		return []string{"-h", "-b"}
	default: // RIGHT
		return []string{"-h"}
	}
}

func SplitWindow(direction PaneDirection, percent int, position []int, layout string) {
	var args = []string{"split-window", "-d"}
	if percent > 0 {
		args = append(args, []string{"-p", strconv.Itoa(percent)}...)
	}
	args = append(args, directionFlags(direction)...)
	args = append(args, []string{"-t", os.Getenv("TMUX_PANE")}...)
	// args = append(args, "-P")
	// args = append(args, []string{"-F", "#{pane_id}"}...)
	args = append(args, "bash --rcfile <(echo '. ~/.bashrc; spacer-bash layout-internal --position "+PositionString(position)+" "+layout+"')")
	ExitIfNonZero(ExecuteCommandInteractive(`tmux "$@"`, args))
}
