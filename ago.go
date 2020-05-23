package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/davecgh/go-spew/spew"
	"github.com/dvwallin/ago/layout"
	"github.com/dvwallin/ago/types"
	"github.com/kkyr/fig"
)

var (
	initFlag = flag.Bool("init", false, "run in a new folder to create a new Ago Blog!")
	cfg      types.Config
)

func init() {
	flag.Parse()
	err := fig.Load(&cfg)
	if err != nil {
		if *initFlag {
			if !fileExists("config.yaml") {
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
	var ph types.Placeholders = types.Placeholders{
		types.Placeholder{
			Tag:   "[[TITLE]]",
			Value: "HEJHEJ",
		},
		types.Placeholder{
			Tag:   "[[DESCRIPTION]]",
			Value: "This is a desc",
		},
		types.Placeholder{
			Tag:   "[[KEYWORDS]]",
			Value: "apa,banan,cyckel",
		},
	}
	layout.GenerateHeader(ph)
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
