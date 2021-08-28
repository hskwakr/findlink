// +build integration

package main_test

import (
	"bytes"
	"strings"
	"testing"

	"github.com/hskwakr/findlink/cli"
)

func TestCrawl(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		in   string
		want int
	}{
		{
			name: "case 1: Proper",
			in:   cli.AppName + " http://localhost/example.html",
			want: 0,
		},
		{
			name: "case 2: Arguments error with no argument",
			in:   cli.AppName,
			want: 1,
		},
		{
			name: "case 3: Option -d",
			in:   cli.AppName + " -d=example.com http://localhost/example.html",
			want: 0,
		},
	}

	outStream, errStream := new(bytes.Buffer), new(bytes.Buffer)
	app := &cli.CLI{
		OutStream: outStream,
		ErrStream: errStream,
	}

	for _, v := range tests {
		test := v

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			args := strings.Split(test.in, " ")
			status := app.Run(args)

			if status != test.want {
				t.Errorf("ExitStatus: %d, want: %d", status, test.want)
			}
		})
	}
}
