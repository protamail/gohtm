package htm

import (
	"html"
	"net/url"
	"strings"
)

// contains well-formed HTML fragment
type HTML struct {
	pieces []string
}

type Attr string

func Element(tag string, attr Attr, body HTML) HTML {
	var r HTML
	switch len(body.pieces) {
	case 0:
		r = HTML{make([]string, 1, 1)}
	case 1:
		if len(body.pieces[0]) < 256 {
			return HTML{[]string{"<" + tag + string(attr) + "\n>" + body.pieces[0] + "</" + tag + ">"}}
		}
		r = HTML{[]string{"", body.pieces[0], ""}}
	default:
		r = body
	}
	r.pieces[0] = "<" + tag + string(attr) + "\n>" + r.pieces[0]
	r.pieces[len(r.pieces)-1] += "</" + tag + ">"
	return r
}

var attrEscaper = strings.NewReplacer(`"`, `&quot;`, `<`, `&lt;`)

func Prepend(doctype string, html HTML) HTML {
	if len(html.pieces) > 0 {
		html.pieces[0] = doctype + html.pieces[0]
		return html
	}
	return HTML{[]string{doctype}}
}

func Attributes(nv ...string) Attr {
	sar := make([]string, 0, len(nv)*5/2)
	for i := 1; i < len(nv); i += 2 {
		sar = append(sar, ` `)
		v := nv[i]
		if strings.Index(v, `"`) >= 0 {
			v = attrEscaper.Replace(v)
		}
		sar = append(sar, nv[i-1], `="`, v, `"`)
	}
	return Attr(strings.Join(sar, ""))
}

func JoinAttributes(attrs ...Attr) Attr {
	var n int
	for _, attr := range attrs {
		n += len(attr)
	}

	var b strings.Builder
	b.Grow(n)
	for _, attr := range attrs {
		b.WriteString(string(attr))
	}
	return Attr(b.String())
}

func If[T ~string | HTML](cond bool, result T) T {
	if cond {
		return result
	}
	var r T
	return r
}

func IfCall[T ~string | HTML](cond bool, call func() T) T {
	if cond {
		return call()
	}
	var r T
	return r
}

func IfElse[T ~string | HTML](cond bool, ifR T, elseR T) T {
	if cond {
		return ifR
	}
	return elseR
}

func IfElseCall[T ~string | HTML](cond bool, ifCall func() T, elseCall func() T) T {
	if cond {
		return ifCall()
	}
	return elseCall()
}

// create HTML tag with no closing, e.g. <input type="text">
func VoidElement(tag string, attr Attr) HTML {
	return HTML{[]string{"<" + tag + string(attr) + "\n>"}}
}

func Append(collect HTML, frags ...HTML) HTML {
	var n int
	for _, frag := range frags {
		n += len(frag.pieces)
	}
	if cap(collect.pieces) < len(collect.pieces)+n {
		var newPieces []string
		if len(collect.pieces) > n {
			newPieces = make([]string, 0, len(collect.pieces)*2)
		} else {
			newPieces = make([]string, 0, len(collect.pieces)+n)
		}
		collect.pieces = append(newPieces, collect.pieces...)
	}

	for _, frag := range frags {
		collect.pieces = append(collect.pieces, frag.pieces...)
	}
	return collect
}

func (c HTML) String() string {
	return strings.Join(c.pieces, "")
}

func AsIs(a ...string) HTML {
	return HTML{[]string{strings.Join(a, "")}}
}

// Used to output HTML text, escaping HTML reserved characters <>&"
func HTMLEncode(a string) HTML {
	return HTML{[]string{html.EscapeString(a)}}
}

var URIComponentEncode = url.QueryEscape

var jsStringEscaper = strings.NewReplacer(
	`"`, `\"`,
	`'`, `\'`,
	"`", "\\`",
	`\`, `\\`,
)

func JSStringEscape(a string) HTML {
	return HTML{[]string{jsStringEscaper.Replace(a)}}
}
