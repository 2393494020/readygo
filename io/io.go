package io

import (
	"fmt"
	"os"
	_ "io/ioutil"
)

func OpenFileLocal(path string) {
	file, err := os.Open(path)
	if err != nil {
		fmt.Printf("error=%v\n", err)
	}

	fmt.Printf("file=%v\n", file)
	
	file.Close()
}


func OsWriteFile(content string)  {
	file, _ := os.OpenFile("C:\\Users\\yjyfzp\\Desktop\\dwprb.txt", os.O_CREATE | os.O_APPEND, 0644)
	file.Write([]byte(content))
	file.Close()
}