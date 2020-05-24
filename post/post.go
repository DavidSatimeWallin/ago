package post

import (
	"fmt"
	"log"
	"os"

	"github.com/dvwallin/ago/util"
)

// Create is used to generate a new empty post file
func Create(postsFolder string, formatedDate string, newAbsolutePostPath string) {
	if !util.FolderExists(postsFolder) {
		os.MkdirAll(postsFolder, os.ModePerm)
	}

	// create new post file
	c := []string{
		"This is the page title",
		fmt.Sprintf("Published %s", formatedDate),
		"Keywords: page, title, post",
		";;;;;;;",
		"This is the page header",
		"=======================",
		"Here goes the content",
	}
	if util.FileExists(newAbsolutePostPath) {
		fmt.Println("cannot create new post file.", newAbsolutePostPath, "already exists")
		os.Exit(1)
	}
	f, err := os.Create(newAbsolutePostPath)
	if err != nil {
		log.Println("error creating a new post file", err)
		f.Close()
		os.Exit(1)
	}
	for _, v := range c {
		fmt.Fprintln(f, v)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(newAbsolutePostPath, "created successfully")

	util.OpenFileInEditor(newAbsolutePostPath)
}
