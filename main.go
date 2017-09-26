package main

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"os/signal"
	"strings"

	"github.com/chzyer/readline"
	"github.com/mattn/go-shellwords"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "usage: %v command [arguments]\n", os.Args[0])
		os.Exit(1)
	}

	exe := os.Args[1]
	args := os.Args[2:]
	ps := strings.Join(os.Args[1:], " ") + " "

	if _, err := exec.LookPath(exe); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(3)
	}

	rl, err := readline.New(ps)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(2)
	}

	signal.Notify(make(chan os.Signal, 1), os.Interrupt)

	sw := shellwords.NewParser()
	sw.ParseBacktick = true
	sw.ParseEnv = true

	for {
		ln, err := rl.Readline()

		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Fprintln(os.Stderr, err)
			continue
		}

		parts, err := sw.Parse(ln)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			continue
		}

		cmd := exec.Command(exe, append(args, parts...)...)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		if err = cmd.Run(); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}
}
