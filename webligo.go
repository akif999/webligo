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
	refWord = kingpin.Arg("refWord", "refWord string").Required().String()
)

func main() {
	kingpin.Parse()

	// if refWord includes multibyte charactor, set mode to JpToEn
	mode := EnToJp
	if utf8.RuneCountInString(*refWord) != len(*refWord) {
		mode = JpToEn
	}

	doc, err := goquery.NewDocument("http://ejje.weblio.jp/content/" + *refWord)
	if err != nil {
		log.Fatal(err)
	}
	if mode == EnToJp {
		fmt.Printf("単語             : %s\n", *refWord)
		fmt.Printf("主な意味         : %s\n", doc.Find(".content-explanation").Text())
		fmt.Printf("音節             : %s\n", doc.Find(".syllableEjje").Text())
		fmt.Printf("発音記号・読み方 : %s\n", doc.Find(".phoneticEjjeDesc").Text())
	} else {
		fmt.Printf("日単語           : %s\n", *refWord)
		fmt.Printf("英単語           : %s\n", doc.Find(".content-explanation").Text())
	}
}
