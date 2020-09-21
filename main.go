package main

import (
	"flag"
	"fmt"
	"os"
	"regexp"
)

var (
	initFlag      = flag.Bool("init", false, "run in a new folder to create a new Ago Blog!")
	transpileFlag = flag.Bool("transpile", false, "transpiles the markdown files into html")
	postFlag      = flag.String("post", "", "used to create a new post")
	helpFlag      = flag.Bool("help", false, "show help section")

	GitCommit, GitState, Version string
)

func init() {
	flag.Parse()

	verifyConfig(initFlag)
	initFolders()
}

func main() {
	if len(*postFlag) > 3 {
		if postnameIsValidFormat := regexp.MustCompile(`^[a-zA-Z0-9-.]+$`).MatchString; !postnameIsValidFormat(*postFlag) {
			fmt.Println("post names can only contain a-zA-Z0-9 . (dot) and -")
			os.Exit(1)
		}
		create(*postFlag)
		os.Exit(0)
	}
	if *transpileFlag {
		transpile()
		os.Exit(0)
	}
	if *helpFlag {
		fmt.Printf("\nVersion: ago%s %s %s\n", Version, GitCommit, GitState)
		fmt.Println(`
This is Ago Blog, a lightweight tool to generate static html blogs.

	ago -init in a new folder to create a new blog
	ago -post to create a new blog post template to edit
	ago -transpile to generate the static site from the posts
	ago -help to show this section

		`)
		os.Exit(0)
	}
}
