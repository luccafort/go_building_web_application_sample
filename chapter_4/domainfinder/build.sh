#!/bin/bash

# go installで各バイナリをインストールしていればこの実行は必要ない
# TODO:各ツール毎に存在チェックしたほうがいいかもしれない...
# domainfinderのビルド
echo domainfinderをビルドします...
go build -o domainfinder

# synonymsのビルド
echo synonymsをビルドします...
cd ../synonyms
go build -o ../domainfinder/lib/synonyms

# availableのビルド
echo availableをビルドします...
cd ../available
go build -o ../domainfinder/lib/available

# sprinkleのビルド
echo sprinkleをビルドします...
cd ../sprinkle
go build -o ../domainfinder/lib/sprinkle

# coolifyのビルド 
echo coolifyをビルドします...
cd ../coolify
go build -o ../domainfinder/lib/coolify

# domainfyのビルド
echo domainfyをビルドします...
cd ../domainify
go build -o ../domainfinder/lib/domainify

# 完了
echo 完了