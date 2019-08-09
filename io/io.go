package io

import (
	"os"
	_ "io/ioutil"
)

func OsWriteFile(content string)  {
	file, _ := os.OpenFile("C:\\Users\\yjyfzp\\Desktop\\dwprb.txt", os.O_CREATE | os.O_APPEND, 0644)
	file.Write([]byte(content))
	file.Close()
}