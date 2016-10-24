package main

import (
	"reflect"
	"testing"
)

func TestParse(t *testing.T) {
	tests := []struct {
		line   string
		parsed []string
	}{
		{"add main.go", []string{"add", "main.go"}},
		{"     add main.go", []string{"add", "main.go"}},
		{"add main.go     ", []string{"add", "main.go"}},
		{"add                     main.go", []string{"add", "main.go"}},
		{"    add     main.go   ", []string{"add", "main.go"}},
		{"ci -m \"initial commit\"", []string{"ci", "-m", "initial commit"}},
		{"commit --ammend --reuse-message=HEAD", []string{"commit", "--ammend", "--reuse-message=HEAD"}},
	}

	for _, ts := range tests {
		ts := ts
		t.Run(ts.line, func(t *testing.T) {
			args, err := Parse(ts.line)

			if err != nil {
				// TODO
			} else if !reflect.DeepEqual(args, ts.parsed) {
				t.Errorf("expected %#v, got %#v\n", ts.parsed, args)
			}
		})
	}
}
