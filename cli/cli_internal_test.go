package cli

import (
	"bytes"
	"strings"
	"testing"
)

func TestParse(t *testing.T) {
	t.Parallel()

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
		{
			name: "case 3: Option -o",
			in:   AppName + "-o=links.json http://www.foo.bar/index.html",
			want: 0,
		},
	}

	outStream, errStream := new(bytes.Buffer), new(bytes.Buffer)
	app := &CLI{
		OutStream: outStream,
		ErrStream: errStream,
		args: args{
			url: "",
			o:   nil,
			d:   nil,
		},
	}

	for _, v := range tests {
		test := v

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			args := strings.Split(test.in, " ")
			status := app.parse(args)

			if status != test.want {
				t.Errorf("ExitStatus: %d, want: %d", status, test.want)
			}
		})
	}
}

func TestUrlValidation(t *testing.T) {
	t.Parallel()

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

	for _, v := range tests {
		test := v

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			got := urlValidation(test.in)
			if got != test.want {
				t.Errorf("got: %v, want: %v", got, test.want)
			}
		})
	}
}
