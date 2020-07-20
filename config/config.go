package config

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/dvwallin/ago/util"
	"github.com/kkyr/fig"
)

// Config contains everything needed to run the blog
type Config struct {
	Domain      string `fig:"domain" default:"ago.ofnir.xyz"`
	Protocol    string `fig:"protocol" default:"https"`
	Author      string `fig:"author" default:"Joane Doe"`
	Email       string `fig:"email" default:"joane.doe@ago.ofnir.xyz"`
	Title       string `fig:"title" default:"an Ago Blog!"`
	Description string `fig:"description" default:"This is an awesome Ago Blog!"`
	Tags        string `fig:"tags" default:"ago,blog,awesome"`
	Intro       string `fig:"intro" default:"You should have a small intro here to describe a little bit about yourself and the purpose of the blog"`
}

var cfg Config

// InitFolders is for setting up needed folder-structure
func InitFolders() {
	for _, folder := range []string{
		GetFolders().PostsFolder,
		GetFolders().EntriesFolder,
		GetFolders().SiteFolder,
		GetFolders().TagsFolder,
	} {
		if !util.Exists(folder) {
			err := os.MkdirAll(folder, os.ModePerm)
			util.ErrIt(err, fmt.Sprintf("could not create %s", folder))
		}
	}
}

// VerifyConfig verifies that we have a config file
func VerifyConfig(initFlag *bool) {
	err := fig.Load(&cfg)
	if err != nil {
		if *initFlag {
			if !util.Exists("config.yaml") {
				var defaultConfig string = `domain: "ago.ofnir.xyz"
protocol: "https"
author: "Joane Doe"
email: "joane.doe@ago.ofnir.xyz"
title: "an Ago Blog!"
description: "This is an awesome Ago Blog!"
tags: "ago,blog,awesome"
intro: "You should have a small intro here to describe a little bit about yourself and the purpose of the blog"`
				err := ioutil.WriteFile("config.yaml", []byte(defaultConfig), 0644)
				util.ErrIt(err, "could not create config.yaml")
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

// GetCfg gives us the config values
func GetCfg() Config {
	err := fig.Load(&cfg)
	util.ErrIt(err, "could not load config")
	return cfg
}
