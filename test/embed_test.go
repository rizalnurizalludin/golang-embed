package test

import (
	"embed"
	"fmt"
	"io/fs"
	"io/ioutil"
	"testing"
)

//go:embed test.txt
var test string

func TestString(t *testing.T) {
	fmt.Println(test)
}

//go:embed logo.png
var logo []byte

func TestByte(t *testing.T) {
	err := ioutil.WriteFile("new_Logo.png", logo, fs.ModePerm)
	if err != nil {
		panic(err)
	}
}

//go:embed files/a.txt
//go:embed files/b.txt
//go:embed files/c.txt
var files embed.FS

func TestMultipleFiles(t *testing.T) {
	a, _ := files.ReadFile("files/a.txt")
	b, _ := files.ReadFile("files/b.txt")
	c, _ := files.ReadFile("files/c.txt")

	fmt.Println(string(a))
	fmt.Println(string(b))
	fmt.Println(string(c))
}

//go:embed files/*.txt
var path embed.FS

func TestPathMatcher(t *testing.T) {
	dir, _ := path.ReadDir("files")
	for _, entry := range dir {
		if !entry.IsDir() {
			fmt.Println(entry.Name())
			content, _ := path.ReadFile("files/" + entry.Name())
			fmt.Println("Content", string(content))
		}
	}
}
