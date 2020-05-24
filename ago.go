package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/dvwallin/ago/post"
	"github.com/dvwallin/ago/types"
	"github.com/dvwallin/ago/util"
	"github.com/kkyr/fig"
)

var (
	// Arg flags
	initFlag = flag.Bool("init", false, "run in a new folder to create a new Ago Blog!")
	postFlag = flag.Bool("post", false, "used to create a new post")
	helpFlag = flag.Bool("help", false, "show help section")

	// config options
	postsFolder         = filepath.Join(".", "posts")
	runtimeDate         = time.Now().Format("2006-01-02_15:04:05")
	formatedDate        = time.Now().Format("2006-01-02 15:04:05 Monday")
	newPostName         = fmt.Sprintf("%s-new-post.md", runtimeDate)
	newAbsolutePostPath = filepath.Join(postsFolder, newPostName)
	cfg                 types.Config
)

func init() {
	flag.Parse()
	err := fig.Load(&cfg)
	if err != nil {
		if *initFlag {
			if !util.FileExists("config.yaml") {
				var defaultConfig string = `domain: "ago.ofnir.xyz"
author: "Joane Doe"
email: "joane.doe@ago.ofnir.xyz"
website_name: "an Ago Blog!"
github_account: "dvwallin"
title: "an Ago Blog!"
description: "This is an awesome Ago Blog!"
keywords: "ago,blog,awesome"`
				err := ioutil.WriteFile("config.yaml", []byte(defaultConfig), 0644)
				if err != nil {
					log.Fatalln("could not create config.yaml")
				}
				fmt.Println("congratulations to your new Ago Blog!")
			} else {
				log.Fatalln("config.yaml already exists")
			}
		} else {
			log.Fatalln("can't find the config.yaml file! please run 'ago -init' to generate a new config.yaml.", err)
		}
	}
}

func main() {
	spew.Dump(cfg)
	fmt.Println(newPostName)
	if *postFlag {
		post.Create(postsFolder, formatedDate, newAbsolutePostPath)
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

	// layout.GenerateHeader(ph)
}
