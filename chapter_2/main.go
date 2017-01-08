package main

import (
	"flag"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"sync"

	"github.com/stretchr/gomniauth"
	"github.com/stretchr/gomniauth/providers/facebook"
	"github.com/stretchr/gomniauth/providers/github"
	"github.com/stretchr/gomniauth/providers/google"
)

// templは1つのテンプレートを表します
type templateHandler struct {
	once     sync.Once
	filename string
	templ    *template.Template
}

// ServerHTTPはHTTPリクエストを処理します
func (t *templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t.once.Do(func() {
		t.templ = template.Must(template.ParseFiles(filepath.Join("templates", t.filename)))
	})
	t.templ.Execute(w, r)
}

func main() {
	var addr = flag.String("host", ":8080", "アプリケーションのアドレス")
	flag.Parse()

	// Gomniauthのセットアップ
	gomniauth.SetSecurityKey("t9UHTgjUUJ")
	gomniauth.WithProviders(
		facebook.New("クライアントID", "秘密の値", "http://localhost:8080/auth/callback/facebook"),
		github.New("クライアントID", "秘密の値", "http://localhost:8080/auth/callback/github"),
		google.New("クライアントID", "秘密の値", "http://localhost:8080/auth/callback/google"),
	)

	r := newRoom()
	//r.tracer = trace.New(os.Stdout)
	http.Handle("/", &templateHandler{filename: "chat.html"})
	http.Handle("/chat", MustAuth(&templateHandler{filename: "chat.html"}))
	http.Handle("/login", &templateHandler{filename: "login.html"})
	http.HandleFunc("/auth/", loginHandler)
	http.Handle("/room", r)

	// assets設定
	http.Handle("/assets/", http.StripPrefix("/assets", http.FileServer(http.Dir("./assets/"))))

	// チャットルームの開始
	go r.run()

	log.Println("Webサーバを開始します。ポート:", *addr)
	if err := http.ListenAndServe(*addr, nil); err != nil {
		log.Fatal("ListenAndServer", err)
	}
}
