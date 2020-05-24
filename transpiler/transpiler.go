package transpiler

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/dvwallin/ago/config"
	"github.com/dvwallin/ago/feed"
	"github.com/dvwallin/ago/layout"
	"github.com/dvwallin/ago/post"
	"github.com/microcosm-cc/bluemonday"
	"gopkg.in/russross/blackfriday.v2"

	"github.com/dvwallin/ago/tmpl"
	"github.com/dvwallin/ago/util"
	"github.com/tdewolff/minify"
	"github.com/tdewolff/minify/html"
)

// Run - lets execute some transpiltaion shall we?
func Run() {
	m := minify.New()
	m.AddFunc("text/html", html.Minify)

	parsedHeader := layout.GenerateHeader()
	indexfile := filepath.Join(config.GetFolders().SiteFolder, "index.html")
	body := posts(10)

	generateCSSFile()

	s, err := m.String("text/html", fmt.Sprintf("%s%s%s", parsedHeader, body, tmpl.Footer))
	util.ErrIt(err, "")
	if util.FileExists(indexfile) {
		err := os.Remove(indexfile)
		util.ErrIt(err, "")
	}
	f, err := os.Create(indexfile)
	defer f.Close()

	_, err = f.WriteString(s)
	util.ErrIt(err, "")

	createAllEntriesPage()

	files := post.GetFiles()
	tags := make(map[string][]string)
	for _, file := range files {
		writeSingleEntry(file)
		tags = buildTagIndex(tags, file)
	}

	writeTagFiles(tags)

	feed.GenerateFeeds()

}

func createAllEntriesPage() {
	m := minify.New()
	m.AddFunc("text/html", html.Minify)

	parsedHeader := layout.GenerateHeader()
	allEntriesFile := filepath.Join(config.GetFolders().SiteFolder, "all_entries.html")
	body := posts(-1)

	generateCSSFile()

	s, err := m.String("text/html", fmt.Sprintf("%s%s%s", parsedHeader, body, tmpl.Footer))
	util.ErrIt(err, "")
	if util.FileExists(allEntriesFile) {
		err := os.Remove(allEntriesFile)
		util.ErrIt(err, "")
	}
	f, err := os.Create(allEntriesFile)
	defer f.Close()

	_, err = f.WriteString(s)
	util.ErrIt(err, "")

}

func generateCSSFile() {
	stylefile := filepath.Join(config.GetFolders().SiteFolder, "ago.css")
	if util.FileExists(stylefile) {
		err := os.Remove(stylefile)
		util.ErrIt(err, "")
	}
	f, err := os.Create(stylefile)
	defer f.Close()

	_, err = f.WriteString(tmpl.Style)
	util.ErrIt(err, "")
}

func posts(limit int) (bodyContent string) {
	cfg := config.GetCfg()
	fullURL := fmt.Sprintf("%s://%s/", cfg.Protocol, cfg.Domain)

	files := post.GetFiles()

	if limit < 0 {
		for _, file := range files {
			bodyContent = generator(bodyContent, file, fullURL)
		}
	} else {
		i := 0
	fileLoop:
		for _, file := range files {
			if i == limit {
				break fileLoop
			}
			bodyContent = generator(bodyContent, file, fullURL)
			i++
		}
	}

	return
}

func generator(bodyContent string, file os.FileInfo, fullURL string) string {
	filename := filepath.Join(config.GetFolders().PostsFolder, file.Name())
	fileContentSlice := strings.Split(post.ReadMDFile(filename), ";;;;;;;")
	headerSlice := strings.Split(fileContentSlice[0], "\n")
	headerSlice[2] = linkTags(headerSlice[2])
	content := fmt.Sprintf(
		"<h2><a href='%sentries/%s.html'>%s</a></h2><small>%s</small><p>%s</p><p>%s</p>",
		fullURL,
		strings.Replace(file.Name(), ".md", "", -1),
		headerSlice[0],
		headerSlice[1],
		post.GetExcerpt(filename),
		headerSlice[2],
	)
	unsafe := blackfriday.Run([]byte(content))
	return fmt.Sprintf("%s%s", bodyContent, bluemonday.UGCPolicy().SanitizeBytes(unsafe))
}

func writeSingleEntry(file os.FileInfo) {
	cfg := config.GetCfg()
	filePath := filepath.Join(config.GetFolders().EntriesFolder, strings.Replace(file.Name(), ".md", ".html", -1))
	if util.FileExists(filePath) {
		err := os.Remove(filePath)
		util.ErrIt(err, "")
	}
	fileContentSlice := strings.Split(post.ReadMDFile(filepath.Join(config.GetFolders().PostsFolder, file.Name())), ";;;;;;;")
	unsafe := blackfriday.Run([]byte(fileContentSlice[1]))
	headerSlice := strings.Split(fileContentSlice[0], "\n")
	headerSlice[2] = linkTags(headerSlice[2])
	fileContent := fmt.Sprintf(
		"%s%s%s%s%s",
		layout.GenerateHeader(),
		fmt.Sprintf("<small>%s</small><hr />", headerSlice[1]),
		bluemonday.UGCPolicy().SanitizeBytes(unsafe),
		fmt.Sprintf("<hr /><p>Written by %s ( %s )</p><p>%s</p>", cfg.Author, strings.Replace(cfg.Email, "@", "[_AT_]", -1), headerSlice[2]),
		layout.GenerateFooter(),
	)
	f, err := os.Create(filePath)
	util.ErrIt(err, "")
	defer f.Close()

	_, err = f.WriteString(fileContent)
	util.ErrIt(err, "")
}

func buildTagIndex(tags map[string][]string, file os.FileInfo) map[string][]string {
	cfg := config.GetCfg()
	fileContentSlice := strings.Split(post.ReadMDFile(filepath.Join(config.GetFolders().PostsFolder, file.Name())), ";;;;;;;")
	headerSlice := strings.Split(fileContentSlice[0], "\n")
	tagSlice := strings.Split(strings.Replace(strings.Replace(headerSlice[2], " ", "", -1), "Keywords:", "", -1), ",")
	for _, tag := range tagSlice {
		tags[tag] = append(tags[tag], fmt.Sprintf("%s://%s/entries/%s", cfg.Protocol, cfg.Domain, strings.Replace(file.Name(), ".md", ".html", -1)))
	}
	return tags
}

func writeTagFiles(tags map[string][]string) {
	for tag, posts := range tags {
		filePath := filepath.Join(config.GetFolders().TagsFolder, fmt.Sprintf("%s.html", tag))
		if util.FileExists(filePath) {
			err := os.Remove(filePath)
			util.ErrIt(err, "")
		}
		for i, p := range posts {
			t := strings.Split(p, "__")
			title := strings.Title(strings.Replace(strings.Replace(t[1], ".html", "", -1), "-", " ", -1))
			posts[i] = fmt.Sprintf("<li><a href='%s'>%s</a></li>", p, title)
		}
		unsafe := blackfriday.Run([]byte(fmt.Sprintf("<ul>%s</ul>", strings.Join(posts, ""))))
		fileContent := fmt.Sprintf(
			"%s%s%s",
			layout.GenerateHeader(),
			bluemonday.UGCPolicy().SanitizeBytes(unsafe),
			layout.GenerateFooter(),
		)
		f, err := os.Create(filePath)
		util.ErrIt(err, "")
		defer f.Close()

		_, err = f.WriteString(fileContent)
		util.ErrIt(err, "")
	}
}

func linkTags(tagString string) string {
	cfg := config.GetCfg()
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
