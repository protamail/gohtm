package main

import (
    "fmt"
    "main/htm"
    "strconv"
    _ "strings"
)

type Book struct {
    Title        string
    Author       string
    CollectionID int
}

var collections map[string][]Book
var arr []string

type Safe = htm.Safe
var E, V, C, H, U, A = htm.Element, htm.VoidElement, htm.Concat, htm.HTMLEncode, htm.URIComponentEncode, htm.AsIs
var empty = Safe{}

func main() {
    type sel map[bool]string;
    fmt.Println("ee="+fmt.Sprintf("%#v", (sel{true: "1"}[false] == "")))
//    var r Safe
    var e Safe
    for i := 0; i < 1000; i++ {
        for j := 0; j < 100; j++ {
            e = C(e, E(
                "li", C(A(`data-href="`), U(`hj&"'>gjh`), func() Safe {
                    if true {
                        return A(" eee")
                    }
                    return empty
                }(), A(`"`)), empty), C(
                    V(
                        "img", C(A(`src="img`), A(strconv.Itoa(j)), A(`"`))),
                    V("br", empty),
                    E("span", A("data-href='ddd'"), H("dsdsdsd")),
                    V("br", empty),
                ),
            )
        }
        _ = E(
            "html", A(`class="heh" data-href="sdsd?sds=1"`), E(
                "body", empty, E(
                    "nav", A(`class="heh" data-href="sdsd?sds=1"`), E(
                        "div", empty, E(
                            "ul", empty, e,
                        ),
                    ),
                ),
            ),
        )
    }
//    fmt.Println(r.String())
}
