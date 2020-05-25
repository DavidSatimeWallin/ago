package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
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
)

func init() {
	flag.Parse()
	config.VerifyConfig(initFlag)
	config.InitFolders()
}

func main() {
	if len(*postFlag) > 3 {
		var (
			formatedDate          = time.Now().Format("2006-01-02 15:04:05 Monday")
			newPostName           = fmt.Sprintf("%d__%s.md", time.Now().Unix(), *postFlag)
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
		c := exec.Command("clear")
		c.Stdout = os.Stdout
		c.Run()
		fmt.Println("\nThis is Ago Blog, a lightweight tool to generate static html blogs.")
		fmt.Printf("\n")
		fmt.Println("\tuse -init in a new folder to create a new blog")
		fmt.Println("\tuse -post to create a new blog post template to edit")
		fmt.Println("\tuse -help to show this section")
		fmt.Printf("\n\n")
		os.Exit(0)
	}
}
