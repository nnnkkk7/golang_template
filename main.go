package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob(("templates/*.html")))
}

func main() {
	var s = bufio.NewScanner(os.Stdin)
	//　入力を受け取る
	s.Scan()
	t := s.Text()
	fmt.Println(t)

	//　ファイルを作成
	f, err := os.Create("go.html")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// 作成したファイルに書き込む
	err = tpl.ExecuteTemplate(f, "index.html", t)
	if err != nil {
		log.Fatal(err)
	}

}
