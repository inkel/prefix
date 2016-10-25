package main

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"

	"bufio"
	"github.com/chzyer/readline"
	"unicode"
	"unicode/utf8"
)

func perror(err error) {
	fmt.Fprintf(os.Stderr, "\x1b[31m%v\x1b[0m\n", err)
}

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "usage: %v command [arguments]\n", os.Args[0])
		os.Exit(1)
	}

	exe := os.Args[1]
	args := os.Args[2:]
	ps := strings.Join(os.Args[1:], " ") + " "

	rl, err := readline.New(ps)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(2)
	}

	for {
		ln, err := rl.Readline()

		if err == io.EOF {
			break
		} else if err != nil {
			perror(err)
			continue
		}

		parts, _ := Parse(ln)

		cmd := exec.Command(exe, append(args, parts...)...)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		if err = cmd.Run(); err != nil {
			perror(err)
		}
	}
}

func scanner(in string) *bufio.Scanner {
	var quoted = false

	s := bufio.NewScanner(strings.NewReader(in))

	s.Split(func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		start := 0

		// Skip leading spaces.
		for width := 0; start < len(data); start += width {
			var r rune
			r, width = utf8.DecodeRune(data[start:])
			if !unicode.IsSpace(r) {
				break
			}
		}

		for width, i := 0, start; i < len(data); i += width {
			var r rune
			r, width = utf8.DecodeRune(data[i:])
			if unicode.IsSpace(r) && !quoted {
				return i + width, data[start:i], nil
			} else if r == '"' {
				quoted = !quoted
				if quoted {
					return i + width, nil, nil
				}
				return i + width, data[start:i], nil
			}
		}

		if atEOF && len(data) > start {
			return len(data), data[start:], nil
		}

		return start, nil, nil
	})

	return s
}

func Parse(in string) ([]string, error) {
	var parts []string

	s := scanner(in)

	for s.Scan() {
		parts = append(parts, s.Text())
	}

	return parts, nil
}
