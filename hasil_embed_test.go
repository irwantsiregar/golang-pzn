package main

import (
	"embed"
	"fmt"
	"io/fs"
	"io/ioutil"
	"testing"
)

// go: embed version.txt
var version string

// go: embed logo.png
var logo []byte


// go: embed files/*.txt
var path embed.FS


func main(t *testing.T) {
	fmt.Println(version)


	err := ioutil.WriteFile("logo_new.png", logo, fs.ModePerm)

		dirEntries, _ := path.ReadDir("files")

	for _, entry := range dirEntries {
		if !entry.IsDir() {
			fmt.Println(entry.Name())

			file, _ := path.ReadFile("files/" + entry.Name())
			fmt.Println(string(file))
		}
	}
}
