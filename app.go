package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/mb-14/gomarkov"
	"log"
	"net/http"
	"strings"
)

func crawl_lyric(url string)  string  {
	var lyric string

	res, _ := http.Get(url)

	doc, err := goquery.NewDocumentFromReader(res.Body)

	if err != nil {
		log.Fatal(err)
	}
	doc.Find(".lyrics").Each(func(i int, s *goquery.Selection) {
		lyric = s.Find("p").Text()

	})

	return lyric

}

func main() {
	sources := []string{
		"https://genius.com/Neset-ertas-evvelim-sen-oldun-lyrics",
		"https://genius.com/Neset-ertas-gonul-dag-lyrics",
		"https://genius.com/Neset-ertas-neredesin-sen-lyrics",
		"https://genius.com/Neset-ertas-zuluf-dokulmus-yuze-lyrics",
		"https://genius.com/Selda-bagcan-niye-cattn-kaslarn-lyrics",
	}

	replace_list:= []string{
		"[Verse 1]",
		"[Verse 2]",
		"[Verse 3]",
		"[Chorus]",
		"[Pre-Chorus]",
		"[?]",
	}

	chain := gomarkov.NewChain(2)
	for _,item := range sources{
		lyric := crawl_lyric(item)

		for _,rm_key :=range replace_list{
			lyric  = strings.Replace(lyric,rm_key,"",-1)
		}

		for _, line := range strings.Split(lyric, "\n") {
			if len(line) != 0 {
				chain.Add(strings.Split(line," "))
			}
		}
	}


	order := chain.Order
	tokens := make([]string, 0)
	for i := 0; i < order; i++ {
		tokens = append(tokens, gomarkov.StartToken)
	}
	for tokens[len(tokens)-1] != gomarkov.EndToken {
		next, _ := chain.Generate(tokens[(len(tokens) - order):])
		tokens = append(tokens, next)
	}
	fmt.Println(strings.Join(tokens[order:len(tokens)-2], " "))


	sum := 0
	for i := 1; i < 20; i++ {
		sum += i
		order := chain.Order
		tokens := make([]string, 0)
		for i := 0; i < order; i++ {
			tokens = append(tokens, gomarkov.StartToken)
		}
		for tokens[len(tokens)-1] != gomarkov.EndToken {
			next, _ := chain.Generate(tokens[(len(tokens) - order):])
			tokens = append(tokens, next)
		}
		fmt.Println(strings.Join(tokens[order:len(tokens)-2], " "))

	}
}