package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
	"unicode"
)

// トップレベルドメインのリスト
var tlds = []string{"com", "net"}

// 許可済み文字列(英数字+"_-")
const allowedChars = "abcdefghijklmnopqrstuvwxyz0123456789_-"

func main() {
	// TODO:コマンドラインフラグを通じてトップレベルドメインを指定できるよう変更
	rand.Seed(time.Now().UTC().UnixNano())
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		text := strings.ToLower(s.Text())
		var newText []rune
		for _, r := range text {
			// スペースは別の文字へ置換
			if unicode.IsSpace(r) {
				r = '-'
			}
			// 許可済み文字列に含まれていない場合は意図しない文字として扱う
			if !strings.ContainsRune(allowedChars, r) {
				continue
			}
			newText = append(newText, r)
		}
		fmt.Println(string(newText) + "." + tlds[rand.Intn(len(tlds))])
	}
}
