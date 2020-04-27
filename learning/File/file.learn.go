package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	absPath, _ := filepath.Abs("./GO/learning/File/test.txt")
	fmt.Println(absPath)

	file, err := os.Open(absPath)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	stat, err := file.Stat()
	if err != nil {
		return
	}

	bs := make([]byte, stat.Size())
	_, err = file.Read(bs)
	if err != nil {
		return
	}

	str := string(bs)
	fmt.Println(str)
}
