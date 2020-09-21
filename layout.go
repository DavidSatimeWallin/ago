package main

import (
	"fmt"
	"reflect"
	"strings"
)

const header = `<!DOCTYPE html><html lang="en"><head><meta charset="utf-8"><title>[[TITLE]]</title><meta name="viewport" content="width=device-width, initial-scale=1"><meta name="description" content="[[DESCRIPTION]]"><meta name="keywords" content="[[TAGS]]"><style>%%STYLE%%</style><link rel="alternate" type="application/rss+xml" title="RSS Feed for [[DOMAIN]]" href="[[PROTOCOL]]://[[DOMAIN]]/ago.rss" /><link rel="alternate" type="application/atom+xml" title="Atom Feed for [[DOMAIN]]" href="[[PROTOCOL]]://[[DOMAIN]]/ago.atom" /></head><body><header><div class="title"><h2><a href="[[PROTOCOL]]://[[DOMAIN]]">[[TITLE]]</a></h2><p><em>[[INTRO]]</em></p><hr /><nav><ul><li><a href="[[PROTOCOL]]://[[DOMAIN]]/all_entries.html">View all entries</a></li><li><a href="[[PROTOCOL]]://[[DOMAIN]]/ago.atom">Atom feed</a></li><li><a href="[[PROTOCOL]]://[[DOMAIN]]/ago.rss">RSS feed</a></li></ul></nav><hr /></div></header><div class="story-container">`

const footer = `</div><footer>generated with the <a href="https://ago.ofnir.xyz">ago blog</a> script. source code located at <a href="https://github.com/dvwallin/ago">GitHub</a>.</footer></body></html>`

func generateHeader() string {
	cfg := getCfg()
	output := parse(header, cfg)
	return output
}

func generateFooter() string {
	cfg := getCfg()
	output := parse(footer, cfg)
	return output
}

func parse(input string, cfg config) string {
	v := reflect.ValueOf(cfg)
	typeOfS := v.Type()
	for i := 0; i < v.NumField(); i++ {
		input = strings.ReplaceAll(
			input,
			fmt.Sprintf(
				"[[%s]]",
				strings.ToUpper(typeOfS.Field(i).Name),
			),
			v.Field(i).Interface().(string),
		)
	}
	return input
}
