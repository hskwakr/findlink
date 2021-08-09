package cli

import (
	"bytes"
	"strings"
	"testing"
)

func TestParse(t *testing.T) {
	tests := []struct {
		name string
		in   string
		want int
	}{
		{
			name: "case 1: Proper",
			in:   AppName + " http://www.foo.bar/index.html",
			want: 0,
		},
		{
			name: "case 2: Arguments error with no argument",
			in:   AppName,
			want: 1,
		},
	}

	outStream, errStream := new(bytes.Buffer), new(bytes.Buffer)
	app := &CLI{outStream, errStream}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			args := strings.Split(test.in, " ")

			status := app.parse(args)
			if status != test.want {
				t.Errorf("ExitStatus: %d, want: %d", status, test.want)
			}
		})
	}
}

func TestUrlValidation(t *testing.T) {
	tests := []struct {
		name string
		in   string
		want bool
	}{
		{
			name: "case 1: Proper",
			in:   "http://example.com",
			want: true,
		},
		{
			name: "case 1: Proper",
			in:   "https://example.com",
			want: true,
		},
		{
			name: "case 2: Empty string",
			in:   "",
			want: false,
		},
		{
			name: "case 3: Input is NOT url (not scheme)",
			in:   "example.com",
			want: false,
		},
		{
			name: "case 3: Input is NOT url (wrong scheme)",
			in:   "ftp://example.com",
			want: false,
		},
		{
			name: "case 3: Input is NOT url (only foo)",
			in:   "foo",
			want: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := urlValidation(test.in)
			if got != test.want {
				t.Errorf("got: %v, want: %v", got, test.want)
			}
		})
	}
}
