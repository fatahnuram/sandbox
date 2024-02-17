package http

import "testing"

func TestParsePathParameter(t *testing.T) {
	suites := []struct {
		name string
		path string
		want []string
	}{
		{name: "root, trailing", path: "/", want: []string{}},
		{name: "root, no trailing", path: "", want: []string{}},
		{name: "department, trailing", path: "/departments/", want: []string{"departments"}},
		{name: "department, no trailing", path: "/departments", want: []string{"departments"}},
		{name: "department id, trailing", path: "/departments/123/", want: []string{"departments", "123"}},
		{name: "department id, no trailing", path: "/departments/123", want: []string{"departments", "123"}},
		{name: "random, trailing", path: "/abce/def/gh/jklm/", want: []string{"abce", "def", "gh", "jklm"}},
		{name: "random, no trailing", path: "/abce/def/gh/jklm", want: []string{"abce", "def", "gh", "jklm"}},
	}

	for _, suite := range suites {
		t.Run(suite.name, func(t *testing.T) {
			got := parsePathParameter(suite.path)
			if len(got) != len(suite.want) {
				t.Errorf("result slice length different, want: %v, got: %v", len(suite.want), len(got))
			}
			if len(got) > 0 {
				for i := range got {
					if got[i] != suite.want[i] {
						t.Errorf("parsed path malformed, want: %v, got: %v", suite.want, got)
					}
				}
			}
		})
	}
}
