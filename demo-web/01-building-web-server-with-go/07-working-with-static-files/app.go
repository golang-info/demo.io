package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
)

type CustomFile struct {
	http.File
}

type CustomFileSystem struct {
	http.FileSystem
}

func (c CustomFile) Readdir(n int) (fis []os.FileInfo, err error) {
	files, err := c.File.Readdir(n)
	for _, file := range files {
		if !strings.Contains(file.Name(), ".json") {
			fmt.Println(file.Name())
			fis = append(fis, file)
		}
	}
	return
}

func (fs CustomFileSystem) Open(name string) (http.File, error) {
	if strings.Contains(name, ".json") {
		return nil, os.ErrPermission
	}
	file, err := fs.FileSystem.Open(name)
	if err != nil {
		return nil, err
	}
	return CustomFile{file}, err
}

func main() {
	fs := CustomFileSystem{http.Dir(".")}
	http.Handle("/", http.FileServer(fs))
	http.ListenAndServe(":3000", nil)
}

