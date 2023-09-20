package htm

import (
	"html"
	"net/url"
	"strings"
)

//contains well-formed HTML fragments
type Safe struct {
	frag []string
}

func Element(tag string, attr Safe, body Safe) Safe {
	//smaller fragments are concatenated as soon as available, larger ones are defered
	ss := make([]string, 0, len(body.frag)+2)
	var result Safe
	if len(attr.frag) > 0 {
		result = AsIs("<"+tag+" "+strings.Join(attr.frag, "")+"\n>")
//		ss = append(ss, "<"+tag+" "+strings.Join(attr.frag, "")+"\n>")
	} else {
		result = AsIs("<"+tag+">")
//		ss = append(ss, "<"+tag+">")
	}
//	ss = append(ss, body.frag...)
//	ss = append(ss, "</"+tag+">")
//	return Safe{ss}
	return Concat(Safe{ss}, result, body, AsIs("</"+tag+">"))
}

//create HTML tag with no closing, e.g. <input type="text">
func VoidElement(tag string, attr Safe) Safe {
	return AsIs("<"+tag+" "+strings.Join(attr.frag, "")+"\n>")
}

func Concat(dst Safe, src ...Safe) Safe {
	for _, s := range src {
		if len(s.frag) > 0 && len(s.frag[0]) > 256 {
			dst.frag = append(dst.frag, s.frag...)
		} else {
			dst.frag = append(dst.frag, strings.Join(s.frag, ""))
		}
	}
	return dst
}

func (c Safe) String() string {
	return strings.Join(c.frag, "")
}

func AsIs(a ...string) Safe {
	return Safe{a}
}

// Used to output HTML text, escaping HTML reserved characters <>&"
func HTMLEncode(a string) Safe {
	return Safe{[]string{html.EscapeString(a)}}
}

func URIComponentEncode(a string) Safe {
	return Safe{[]string{url.QueryEscape(a)}}
}

var JSStringEscaper = strings.NewReplacer(
	`"`, `\"`,
	`'`, `\'`,
	"`", "\\`",
	`\`, `\\`,
)

func JSStringEscape(a string) Safe {
	return Safe{[]string{JSStringEscaper.Replace(a)}}
}
