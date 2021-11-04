package main

import "net/http"

func main() {
	/*
	* FileServer: ファイルを表示する
	* http.Dir("."): 引数にはパスを指定
	*/
	// http.ListenAndServe("", http.FileServer((http.Dir("."))))

	hh := func(w http.ResponseWriter, rq *http.Request) {
		w.Write([]byte("Hello, This is Go-server!!"))
	}

	/*
	* HandleFunc
	* アクセスされるアドレスごとに処理を個別に組み込む
	* ルーティングのようなもの？
	*  第一引数：アドレス
	*  第二引数: 処理
	*/
	http.HandleFunc("/hello", hh)

	// 実際webサーバーではファイルをそのまま返すことはないので、
	// ListenAndServeの第二引数はnilにしておくことが多い
	http.ListenAndServe("", nil)
}
