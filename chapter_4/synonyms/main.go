package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/luccafort/building_web_application/chapter_4/thesaurus"
)

func main() {
	apiKey := os.Getenv("BHT_APIKEY")

	if len(apiKey) == 0 {
		log.Fatal("APIキーの取得に失敗しました")
	}

	thesaurus := &thesaurus.BigHuge{APIKey: apiKey}
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		word := s.Text()
		syns, err := thesaurus.Synonyms(word)
		if err != nil {
			log.Fatalf("'%q'の類語検索に失敗しました: %v\n", word, err)
		}
		if len(syns) == 0 {
			log.Fatalf("'%q'の類語はありませんでした\n")
		}
		//fmt.Printf("'%v'の類語を出力します\n\n", word) // コマンドラインで実行するとノイズになる
		for _, syn := range syns {
			fmt.Println(syn)
		}
	}
}
