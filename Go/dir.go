package main

import (
	"fmt"
	"io/ioutil"
	"path"
)

var pathS [string]string

func readdir(paths string) {
	files, _ := ioutil.ReadDir(paths)
	for _, f := range files {
		fmt.Println(path.Join(paths, f.Name()))
		fmt.Println(f.IsDir())
		if f.IsDir() {
			fmt.Println(path.Join(paths, f.Name()))
			fmt.Println("--------------")
			readdir(path.Join(paths, f.Name()))
		} else {
			pathS = append(pathS, path.Join(paths, f.Name()))
		}
	}
}

func main() {
	readdir("./")
	fmt.Println(pathS)
}
