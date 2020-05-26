package main

import (
	"io/ioutil"
	"os"
)

//ExportFile 指定したファイル名でデータを保存
func ExportFile(filename string, bytes []byte) {
	ioutil.WriteFile(filename, bytes, os.ModePerm)
}

//ReadFile 指定したファイルをbyte形式で読み込む
func ReadFile(filename string) ([]byte, error) {
	bytes, err := ioutil.ReadFile(filename)
	return bytes, err
}
