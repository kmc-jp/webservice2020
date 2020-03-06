package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

//credit Put credential data
var credit credential

type credential struct {
	clientID string `json:"clientID"`
	secret   string `json:"secret"`
}

func getKeys() {
	bytes, err := ReadFile("credential.json")
	if err != nil {
		panic("panic!:Cannot read credentials!\n")
	}
	err := json.Unmarshal(bytes, &credit)
	if err != nil {
		log.Printf("json error:\n%v", err)
		panic("panic!:Cannot marshal credential json!\n")
	}
}

//ExportFile 指定したファイル名でデータを保存
func ExportFile(filename string, bytes []byte) {
	ioutil.WriteFile(filename, bytes, os.ModePerm)
}

//ReadFile 指定したファイルをbyte形式で読み込む
func ReadFile(filename string) ([]byte, error) {
	bytes, err := ioutil.ReadFile(filename)
	return bytes, err
}
