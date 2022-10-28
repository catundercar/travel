package main

import (
	"embed"
	"fmt"
	"io/fs"
	"log"
	"net/http"
)

//go:embed hello.txt
var hello string

//go:embed files
var files embed.FS

func main()  {
	fmt.Println(hello)

	//Sub返回与以files的目录为根的子树相对应的FS。
	fs.Sub(files, "files")

	if err := http.ListenAndServe(":8000", http.FileServer(http.FS(files))); err != nil {
		log.Fatal(err)
	}
}
