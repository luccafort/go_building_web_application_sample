package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

const (
	duplicateVowel bool = false // 母音を重ねる
	removeVowel    bool = true  // 母音を削除する
)

func randBool() bool {
	return rand.Intn(2) == 0
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		word := []byte(s.Text())
		// 変換を行うかどうかは1/2の確率(1/2)
		if randBool() {
			var vI int = -1
			for i, char := range word {
				switch char {
				case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
					// さらに1/2の確率で母音後がある場合のみ処理を実行(1/4)
					if randBool() {
						vI = i
					}
				}
			}
			if vI >= 0 {
				// さらにさらに1/2の確率で文字列操作を実行する(1/8)
				switch randBool() {
				case duplicateVowel:
					// 1文字追加
					word = append(word[:vI+1], word[vI:]...)
				case removeVowel:
					// 1文字除去
					word = append(word[:vI], word[vI+1:]...)
				}
			}
		}
		fmt.Println(string(word))
	}
}
