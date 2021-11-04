package main

import (
	"net/http"
)

func main() {
	/*
	* FileServer: ファイルを表示する
	* http.Dir("."): 引数にはパスを指定
	*/
	http.ListenAndServe("", http.FileServer((http.Dir("."))))
}
