package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/fatih/color"
	"log"
	"net/http"
	"os"
	"strings"
)

const (
	ApiUrl = "http://dict.youdao.com/search?q="
)

type Word struct {
	// 关键词
	Keyword string
	// 发音
	Pronounces []string
	// 翻译
	Translations []string
}

func removeEmpty(s string) string {
	s = strings.ReplaceAll(s, "\n", "")
	s = strings.ReplaceAll(s, " ", "")
	s = strings.ReplaceAll(s, "\t", "")
	return s
}

func (w *Word) Parse(keyword string) {
	resp, err := http.Get(ApiUrl + keyword)
	handleErr(err, "response error")
	//defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	handleErr(err, "goquery error")

	// 解析关键词
	w.Keyword = doc.Find("#phrsListTab > h2 > span").Text()

	// 解析音标
	doc.Find("#phrsListTab > h2 > div > span").Each(func(i int, s *goquery.Selection) {
		w.Pronounces = append(w.Pronounces, removeEmpty(s.Text()))
	})

	doc.Find("#phrsListTab > div > ul > li").Each(func(i int, s *goquery.Selection) {
		trans := s.Text()
		w.Translations = append(w.Translations, trans)
	})

}

func Display(word Word) {

	hiGreen := color.New(color.FgHiGreen)
	hiCyan := color.New(color.FgCyan)
	//white:=color.New(color.FgWhite)

	color.HiMagenta("%s\n", word.Keyword)

	for _, v := range word.Pronounces {
		_, err := hiCyan.Printf("  %s  ", v)
		handleErr(err, "print error")
	}
	fmt.Println()

	for _, v := range word.Translations {
		vs := strings.SplitN(v, ".", 2)
		if len(vs) > 1 {
			_, err := hiGreen.Printf("  %s.", vs[0])
			handleErr(err, "print error")
			color.HiBlue("%s\n", vs[1])
		}else {
			color.HiBlue("  %s\n", v)
		}

	}

}

func handleErr(err error, msg string) {
	if err != nil {
		log.Fatalf("%s > %s\n", msg, err)
	}
}

func main() {
	var args = os.Args
	if len(args) < 2 {
		fmt.Println("请输入要查询的单词.")
		os.Exit(0)
	}
	w := &Word{}
	w.Parse(args[1])
	Display(*w)
}
