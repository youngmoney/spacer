package main

import (
	"os"
	"os/exec"
)

func ExitIfNonZero(err interface{}) {
	if err != nil {
		if e, ok := err.(interface{ ExitCode() int }); ok {
			os.Exit(e.ExitCode())
		}
	}
}

func ExecuteCommandQuietly(command string, args []string) error {
	bashArgs := []string{"-c", command, "command"}
	cmd := exec.Command("bash", append(bashArgs, args...)...)

	return cmd.Run()
}

func ExecuteCommandQuietlyCaptureOutput(command string, args []string) (string, error) {
	bashArgs := []string{"-c", command, "command"}
	cmd := exec.Command("bash", append(bashArgs, args...)...)
	out, err := cmd.CombinedOutput()
	return string(out), err
}

func ExecuteCommandInteractive(command string, args []string) error {
	bashArgs := []string{"-c", command, "command"}
	cmd := exec.Command("bash", append(bashArgs, args...)...)

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}
