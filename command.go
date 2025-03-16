package main

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
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

func ExecuteCommandQuietlyCaptureOutput(command string, args []string, in string) (string, error) {
	bashArgs := []string{"-c", command, "command"}
	cmd := exec.Command("bash", append(bashArgs, args...)...)

	stdin, err := cmd.StdinPipe()
	if err != nil {
		return "", err
	}

	go func() {
		defer stdin.Close()
		io.WriteString(stdin, in)
	}()

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

func VerifyDirectory(in string) string {
	s := strings.Trim(in, "\n\r")
	if s == "" {
		fmt.Fprintln(os.Stderr, "command has no directory:", s)
		os.Exit(1)
	}
	if !IsDir(s) {
		fmt.Fprintln(os.Stderr, "command ended in non-existant directory:", s)
		os.Exit(1)
	}
	return s
}

func GetAndVerifyDirectory(f *os.File) string {
	data := make([]byte, 1000)
	c, err := f.Read(data)
	if err != nil {
		fmt.Fprintln(os.Stderr, "error reading from cwd file:", err)
		os.Exit(1)
	}
	return VerifyDirectory(string(data[:c]))
}

func ExecuteCommandInteractiveCaptureCwd(command string, args []string) (string, error) {
	f, ferr := os.CreateTemp("", "example")
	if ferr != nil {
		fmt.Fprintln(os.Stderr, "error creating tempfile:", ferr)
		os.Exit(1)
	}
	defer os.Remove(f.Name())

	var c = command
	c = c + "\n"
	c = c + "O=$?;\n"
	c = c + "CWDFILE=\"" + f.Name() + "\";\n"
	c = c + "[ -f \"$CWDFILE\" ] && pwd > \"$CWDFILE\";\n"
	c = c + "exit $O;"

	err := ExecuteCommandInteractive(c, args)
	if err != nil {
		return "", err
	}
	return GetAndVerifyDirectory(f), nil
}
