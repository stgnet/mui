package mui

import (
	"fmt"
	"strings"
)

var indent = "  "

var id_counter = 0

func getId(prefix string) string {
	id_counter += 1
	return fmt.Sprintf("%s%d", prefix, id_counter)
}

var dont_break = []string{
	"li",
	"a",
}

var dont_close = []string{
	"area",
	"base",
	"br",
	"col",
	"command",
	"embed",
	"hr",
	"img",
	"input",
	"keygen",
	"link",
	"meta",
	"param",
	"source",
	"track",
	"wbr",
}

var dont_self_close = []string{
	"script",
	"i",
	"iframe",
	"div",
	"span",
	"title",
}

func isIn(value string, list []string) bool {
	for _, entry := range list {
		if entry == value {
			return true
		}
	}
	return false
}

type Params map[string]interface{}

func NewClass(classes []string) Params {
	return Params{"class": classes}
}

type Tag struct {
	Name     string
	Params   Params
	Contents interface{}
}

// smartly generate neatly formatted nested tags
func (t *Tag) String() string {
	switch content := t.Contents.(type) {
	case []Tag:
		set := []string{}
		for _, element := range content {
			set = append(set, element.String())
		}
		t.Contents = set
		return t.String()
	case Tag:
		t.Contents = content.String()
		return t.String()
	case []string:
		t.Contents = strings.Join(content, "\n")
		return t.String()
	case string:
		html := "<" + t.Name
		tag := strings.Split(t.Name, " ")[0]
		for name, value := range t.Params {
			html += " " + name + "=\"" + fmt.Sprintf("%v", value) + "\""
		}
		if len(content) == 0 && !isIn(tag, dont_self_close) {
			html += "/>"
			return html
		}
		html += ">"
		if len(content) > 0 {
			need_break := strings.Contains(content, "\n") || len(content) > 40
			if need_break && !isIn(tag, dont_break) {
				// note: this wil indent a <pre> in content -- to avoid:
				// process individual lines to add indent with <pre detection
				html += "\n" + indent + strings.Replace(content, "\n", "\n"+indent, -1) + "\n"
			} else {
				html += content
			}
		}
		if !isIn(tag, dont_close) {
			html += "</" + tag + ">"
		}
		return html
	default:
		panic(fmt.Sprintf("Don't know how to tag %#v", t.Contents))
	}
}
