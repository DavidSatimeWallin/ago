package feed

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/dvwallin/ago/config"
	"github.com/dvwallin/ago/post"
	"github.com/dvwallin/ago/util"
	"github.com/gorilla/feeds"
)

// GenerateFeeds are the main function to initiate the generation of RSS and Atom
func GenerateFeeds() {
	cfg := config.GetCfg()
	fullURL := fmt.Sprintf("%s://%s/", cfg.Protocol, cfg.Domain)
	now := time.Now()
	feed := &feeds.Feed{
		Title:       cfg.Title,
		Link:        &feeds.Link{Href: fullURL},
		Description: cfg.Description,
		Author:      &feeds.Author{Name: cfg.Author, Email: cfg.Email},
		Created:     now,
	}

	files := post.GetFiles()
	for _, file := range files {
		filename := filepath.Join(config.GetFolders().PostsFolder, file.Name())
		fileContentSlice := strings.Split(post.ReadMDFile(filename), ";;;;;;;")
		headerSlice := strings.Split(fileContentSlice[0], "\n")
		feed.Items = append(feed.Items, &feeds.Item{
			Title:       headerSlice[0],
			Link:        &feeds.Link{Href: fmt.Sprintf("%sentries/%s", fullURL, strings.Replace(file.Name(), ".md", ".html", -1))},
			Description: post.GetExcerpt(filename),
			Author:      &feeds.Author{Name: cfg.Author, Email: cfg.Email},
			Created:     now,
		})
	}

	atom, err := feed.ToAtom()
	util.ErrIt(err, "")
	createFeedFile(atom, "ago.atom")

	rss, err := feed.ToRss()
	util.ErrIt(err, "")
	createFeedFile(rss, "ago.rss")
}

func createFeedFile(content string, name string) {
	outputFile := filepath.Join(config.GetFolders().SiteFolder, name)
	if util.FileExists(outputFile) {
		err := os.Remove(outputFile)
		util.ErrIt(err, "")
	}
	f, err := os.Create(outputFile)
	util.ErrIt(err, "")
	defer f.Close()

	_, err = f.WriteString(content)
	util.ErrIt(err, "")
}
