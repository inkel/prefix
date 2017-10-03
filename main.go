package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"

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

	if _, err := exec.LookPath(exe); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(3)
	}

	rl, err := readline.New(pretty(exe, args) + " ")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(2)
	}

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

func pretty(cmd string, args []string) string {
	buf := bytes.NewBufferString(cmd)

	for _, arg := range args {
		fmts := " %s"
		for _, b := range arg {
			if isSpace(b) {
				fmts = " %q"
				break
			}
		}

		fmt.Fprintf(buf, fmts, arg)
	}
	return buf.String()
}

// Borrowed from unicode.IsSpace
func isSpace(r rune) bool {
	switch r {
	case '\t', '\n', '\v', '\f', '\r', ' ', 0x85, 0xA0:
		return true
	}
	return false
}
