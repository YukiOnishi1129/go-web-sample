package main

// func main() {
// 	/*
// 	* FileServer: ファイルを表示する
// 	* http.Dir("."): 引数にはパスを指定
// 	*/
// 	// http.ListenAndServe("", http.FileServer((http.Dir("."))))

// 	// html := `<html>
// 	// <body>
// 	// <h1>Hello</h1>
// 	// <p>This is sample message</p>
// 	// </body>
// 	// </html>
// 	// `

// 	// テンプレートを使用
// 	// template.New().Parse(): Parseに入れる引数はhtmlのソースコード
// 	// tf, er := template.New("index").Parse(html)
// 	/*
// 	* template.ParseFiles()
// 	* ファイルパスにテンプレートのhtmlファイルを埋め込める
// 	*/
// 	tf, er := template.ParseFiles("templates/hello.html")
// 	if er != nil {
// 		// log.Fatal(er)
// 		// templateを読み込めなかった際に表示するエラーページ
// 		tf ,_ = template.New("index").Parse("<html><body><h1>NO TEMPLATE.</h1></body></html>")
// 	}
// 	hh := func(w http.ResponseWriter, rq *http.Request) {
// 		// w.Write([]byte("Hello, This is Go-server!!"))
// 		// htmlも出力できる

// 		// Execute: templateのメソッド。ソースコードをwebページとして出力
// 		er = tf.Execute(w, nil)
// 		if er != nil {
// 			log.Fatal((er))
// 		}
// 	}

// 	/*
// 	* HandleFunc
// 	* アクセスされるアドレスごとに処理を個別に組み込む
// 	* ルーティングのようなもの？
// 	*  第一引数：アドレス
// 	*  第二引数: 処理
// 	*/
// 	http.HandleFunc("/hello", hh)

// 	// 実際webサーバーではファイルをそのまま返すことはないので、
// 	// ListenAndServeの第二引数はnilにしておくことが多い
// 	http.ListenAndServe("", nil)
// }
