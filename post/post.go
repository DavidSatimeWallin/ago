package post

import (
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/dvwallin/ago/config"
	"github.com/dvwallin/ago/util"
	stripmd "github.com/writeas/go-strip-markdown"
)

// Create is used to generate a new empty post file
func Create(formatedDate string, newAbsolutePostPath string) {
	t := strings.Split(newAbsolutePostPath, "__")
	if util.FileExists(newAbsolutePostPath) {
		fmt.Println("cannot create new post file.", newAbsolutePostPath, "already exists")
		os.Exit(1)
	}
	f, err := os.Create(newAbsolutePostPath)
	defer f.Close()
	util.ErrIt(err, "error creating a new post file")
	for _, v := range []string{
		strings.Title(strings.Replace(strings.Replace(t[1], ".md", "", -1), "-", " ", -1)),
		fmt.Sprintf("Published %s", formatedDate),
		"Tags: page, title, post",
		";;;;;;;",
		"This is the page header",
		"=======================",
		"Here goes the content",
	} {
		fmt.Fprintln(f, v)
		util.ErrIt(err, "")
	}
	err = f.Close()
	util.ErrIt(err, "")
	fmt.Println(newAbsolutePostPath, "created successfully")
	util.OpenFileInEditor(newAbsolutePostPath)
}

// ReadMDFile reads in an entire file into a single string
func ReadMDFile(newAbsolutePostPath string) string {
	b, err := ioutil.ReadFile(newAbsolutePostPath)
	util.ErrIt(err, fmt.Sprintf("could not read in %s", newAbsolutePostPath))
	return string(b)
}

// GetFiles gets all post files
func GetFiles() []os.FileInfo {
	f, err := os.Open(config.GetFolders().PostsFolder)
	util.ErrIt(err, "")
	files, err := f.Readdir(-1)
	f.Close()
	util.ErrIt(err, "")

	m := make(map[int]os.FileInfo)
	var returnFiles []os.FileInfo
	var keys []int
	for _, file := range files {
		s := strings.Split(file.Name(), "__")
		i, err := strconv.Atoi(s[0])
		util.ErrIt(err, "could not convert string to int")
		keys = append(keys, i)
		m[i] = file
	}
	sort.Sort(sort.Reverse(sort.IntSlice(keys)))
	for _, k := range keys {
		returnFiles = append(returnFiles, m[k])
	}
	return returnFiles
}

// GetExcerpt returns the first 100 characters of a blog post
func GetExcerpt(file string) (content string) {
	if util.FileExists(file) {
	content = stripmd.Strip(strings.Split(ReadMDFile(file), ";;;;;;;")[1])
	if len(content) > 200 {
		content = fmt.Sprintf("%s...", content[0:200])
	}
	strings.Replace(content, "\n", " ", -1)
	}
	return
}
