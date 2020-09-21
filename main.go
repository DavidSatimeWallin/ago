package main

import (
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"

	"github.com/kkyr/fig"
)

type (
	config struct {
		Domain      string `fig:"domain" default:"ago.ofnir.xyz"`
		Protocol    string `fig:"protocol" default:"https"`
		Author      string `fig:"author" default:"Joane Doe"`
		Email       string `fig:"email" default:"joane.doe@ago.ofnir.xyz"`
		Title       string `fig:"title" default:"an Ago Blog!"`
		Description string `fig:"description" default:"This is an awesome Ago Blog!"`
		Tags        string `fig:"tags" default:"ago,blog,awesome"`
		Intro       string `fig:"intro" default:"You should have a small intro here to describe a little bit about yourself and the purpose of the blog"`
	}

	hard struct {
		PostsFolder   string
		SiteFolder    string
		EntriesFolder string
		TagsFolder    string
	}
)

var (
	initFlag      = flag.Bool("init", false, "run in a new folder to create a new Ago Blog!")
	transpileFlag = flag.Bool("transpile", false, "transpiles the markdown files into html")
	postFlag      = flag.String("post", "", "used to create a new post")
	helpFlag      = flag.Bool("help", false, "show help section")
	cfg           config

	GitCommit, GitState, Version string
)

func main() {
	flag.Parse()

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

	verifyConfig(initFlag)
	initFolders()

	err := fig.Load(&cfg)
	errIt(err, "could not load config")
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
}

func initFolders() {
	for _, folder := range []string{
		getFolders().PostsFolder,
		getFolders().EntriesFolder,
		getFolders().SiteFolder,
		getFolders().TagsFolder,
	} {
		if !exists(folder) {
			err := os.MkdirAll(folder, os.ModePerm)
			errIt(err, fmt.Sprintf("could not create %s", folder))
		}
	}
}

func verifyConfig(initFlag *bool) {
	err := fig.Load(&cfg)
	if err != nil {
		if *initFlag {
			if !exists("config.yaml") {
				var defaultConfig string = `domain: "ago.ofnir.xyz"
protocol: "https"
author: "Joane Doe"
email: "joane.doe@ago.ofnir.xyz"
title: "an Ago Blog!"
description: "This is an awesome Ago Blog!"
tags: "ago,blog,awesome"
intro: "You should have a small intro here to describe a little bit about yourself and the purpose of the blog"`
				err := ioutil.WriteFile("config.yaml", []byte(defaultConfig), 0644)
				errIt(err, "could not create config.yaml")
				fmt.Println("blog initialized..")
			} else {
				errIt(errors.New("config.yaml already exists"), "")
			}
		} else {
			errIt(err, "can't find the config.yaml file! please run 'ago -init' to generate a new config.yaml.")
		}
	}
}

func getFolders() hard {
	return hard{
		PostsFolder:   filepath.Join(".", "posts"),
		SiteFolder:    filepath.Join(".", "site"),
		EntriesFolder: filepath.Join(".", "site", "entries"),
		TagsFolder:    filepath.Join(".", "site", "tags"),
	}
}
