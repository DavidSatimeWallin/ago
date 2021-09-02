package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/microcosm-cc/bluemonday"
	"github.com/russross/blackfriday/v2"
	"github.com/tdewolff/minify"
	"github.com/tdewolff/minify/html"
)

func transpile() {
	m := minify.New()
	m.AddFunc("text/html", html.Minify)
	parsedHeader := applyStyle(generateHeader())
	indexfile := filepath.Join(getFolders().SiteFolder, "index.html")
	s, err := m.String("text/html", fmt.Sprintf("%s%s%s", parsedHeader, posts(10), footer))
	errIt(err, "")
	delFileIfExists(indexfile)
	createFile(indexfile, s)
	createAllEntriesPage()
	tags := make(map[string][]string)
	for _, file := range getFiles() {
		writeSingleEntry(file)
		tags = buildTagIndex(tags, file)
	}
	writeTagFiles(tags)
	generateFeeds()
}

func applyStyle(input string) string {
	return strings.Replace(input, "%%STYLE%%", `h1{font-size:45px}h2{font-size:30px}p{font-size:16px}ul{list-style:none}ul li{display:inline;margin:15px;}ul li a{padding:5px}a{text-decoration:none;color:#0074D9}a:hover{text-decoration:underline dotted}hr{border:0;height:0;border-top:1px solid rgba(0,0,0,0.1);border-bottom:1px solid rgba(255,255,255,0.3)}.title{text-align:center;line-height:30px;height:100%}.story-container{padding:50px;}footer{padding:50px}`, -1)
}

func createAllEntriesPage() {
	m := minify.New()
	m.AddFunc("text/html", html.Minify)
	allEntriesFile := filepath.Join(getFolders().SiteFolder, "all_entries.html")
	s, err := m.String("text/html", fmt.Sprintf("%s%s%s", applyStyle(generateHeader()), posts(-1), footer))
	errIt(err, "")
	delFileIfExists(allEntriesFile)
	createFile(allEntriesFile, s)

}

func posts(limit int) (bodyContent string) {
	fullURL := fmt.Sprintf("%s://%s/", cfg.Protocol, cfg.Domain)
	files := getFiles()
	i := 0
fileLoop:
	for _, file := range files {
		if i == limit {
			break fileLoop
		}
		bodyContent = generator(bodyContent, file, fullURL)
		i++
	}
	return
}

func generator(bodyContent string, file os.FileInfo, fullURL string) string {
	filename := filepath.Join(getFolders().PostsFolder, file.Name())
	fileContentSlice := strings.Split(readMDFile(filename), ";;;;;;;")
	headerSlice := strings.Split(fileContentSlice[0], "\n")
	headerSlice[2] = linkTags(headerSlice[2])
	content := fmt.Sprintf(
		"<div><a href='%sentries/%s.html'>%s</a><small>%s</small></div>",
		fullURL,
		strings.Replace(file.Name(), ".md", "", -1),
		headerSlice[0],
		headerSlice[1],
	)
	unsafe := blackfriday.Run([]byte(content))
	return fmt.Sprintf("%s%s", bodyContent, bluemonday.UGCPolicy().SanitizeBytes(unsafe))
}

func writeSingleEntry(file os.FileInfo) {
	filePath := filepath.Join(getFolders().EntriesFolder, strings.Replace(file.Name(), ".md", ".html", -1))
	fileContentSlice := strings.Split(readMDFile(filepath.Join(getFolders().PostsFolder, file.Name())), ";;;;;;;")
	unsafe := blackfriday.Run([]byte(fileContentSlice[1]))
	headerSlice := strings.Split(fileContentSlice[0], "\n")
	headerSlice[2] = linkTags(headerSlice[2])
	m := minify.New()
	m.AddFunc("text/html", html.Minify)
	content, err := m.String("text/html", string(bluemonday.UGCPolicy().SanitizeBytes(unsafe)))
	errIt(err, "")
	fileContent := fmt.Sprintf(
		"%s%s%s%s%s",
		applyStyle(generateHeader()),
		fmt.Sprintf("<small>%s</small><hr />", headerSlice[1]),
		content,
		fmt.Sprintf(
			"<hr /><p>Written by %s ( %s )</p><p>%s</p>",
			cfg.Author,
			strings.Replace(
				cfg.Email,
				"@",
				"[_AT_]",
				-1,
			),
			headerSlice[2],
		),
		generateFooter(),
	)
	createFile(filePath, fileContent)
}

func buildTagIndex(tags map[string][]string, file os.FileInfo) map[string][]string {
	fileContentSlice := strings.Split(
		readMDFile(
			filepath.Join(
				getFolders().PostsFolder,
				file.Name(),
			),
		),
		";;;;;;;",
	)
	headerSlice := strings.Split(fileContentSlice[0], "\n")
	tagSlice := strings.Split(strings.Replace(strings.Replace(headerSlice[2], " ", "", -1), "Tags:", "", -1), ",")
	for _, tag := range tagSlice {
		tags[tag] = append(
			tags[tag],
			fmt.Sprintf(
				"%s://%s/entries/%s",
				cfg.Protocol,
				cfg.Domain,
				strings.Replace(
					file.Name(),
					".md",
					".html",
					-1,
				),
			),
		)
	}
	return tags
}

func writeTagFiles(tags map[string][]string) {
	for tag, posts := range tags {
		file := filepath.Join(getFolders().TagsFolder, fmt.Sprintf("%s.html", tag))
		delFileIfExists(file)
		for i, p := range posts {
			t := strings.Split(p, "__")
			title := strings.Title(strings.Replace(strings.Replace(t[1], ".html", "", -1), "-", " ", -1))
			posts[i] = fmt.Sprintf("<li><a href='%s'>%s</a></li>", p, title)
		}
		content := fmt.Sprintf(
			"%s%s%s",
			applyStyle(generateHeader()),
			bluemonday.UGCPolicy().SanitizeBytes(
				blackfriday.Run(
					[]byte(
						fmt.Sprintf("<ul>%s</ul>",
							strings.Join(posts, ""),
						),
					),
				),
			),
			generateFooter(),
		)
		createFile(file, content)
	}
}

func linkTags(tagString string) string {
	fullURL := fmt.Sprintf("%s://%s/", cfg.Protocol, cfg.Domain)
	tagString = strings.Replace(tagString, " ", "", -1)
	tagSlice := strings.Split(tagString, ":")
	wordSlice := strings.Split(tagSlice[1], ",")
	for key, word := range wordSlice {
		wordSlice[key] = fmt.Sprintf("<a href='%stags/%s.html'>%s</a>", fullURL, word, word)
	}
	tagSlice[1] = strings.Join(wordSlice, ", ")
	tagString = strings.Join(tagSlice, ": ")
	return tagString
}
