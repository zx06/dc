package main

import (
	"regexp"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	API_URL = "http://dict.youdao.com/search?q="
	
)
var(
	// 音标
	R_PS = regexp.MustCompile(`"phonetic">(.*?)</span>`)
)

func handleErr(err error,msg string)  {
	// TODO
}

func getContent(word string)  {
	resp,err:=http.Get(API_URL+word)
	handleErr(err,"response error")
	defer resp.Body.Close()
	body,_:=ioutil.ReadAll(resp.Body)

	t:=R_PS.FindSubmatch(body)
	fmt.Println(string(t[1]))
}

func main() {
	getContent("love")
}