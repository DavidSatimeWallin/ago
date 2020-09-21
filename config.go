package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/kkyr/fig"
)

type config struct {
	Domain      string `fig:"domain" default:"ago.ofnir.xyz"`
	Protocol    string `fig:"protocol" default:"https"`
	Author      string `fig:"author" default:"Joane Doe"`
	Email       string `fig:"email" default:"joane.doe@ago.ofnir.xyz"`
	Title       string `fig:"title" default:"an Ago Blog!"`
	Description string `fig:"description" default:"This is an awesome Ago Blog!"`
	Tags        string `fig:"tags" default:"ago,blog,awesome"`
	Intro       string `fig:"intro" default:"You should have a small intro here to describe a little bit about yourself and the purpose of the blog"`
}

var cfg config

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
				fmt.Println("congratulations to your new Ago Blog!")
			} else {
				fmt.Println("config.yaml already exists")
				os.Exit(1)
			}
		} else {
			fmt.Println("can't find the config.yaml file! please run 'ago -init' to generate a new config.yaml.", err)
			os.Exit(1)
		}
	}
}

func getCfg() config {
	err := fig.Load(&cfg)
	errIt(err, "could not load config")
	return cfg
}
