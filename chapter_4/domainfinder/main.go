package main

import (
	"log"
	"os"
	"os/exec"
)

var cmdChain = []*exec.Cmd{
	exec.Command("lib/synonyms"),
	exec.Command("lib/sprinkle"),
	exec.Command("lib/coolify"),
	exec.Command("lib/domainify"),
	exec.Command("lib/available"),
}

func main() {
	cmdChain[0].Stdin = os.Stdin                 // synonymsに標準入力を読み込ませるようにしている
	cmdChain[len(cmdChain)-1].Stdout = os.Stdout // availableに標準出力を受け取るようにしている

	// 各プログラムの入力元（データの読込）と出力先（処理結果）を設定
	for i := 0; i < len(cmdChain)-1; i++ {
		this := cmdChain[i]
		next := cmdChain[i+1]
		stdout, err := this.StdoutPipe()
		if err != nil {
			log.Panicln(err)
		}
		next.Stdin = stdout
	}

	// 各コマンドをStart()を使ってバックグラウンドで実行
	for _, cmd := range cmdChain {
		// Run()でも実行可能だがコマンド終了まで呼び出し元のコードをブロックしてしまう
		// 今回は全てのコマンドを同時に実行してデータを引き渡す必要があるため不採用
		if err := cmd.Start(); err != nil {
			log.Panicln(err)
		} else {
			// 成功したら最後にコマンド実行のプロセスを殺すことで終了を担保する
			defer cmd.Process.Kill()
		}
	}

	// コマンド実行中に早期にdomainfinderが終了し、実行中のコマンドを強制終了しないようにWaitさせる
	for _, cmd := range cmdChain {
		if err := cmd.Wait(); err != nil {
			log.Panicln(err)
		}
	}
}
