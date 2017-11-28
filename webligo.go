package main

import (
	"fmt"
	"log"
	"unicode/utf8"

	"github.com/PuerkitoBio/goquery"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	word = kingpin.Arg("word", "target word").Required().String()
)

func main() {
	kingpin.Parse()

	doc, err := goquery.NewDocument("http://ejje.weblio.jp/content/" + *word)
	if err != nil {
		log.Fatal(err)
	}
	if isIncludeMultibyte(*word) {
		fmt.Printf("日単語           : %s\n", *word)
		fmt.Printf("英単語           : %s\n", doc.Find(".content-explanation").Text())
	} else {
		fmt.Printf("単語             : %s\n", *word)
		fmt.Printf("主な意味         : %s\n", doc.Find(".content-explanation").Text())
		fmt.Printf("音節             : %s\n", doc.Find(".syllableEjje").Text())
		fmt.Printf("発音記号・読み方 : %s\n", doc.Find(".phoneticEjjeDesc").Text())
	}
}

func isIncludeMultibyte(str string) bool {
	return utf8.RuneCountInString(str) != len(str)
}
