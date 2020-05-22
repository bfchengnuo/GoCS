package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

func main() {
	pwd, _ := os.Getwd()
	println("pwd: ", pwd)
	println()

	const path = "/Users/kerronex/Documents"

	listFile(path)
	//listFile4Filepath(path)
	//useFilepath()
}

func useFilepath() {
	//pwd,_ := os.Getwd()
	//files, _ := filepath.Glob(filepath.Join(pwd, "*"))

	// 默认取当前目录, 列出与指定的模式 pattern 完全匹配的文件或目录
	files, _ := filepath.Glob("*")
	fmt.Println(files)
}

func listFile4Filepath(path string) {
	// 大目录效率比较低
	filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		fmt.Println("path: ", path)
		// 文件或者目录名
		fmt.Println(info.Name())
		return nil
	})
}

// 列出目录下的所有文件，递归调用
func listFile(path string) {
	dir, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range dir {
		if f.IsDir() {
			listFile(path + "/" + f.Name())
		} else {
			fmt.Println(f.Name())
		}
	}

}
