package crawler

import (
	"fmt"
	"os"
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

func TestGetLinks(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		in   string
		want []Link
	}{
		{
			name: "case 1: Proper",
			in:   "http://go-colly.org",
			want: []Link{
				{URL: "http://go-colly.org/"},
				{URL: "/docs/"},
				{URL: "/articles/"},
				{URL: "/services/"},
				{URL: "/datasets/"},
				{URL: "https://godoc.org/github.com/gocolly/colly"},
				{URL: "https://github.com/gocolly/colly"},
				{URL: "http://go-colly.org/"},
				{URL: ""},
				{URL: "/docs/"},
				{URL: "/articles/"},
				{URL: "/services/"},
				{URL: "/datasets/"},
				{URL: "https://godoc.org/github.com/gocolly/colly"},
				{URL: "https://github.com/gocolly/colly"},
				{URL: "https://github.com/gocolly/colly"},
				{URL: "http://go-colly.org/docs/"},
				{URL: "https://github.com/gocolly/colly/blob/master/LICENSE.txt"},
				{URL: "https://github.com/gocolly/colly"},
				{URL: "#"},
				{URL: "http://go-colly.org/contact/"},
				{URL: "http://go-colly.org/docs/"},
				{URL: "http://go-colly.org/services/"},
				{URL: "https://github.com/gocolly/colly"},
				{URL: "https://github.com/gocolly/site/"},
				{URL: "http://go-colly.org/sitemap.xml"},
			},
		},
	}

	for _, v := range tests {
		test := v
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			got, err := GetLinks(test.in, "", os.Stdout)
			if err != nil {
				t.Errorf("error: %v", err)
			}

			if !equal(got, test.want) {
				t.Errorf("got\twant\n%v", toStringTwoLinkSlices(got, test.want))
			}
		})
	}
}

func TestFilterByDomain(t *testing.T) {
	t.Parallel()

	type input struct {
		links  []Link
		domain string
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

			got := filterByDomain(test.in.links, test.in.domain)

			if !equal(got, test.want) {
				t.Errorf("got want\n%v", toStringTwoLinkSlices(got, test.want))
			}
		})
	}
}
