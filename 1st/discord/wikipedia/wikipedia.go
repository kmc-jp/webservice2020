package wikipedia

import (
	"bytes"
	"io"
	"net/http"
	"regexp"
)

//Wikipedia put wikipedia page data
type Wikipedia struct {
	Text  string
	Title string
}

//Head put wikipedia head data
var Head *regexp.Regexp

func init() {
	Head = regexp.MustCompile("<h1 id=\"firstHeading\".*</h1>")
}

//Get gets wikipedia page text and title at random
func (w *Wikipedia) Get() {
	resp, err := http.Get("http://ja.wikipedia.org/wiki/Special:Randompage")
	if err != nil {
		return
	}
	defer resp.Body.Close()

	var buf *bytes.Buffer = new(bytes.Buffer)
	io.Copy(buf, resp.Body)
	var ret []byte = buf.Bytes()

	var top []byte = bytes.Split(ret, []byte("<div class=\"mw-parser-output\">"))[1]
	var content []byte = bytes.Split(bytes.Split(top, []byte("<p>"))[1], []byte("</p>"))[0]

	var contentText []byte

	for i, c := range bytes.Split(content, []byte("<")) {
		if i == 0 {
			contentText = append(contentText, c...)
			continue
		}
		contentText = append(contentText, bytes.Split(c, []byte(">"))[1]...)
	}
	w.Text = string(contentText)
	w.Title = string(bytes.Split(bytes.Split(ret, []byte("class=\"firstHeading\" lang=\"ja\">"))[1], []byte("</h1>"))[0])
}
