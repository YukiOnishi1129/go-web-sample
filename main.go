package main

import (
	"log"
	"net/http"
	"text/template"

	"github.com/gorilla/sessions"
)

// var flg bool = true

// セッションのストアを作成 (セッションの値を保持する構造体)
// CookieStore: クッキー情報をもとにストアを生成
var cs *sessions.CookieStore = sessions.NewCookieStore([]byte("secret-key-12345"))

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
	
	// rq.FormValue: getパラメーターを取得
	// id := rq.FormValue("id")
	// nm := rq.FormValue("name")
	// msg := "id: " + id + ", Name: " + nm

	msg := "login name and password"

	// sessionを用意
	// <Store>.get(キー名)
	ses, _ := cs.Get(rq, "hello-session")

	if rq.Method == "POST" {
		// セッションに値を保管
		ses.Values["login"] = nil
		ses.Values["name"] = nil
		// PostFormValue: postパラメーターを取得
		nm := rq.PostFormValue("name")
		pw := rq.PostFormValue("pass")
		// nameとpasswordが同じならsessionに保持
		if nm == pw {
			ses.Values["login"] = true
			ses.Values["name"] = nm
		}
		// セッションの保管
		ses.Save(rq, w)
	}

	

	flg, _ := ses.Values["login"].(bool)
	lname, _ := ses.Values["name"].(string)

	if flg {
		msg = "logined: " + lname
	}
	
	// templateに渡されるデータを構造体として定義
	item := struct {
		Title    string
		Message  string
	}{
		Title: "Session",
		Message: msg,
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