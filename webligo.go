package main

import (
	"fmt"
	"log"

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
	fmt.Printf("語彙             : %s\n", *searchWord)
	fmt.Printf("主な意味         : %s\n", doc.Find(".content-explanation").Text())
	fmt.Printf("音節             : %s\n", doc.Find(".syllableEjje").Text())
	fmt.Printf("発音記号・読み方 : %s\n", doc.Find(".phoneticEjjeDesc").Text())
}
