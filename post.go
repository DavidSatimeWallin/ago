package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"time"

	stripmd "github.com/writeas/go-strip-markdown"
)

func create(postFlag string) {

	var (
		formatedDate = time.Now().Format("2006-01-02 15:04:05 Monday")
		newPost      = filepath.Join(
			getFolders().PostsFolder,
			fmt.Sprintf(
				"%d__%s.md",
				time.Now().Unix(),
				postFlag,
			),
		)
	)

	if exists(newPost) {
		fmt.Println("cannot create new post file.", newPost, "already exists")
		os.Exit(1)
	}
	f, err := os.Create(newPost)
	errIt(err, "error creating a new post file")
	defer f.Close()
	for _, v := range []string{
		strings.Title(
			strings.Replace(
				strings.Replace(
					strings.Split(newPost, "__")[1],
					".md",
					"",
					-1,
				),
				"-",
				" ",
				-1,
			),
		),
		fmt.Sprintf("published %s", formatedDate),
		"Tags: page, title, post",
		";;;;;;;",
		"This is the page header",
		"=======================",
		"Here goes the content",
	} {
		fmt.Fprintln(f, v)
	}
	fmt.Println(newPost, "created successfully")
	errIt(openFileInEditor(newPost), "")
}

func readMDFile(newPost string) string {
	b, err := ioutil.ReadFile(newPost)
	errIt(err, fmt.Sprintf("could not read in %s", newPost))
	return string(b)
}

func getFiles() []os.FileInfo {
	f, err := os.Open(getFolders().PostsFolder)
	errIt(err, "")
	files, err := f.Readdir(-1)
	f.Close()
	errIt(err, "")

	var (
		returnFiles []os.FileInfo
		keys        []int
		m           map[int]os.FileInfo = make(map[int]os.FileInfo)
	)

	for _, file := range files {
		s := strings.Split(file.Name(), "__")
		i, err := strconv.Atoi(s[0])
		errIt(err, "could not convert string to int")
		keys = append(keys, i)
		m[i] = file
	}
	sort.Sort(sort.Reverse(sort.IntSlice(keys)))
	for _, k := range keys {
		returnFiles = append(returnFiles, m[k])
	}
	return returnFiles
}

func getExcerpt(file string) string {
	if exists(file) {
		return shortenContent(parsePostContent(file))
	}
	return ""
}

func parsePostContent(file string) string {
	return stripmd.Strip(strings.Split(readMDFile(file), ";;;;;;;")[1])
}

func shortenContent(content string) string {
	if len(content) > 200 {
		content = fmt.Sprintf("%s...", content[0:200])
	}
	return strings.Replace(content, "\n", " ", -1)
}
