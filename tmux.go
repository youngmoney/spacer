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
	// -c CWD
	// var command = `split-window -l {percent} -d {directions} -t {pane} -P -F '#{{pane_id}}' '; bash'`
	var args = []string{"split-window", "-d"}
	if percent > 0 {
		args = append(args, []string{"-l", strconv.Itoa(percent)}...)
	}
	args = append(args, directionFlags(direction)...)
	args = append(args, []string{"-t", os.Getenv("TMUX_PANE")}...)
	// args = append(args, "-P")
	// args = append(args, []string{"-F", "#{pane_id}"}...)
	args = append(args, "bash --rcfile <(echo '. ~/.bashrc; spacer::cd layout --position "+PositionString(position)+" "+layout+"')")
	// print(f"\\tmux send-keys -t \"$new\" '{command}' Enter")
	ExitIfNonZero(ExecuteCommandInteractive(`tmux "$@"`, args))
}
