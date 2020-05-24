package post

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/dvwallin/ago/config"
	"github.com/dvwallin/ago/util"
)

// Create is used to generate a new empty post file
func Create(formatedDate string, newAbsolutePostPath string) {
	if !util.FolderExists(config.GetFolders().PostsFolder) {
		os.MkdirAll(config.GetFolders().PostsFolder, os.ModePerm)
	}

	t := strings.Split(newAbsolutePostPath, "__")
	title := strings.Title(strings.Replace(strings.Replace(t[1], ".md", "", -1), "-", " ", -1))

	// create new post file
	c := []string{
		title,
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
	defer f.Close()
	if err != nil {
		log.Println("error creating a new post file", err)
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

// ReadMDFile reads in an entire file into a single string
func ReadMDFile(newAbsolutePostPath string) string {
	b, err := ioutil.ReadFile(newAbsolutePostPath)
	if err != nil {
		log.Print("could not read in", newAbsolutePostPath, err)
	}
	return string(b)
}

// GetFiles gets all post files
func GetFiles() []os.FileInfo {
	f, err := os.Open(config.GetFolders().PostsFolder)
	if err != nil {
		log.Fatal(err)
	}
	files, err := f.Readdir(-1)
	f.Close()
	if err != nil {
		log.Fatal(err)
	}

	m := make(map[int]os.FileInfo)
	var returnFiles []os.FileInfo
	var keys []int
	for _, file := range files {
		s := strings.Split(file.Name(), "__")
		i, err := strconv.Atoi(s[0])
		if err != nil {
			log.Println("could not convert string to int", err)
		}
		keys = append(keys, i)
		m[i] = file
	}
	sort.Sort(sort.Reverse(sort.IntSlice(keys)))
	for _, k := range keys {
		returnFiles = append(returnFiles, m[k])
	}
	return returnFiles
}
