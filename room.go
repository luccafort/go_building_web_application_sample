package main

type room struct {
	forward chan []byte      // forwardは他のクライアントに転送するためにメッセージを保持するチャネル
	join    chan *client     // join はチャットルームに参加しようとしているクライアントのためのチャネル
	leave   chan *client     // leave はチャットルームから退室しようとしているクライアントのためのチャネル
	clients map[*client]bool // clients には在室している全てのクライアントが保持されます
}
