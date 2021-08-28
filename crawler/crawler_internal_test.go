package crawler

import (
	"fmt"
	"reflect"
	"testing"
)

// Compare two variables.
func equal(a, b interface{}) bool {
	return reflect.DeepEqual(a, b)
}

func toStringTwoLinkSlices(a, b []Link) string {
	var r string

	var str []string

	if len(a) >= len(b) {
		str = formatLinks(a, b)
	} else {
		str = formatLinks(b, a)
	}

	for _, s := range str {
		r += fmt.Sprintf("%v\n", s)
	}

	return r
}

func formatLinks(v1, v2 []Link) []string {
	var r []string

	for i := range v1 {
		if i < len(v2) {
			r = append(r, fmt.Sprintf("%v %v", v1[i], v2[i]))
		} else {
			r = append(r, fmt.Sprintf("%v {}", v1[i]))
		}
	}

	return r
}

func TestFilterByString(t *testing.T) {
	t.Parallel()

	type input struct {
		links []Link
		str   string
	}

	tests := []struct {
		name string
		in   input
		want []Link
	}{
		{
			name: "case 1: Proper",
			in: input{
				[]Link{
					{URL: "http://example.com"},
					{URL: "http://example.jp"},
					{URL: "http://foo.com"},
					{URL: "http://www.example.com"},
				},
				"example.com",
			},
			want: []Link{
				{URL: "http://example.com"},
				{URL: "http://www.example.com"},
			},
		},
	}

	for _, v := range tests {
		test := v

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			got := filterByString(test.in.links, test.in.str)

			if !equal(got, test.want) {
				t.Errorf("got want\n%v", toStringTwoLinkSlices(got, test.want))
			}
		})
	}
}
