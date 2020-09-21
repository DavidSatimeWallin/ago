package main

import (
	"fmt"
	"path/filepath"
	"strings"
	"time"

	"github.com/gorilla/feeds"
)

func generateFeeds() {
	fullURL := fmt.Sprintf("%s://%s/", cfg.Protocol, cfg.Domain)
	now := time.Now()
	feed := &feeds.Feed{
		Title:       cfg.Title,
		Link:        &feeds.Link{Href: fullURL},
		Description: cfg.Description,
		Author:      &feeds.Author{Name: cfg.Author, Email: cfg.Email},
		Created:     now,
	}

	for _, file := range getFiles() {
		filename := filepath.Join(getFolders().PostsFolder, file.Name())
		fileContentSlice := strings.Split(readMDFile(filename), ";;;;;;;")
		headerSlice := strings.Split(fileContentSlice[0], "\n")
		feed.Items = append(feed.Items, &feeds.Item{
			Title:       headerSlice[0],
			Link:        &feeds.Link{Href: fmt.Sprintf("%sentries/%s", fullURL, strings.Replace(file.Name(), ".md", ".html", -1))},
			Description: getExcerpt(filename),
			Author:      &feeds.Author{Name: cfg.Author, Email: cfg.Email},
			Created:     now,
		})
	}

	atom, err := feed.ToAtom()
	errIt(err, "")
	createFeedFile(atom, "ago.atom")

	rss, err := feed.ToRss()
	errIt(err, "")
	createFeedFile(rss, "ago.rss")
}

func createFeedFile(content string, name string) {
	outputFile := filepath.Join(getFolders().SiteFolder, name)
	delFileIfExists(outputFile)
	createFile(outputFile, content)
}
