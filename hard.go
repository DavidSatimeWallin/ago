package main

import (
	"path/filepath"
)

type hard struct {
	PostsFolder   string
	SiteFolder    string
	EntriesFolder string
	TagsFolder    string
}

func getFolders() hard {
	return hard{
		PostsFolder:   filepath.Join(".", "posts"),
		SiteFolder:    filepath.Join(".", "site"),
		EntriesFolder: filepath.Join(".", "site", "entries"),
		TagsFolder:    filepath.Join(".", "site", "tags"),
	}
}
