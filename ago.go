package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"time"

	"github.com/dvwallin/ago/config"
	"github.com/dvwallin/ago/post"
	"github.com/dvwallin/ago/transpiler"
)

var (
	// Arg flags
	initFlag      = flag.Bool("init", false, "run in a new folder to create a new Ago Blog!")
	transpileFlag = flag.Bool("transpile", false, "transpiles the markdown files into html")
	postFlag      = flag.String("post", "", "used to create a new post")
	helpFlag      = flag.Bool("help", false, "show help section")

	runtimeUnix = time.Now().Unix()
)

func init() {
	flag.Parse()
	config.VerifyConfig(initFlag)
}

func main() {

	files := post.GetFiles()
	for _, file := range files {
		fmt.Println(file.Name())
	}

	if len(*postFlag) > 3 {
		var (
			formatedDate          = time.Now().Format("2006-01-02 15:04:05 Monday")
			newPostName           = fmt.Sprintf("%d__%s.md", runtimeUnix, *postFlag)
			newAbsolutePostPath   = filepath.Join(config.GetFolders().PostsFolder, newPostName)
			postnameIsValidFormat = regexp.MustCompile(`^[a-zA-Z0-9-.]+$`).MatchString
		)

		if !postnameIsValidFormat(*postFlag) {
			fmt.Println("blog post names can only contain a-zA-Z0-9 . (dot) and -")
			os.Exit(1)
		}
		post.Create(formatedDate, newAbsolutePostPath)
		os.Exit(0)
	}
	if *transpileFlag {
		transpiler.Run()
		os.Exit(0)
	}
	if *helpFlag {
		fmt.Println("~~~~~")
		fmt.Println("this is Ago Blog, a lightweight tool to generate static html blogs.")
		fmt.Println("use -init in a new folder to create a new blog")
		fmt.Println("use -post to create a new blog post template to edit")
		fmt.Println("use -help to show this section")
		fmt.Println("~~~~~")
		os.Exit(0)
	}
}
