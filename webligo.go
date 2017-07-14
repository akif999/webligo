package main

import (
	"fmt"
	"log"
	"unicode/utf8"

	"github.com/PuerkitoBio/goquery"
	"gopkg.in/alecthomas/kingpin.v2"
)

const (
	EnToJp = iota
	JpToEn
)

var (
	searchWord = kingpin.Arg("searchWord", "searchWord string").Required().String()
)

func main() {
	kingpin.Parse()

	mode := EnToJp
	// if searchWord includes multibyte charactor, set mode to JpToEn
	if utf8.RuneCountInString(*searchWord) != len(*searchWord) {
		mode = JpToEn
	}
	doc, err := goquery.NewDocument("http://ejje.weblio.jp/content/" + *searchWord)
	if err != nil {
		log.Fatal(err)
	}
	if mode == EnToJp {
		fmt.Printf("語彙             : %s\n", *searchWord)
		fmt.Printf("主な意味         : %s\n", doc.Find(".content-explanation").Text())
		fmt.Printf("音節             : %s\n", doc.Find(".syllableEjje").Text())
		fmt.Printf("発音記号・読み方 : %s\n", doc.Find(".phoneticEjjeDesc").Text())
	} else {
		fmt.Printf("日語彙           : %s\n", *searchWord)
		fmt.Printf("英語彙           : %s\n", doc.Find(".content-explanation").Text())
	}
}
