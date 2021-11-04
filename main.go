package main

import (
	"log"
	"net/http"
	"text/template"
)

// Temps is template structure
type Temps struct {
	notemp *template.Template
	index  *template.Template
	hello  *template.Template
}

// Template for no-template.
func notemp() *template.Template {
	src := "<html><body><h1>NO TEMPLATE.</h1></body></html>"
	tmp,_ := template.New("index").Parse(src)
	return tmp
}

// setup template function.
func setupTemp() *Temps {
	temps := new(Temps)

	// ページがない場合の表示画面をセット
	temps.notemp = notemp()

	// set index template.
	index, er := template.ParseFiles("templates/index.html")
	if er != nil {
		index = temps.notemp
	}
	temps.index = index

	// set hello template.
	hello, er := template.ParseFiles("templates/hello.html")
	if er != nil {
		hello = temps.notemp
	}
	temps.hello = hello

	return temps
}

// index handler
func index(w http.ResponseWriter, rq *http.Request, tmp *template.Template) {
	// Execute: ファイル出力 (この場合templatesのhtmlを表示)
	er := tmp.Execute(w, nil)
	if er != nil {
		log.Fatal(er)
	}
}

// hello handler
func hello(w http.ResponseWriter, rq *http.Request, tmp *template.Template) {
	// templateに渡されるデータを構造体として定義
	item := struct {
		Title   string
		Message string
	}{
		Title: "Send values",
		Message: "This is Sample message. <br>これはサンプルです。",
	}
	// Execute: ファイル出力 (この場合templatesのhtmlを表示)
	// 第二引数：templatesに渡すデータ
	er := tmp.Execute(w, item)
	if er != nil {
		log.Fatal(er)
	}
}


// main program
func main() {
	temps := setupTemp()
	// index handler
	http.HandleFunc("/", func(w http.ResponseWriter, rq *http.Request) {
		hello(w, rq, temps.index)
	})

	// hello handling
	http.HandleFunc("/hello", func(w http.ResponseWriter, rq *http.Request) {
		hello(w, rq, temps.hello)
	})

	http.ListenAndServe("", nil)
}