
package main

import (
	"net/http"
	"sync"
	"html/template"
	"path/filepath"
	"log"
	"flag"
	"os"
	"gostudy/trace"
)

type templateHandler struct {
	// 関数を一度だけ呼び出したい場合に使用
	// https://play.golang.org/p/vrZ6thRhct
	once	sync.Once
	filename	string
	//
	temp1	*template.Template
}

func (t *templateHandler) ServeHTTP (w http.ResponseWriter, r *http.Request){
	t.once.Do(func(){
		t.temp1 =
			template.Must(template.ParseFiles(filepath.Join("templates",
				t.filename)))
	})
	t.temp1.Execute(w, r)
}



func main(){
	var addr = flag.String("addr", ":8080", "アプリケーションのアドレス")
	flag.Parse() // フラグを解釈する

	r := newRoom()
	r.tracer = trace.New(os.Stdout)
	http.Handle("/", &templateHandler{filename: "chat.html"})
	http.Handle("/room", r)
	// チャットルーム開始します
	go r.run()
	// webサーバーを起動
	log.Println("webサーバーを開始します。ポート：", *addr)
	if err := http.ListenAndServe(*addr, nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

