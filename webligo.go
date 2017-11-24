package main

import (
	"fmt"
	"log"
	"unicode/utf8"

	"github.com/PuerkitoBio/goquery"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	searchWord = kingpin.Arg("searchWord", "searchWord string").Required().String()
)

func main() {
	kingpin.Parse()

	doc, err := goquery.NewDocument("http://ejje.weblio.jp/content/" + *searchWord)
	if err != nil {
		log.Fatal(err)
	}
	if isIncludeMultibyte(*searchWord) {
		fmt.Printf("単語             : %s\n", *searchWord)
		fmt.Printf("主な意味         : %s\n", doc.Find(".content-explanation").Text())
		fmt.Printf("音節             : %s\n", doc.Find(".syllableEjje").Text())
		fmt.Printf("発音記号・読み方 : %s\n", doc.Find(".phoneticEjjeDesc").Text())
	} else {
		fmt.Printf("日単語           : %s\n", *searchWord)
		fmt.Printf("英単語           : %s\n", doc.Find(".content-explanation").Text())
	}
}

func isIncludeMultibyte(str string) bool {
	if utf8.RuneCountInString(str) != len(str) {
		return true
	}
	return false
}
