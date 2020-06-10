package idol

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

//Idol put Idol info
type Idol struct {
	Type   string `json:"type"`
	Name   string `json:"name"`
	URL    string `json:"url"`
	Age    string `json:"age"`
	Birth  string `json:"birth"`
	Star   string `json:"star"`
	Blood  string `json:"blood"`
	Height string `json:"height"`
	Weight string `json:"weight"`
	B      string `json:"b"`
	W      string `json:"w"`
	H      string `json:"j"`
	Hand   string `json:"hand"`
	Locate string `json:"locate"`
	Hobby  string `json:"hobby"`
	CV     string `json:"cv"`
	Deploy string `json:"deploy"`
}

//MakeDict make dictionary file
func MakeDict(Path string) error {
	var x []Idol = Get()
	b, err := json.MarshalIndent(x, "", "    ")
	if err != nil {
		return fmt.Errorf("Marshal Error")
	}
	err = ioutil.WriteFile(Path, b, os.ModePerm)
	if err != nil {
		return fmt.Errorf("Write Error")
	}
	return nil
}

//Get gets idol info from web
func Get() []Idol {
	resp, err := http.Get("https://imascg-slstage-wiki.gamerch.com/%E3%82%A2%E3%82%A4%E3%83%89%E3%83%AB%E4%B8%80%E8%A6%A7")
	if err != nil {
		return []Idol{}
	}
	defer resp.Body.Close()

	var buf *bytes.Buffer = new(bytes.Buffer)
	io.Copy(buf, resp.Body)
	var ret []byte = buf.Bytes()

	var top []byte = bytes.Split(ret, []byte("id=\"content_1_2\""))[1]
	var content string = string(bytes.Split(bytes.Split(bytes.Split(
		top,
		[]byte("id=\"content_1_2\""))[0],
		[]byte("t-line-img word-keep-all"))[1],
		[]byte("<tbody>"))[1],
	)

	var table []Idol

	for i, x := range strings.Split(content, "<tr>") {
		if i == 0 {
			continue
		}

		var idol Idol

		for _, y := range strings.Split(x, "<td") {
			switch {
			case strings.Contains(y, "data-col=\"1\">"):
				if len(strings.Split(y, ">")) < 2 || len(strings.Split(strings.Split(strings.Split(y, ">")[1], "</td")[0], "a href=\"")) < 2 {
					continue
				}
				idol.URL = strings.Split(strings.Split(strings.Split(strings.Split(y, ">")[1], "</td")[0], "a href=\"")[1], "\"")[0]
				idol.Name = strings.Split(strings.Split(y, "title=\"")[1], "\"")[0]
			case strings.Contains(y, "data-col=\"2\">"):
				idol.Age = strings.Split(strings.Split(y, "data-col=\"2\">")[1], "</")[0]
			case strings.Contains(y, "data-col=\"3\">"):
				idol.Birth = strings.Split(strings.Split(y, "data-col=\"3\">")[1], "</")[0]
			case strings.Contains(y, "data-col=\"4\">"):
				idol.Star = strings.Split(strings.Split(y, "data-col=\"4\">")[1], "</")[0]
			case strings.Contains(y, "data-col=\"5\">"):
				idol.Blood = strings.Split(strings.Split(y, "data-col=\"5\">")[1], "</")[0]
			case strings.Contains(y, "data-col=\"6\">"):
				idol.Height = strings.Split(strings.Split(y, "data-col=\"6\">")[1], "</")[0]
			case strings.Contains(y, "data-col=\"7\">"):
				if strings.Contains(y, "span") {
					idol.Weight = strings.Split(strings.Split(y, "size8\">")[1], "</")[0]
					continue
				}
				idol.Weight = strings.Split(strings.Split(y, "data-col=\"7\">")[1], "</")[0]

			case strings.Contains(y, "data-col=\"8\">"):
				if len(strings.Split(y, "</span>")) < 2 {
					continue
				}
				idol.B = strings.Split(strings.Split(y, "</span>")[1], "</")[0]
			case strings.Contains(y, "data-col=\"9\">"):
				idol.W = strings.Split(strings.Split(y, "data-col=\"9\">")[1], "</")[0]
			case strings.Contains(y, "data-col=\"10\">"):
				idol.H = strings.Split(strings.Split(y, "data-col=\"10\">")[1], "</")[0]
			case strings.Contains(y, "data-col=\"11\">"):
				idol.Hand = strings.Split(strings.Split(y, "data-col=\"11\">")[1], "</")[0]
			case strings.Contains(y, "data-col=\"12\">"):
				idol.Locate = strings.Split(strings.Split(y, "data-col=\"12\">")[1], "</")[0]
			case strings.Contains(y, "data-col=\"13\">"):
				idol.Hobby = strings.Split(strings.Split(y, "data-col=\"13\">")[1], "</")[0]
			case strings.Contains(y, "data-col=\"14\">"):
				idol.CV = strings.Split(strings.Split(y, "data-col=\"14\">")[1], "</")[0]
			case strings.Contains(y, "data-col=\"15\">"):
				idol.Deploy = strings.Split(strings.Split(y, "data-col=\"15\">")[1], "</")[0]
			}

		}
		table = append(table, idol)

	}
	fmt.Printf("%#v\n", table)
	return table
}

//GetImgURL Get image url from URL page
func (idol Idol) GetImgURL() string {
	resp, err := http.Get(idol.URL)
	if err != nil {
		return ""
	}
	defer resp.Body.Close()

	var buf *bytes.Buffer = new(bytes.Buffer)
	io.Copy(buf, resp.Body)
	var ret []byte = buf.Bytes()

	if !strings.Contains(string(ret), "id=\"content_1_1\"") {
		return ""
	}
	var top []byte = bytes.Split(ret, []byte("id=\"content_1_1\""))[1]

	if !strings.Contains(string(top), "a href=\"") {
		return ""
	}
	var url []byte = bytes.Split(bytes.Split(top, []byte("a href=\""))[1], []byte("\""))[0]

	return string(url)
}

//Read read idol data
func Read(Path string) []Idol {
	b, err := ioutil.ReadFile(Path)
	if err != nil {
		fmt.Printf("Error: Cannot Read file\n")
		return []Idol{}
	}

	var info []Idol

	json.Unmarshal(b, &info)
	return info
}
