package crawler

import (
	"fmt"
	"reflect"
	"testing"
)

// Compare two variables
func equal(a, b interface{}) bool {
	return reflect.DeepEqual(a, b)
}

func toStringTwoLinkSlices(a, b []Link) string {
	var r string
	str := make([]string, 0)

	if len(a) >= len(b) {
		for i := range a {
			if i < len(b) {
				str = append(str, fmt.Sprintf("%v %v", a[i], b[i]))
			} else {
				str = append(str, fmt.Sprintf("%v {}", a[i]))
			}
		}
	} else {
		for i := range b {
			if i < len(a) {
				str = append(str, fmt.Sprintf("%v %v", a[i], b[i]))
			} else {
				str = append(str, fmt.Sprintf("{} %v", b[i]))
			}
		}
	}

	for _, s := range str {
		r += fmt.Sprintf("%v\n", s)
	}

	return r
}

func TestGetLinks(t *testing.T) {
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

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, err := GetLinks(test.in)
			if err != nil {
				t.Errorf("error: %v", err)
			}

			if !equal(got, test.want) {
				//t.Errorf("got:\n%v\nwant:\n%v\n", got, test.want)
				t.Errorf("got\twant\n%v", toStringTwoLinkSlices(got, test.want))
			}
		})
	}
}
