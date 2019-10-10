package main

import (
	"fmt"
	"log"
	"unicode/utf8"

	"github.com/PuerkitoBio/goquery"
	"gopkg.in/alecthomas/kingpin.v2"
)

type JpEn struct {
	JpWord string
	EnWord string
}

type EnJp struct {
	Word           string
	Meaning        string
	Syllable       string
	PhoneticSymbol string
}

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
		je := &JpEn{
			JpWord: *word,
			EnWord: doc.Find(".content-explanation").Text(),
		}
		je.String()
	} else {
		ej := &EnJp{
			Word:           *word,
			Meaning:        doc.Find(".content-explanation").Text(),
			Syllable:       doc.Find(".syllableEjje").Text(),
			PhoneticSymbol: doc.Find(".phoneticEjjeDesc").Text(),
		}
		ej.String()
	}
}

func (j *JpEn) String() {
	fmt.Printf("日単語           : %s\n", j.JpWord)
	fmt.Printf("英単語           : %s\n", j.EnWord)
}

func (e *EnJp) String() {
	fmt.Printf("単語             : %s\n", e.Word)
	fmt.Printf("主な意味         : %s\n", e.Meaning)
	fmt.Printf("音節             : %s\n", e.Syllable)
	fmt.Printf("発音記号・読み方 : %s\n", e.PhoneticSymbol)
}

func isIncludeMultibyte(str string) bool {
	return utf8.RuneCountInString(str) != len(str)
}
