package mui

import (
	"testing"
)

type combo struct {
	Tag    Tag
	Output string
}

var tests = []combo{
	combo{
		Tag: Tag{
			Name:     "p",
			Contents: "text123",
		},
		Output: "<p>text123</p>",
	},
	combo{
		Tag: Tag{
			Name: "one",
			Contents: Tag{
				Name:     "two",
				Contents: "nada",
			},
		},
		Output: "<one><two>nada</two></one>",
	},
	combo{
		Tag: Tag{
			Name: "one",
			Contents: []Tag{
				Tag{
					Name:     "two",
					Contents: "foo",
				},
				Tag{
					Name:     "three",
					Contents: "bar",
				},
			},
		},
		Output: "<one>\n  <two>foo</two>\n  <three>bar</three>\n</one>",
	},
}

func TestTag(t *testing.T) {
	for _, check := range tests {
		out := check.Tag.String()
		if out != check.Output {
			t.Errorf("Output %#v does not match expected %#v", out, check.Output)
		}
	}
}
