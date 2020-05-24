package config

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/dvwallin/ago/agotypes"
	"github.com/dvwallin/ago/util"
	"github.com/kkyr/fig"
)

var cfg agotypes.Config

// InitFolders is for setting up needed folder-structure
func InitFolders() {
	createIfNotExists(GetFolders().PostsFolder)
	createIfNotExists(GetFolders().EntriesFolder)
	createIfNotExists(GetFolders().SiteFolder)
	createIfNotExists(GetFolders().TagsFolder)
}

func createIfNotExists(folder string) {
	if !util.FolderExists(folder) {
		os.MkdirAll(folder, os.ModePerm)
	}
}

// VerifyConfig verifies that we have a config file
func VerifyConfig(initFlag *bool) {
	err := fig.Load(&cfg)
	if err != nil {
		if *initFlag {
			if !util.FileExists("config.yaml") {
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
func GetCfg() agotypes.Config {
	err := fig.Load(&cfg)
	util.ErrIt(err, "could not load config")
	return cfg
}
