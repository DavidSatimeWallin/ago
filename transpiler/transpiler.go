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
	"github.com/dvwallin/ago/util"
	"github.com/microcosm-cc/bluemonday"
	"github.com/russross/blackfriday/v2"
	"github.com/tdewolff/minify"
	"github.com/tdewolff/minify/html"
)

// Run - lets execute some transpiltaion shall we?
func Run() {
	m := minify.New()
	m.AddFunc("text/html", html.Minify)
	parsedHeader := applyStyle(layout.GenerateHeader())
	indexfile := filepath.Join(config.GetFolders().SiteFolder, "index.html")
	s, err := m.String("text/html", fmt.Sprintf("%s%s%s", parsedHeader, posts(10), layout.Footer))
	util.ErrIt(err, "")
	util.DelFileIfExists(indexfile)
	util.GenerateFile(indexfile, s)
	createAllEntriesPage()
	tags := make(map[string][]string)
	for _, file := range post.GetFiles() {
		writeSingleEntry(file)
		tags = buildTagIndex(tags, file)
	}
	writeTagFiles(tags)
	feed.GenerateFeeds()
}

func applyStyle(input string) string {

	return strings.Replace(input, "%%STYLE%%", `h1{font-size:45px}h2{font-size:30px}p{font-size:16px}ul{list-style:none}ul li{display:inline;margin:15px;}ul li a{padding:5px}a{text-decoration:none;color:#0074D9}a:hover{text-decoration:underline dotted}hr{width:500px;border:0;height:0;border-top:1px solid rgba(0,0,0,0.1);border-bottom:1px solid rgba(255,255,255,0.3)}.title{text-align:center;line-height:30px;height:100%;max-width:600px;margin:0 auto}.story-container{max-width:600px;margin:50px auto;padding:50px;-moz-box-shadow:rgba(0,0,0,0.1) 0 10px 30px;-webkit-box-shadow:rgba(0,0,0,0.1) 0 10px 30px;box-shadow:rgba(0,0,0,0.1) 0 10px 30px}footer{max-width:600px;margin:0 auto;padding:5px}`, -1)
}

func createAllEntriesPage() {
	m := minify.New()
	m.AddFunc("text/html", html.Minify)
	parsedHeader := applyStyle(layout.GenerateHeader())
	allEntriesFile := filepath.Join(config.GetFolders().SiteFolder, "all_entries.html")
	body := posts(-1)
	s, err := m.String("text/html", fmt.Sprintf("%s%s%s", parsedHeader, body, layout.Footer))
	util.ErrIt(err, "")
	util.DelFileIfExists(allEntriesFile)
	util.GenerateFile(allEntriesFile, s)

}

func posts(limit int) (bodyContent string) {
	cfg := config.GetCfg()
	fullURL := fmt.Sprintf("%s://%s/", cfg.Protocol, cfg.Domain)
	files := post.GetFiles()
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
	fileContentSlice := strings.Split(post.ReadMDFile(filepath.Join(config.GetFolders().PostsFolder, file.Name())), ";;;;;;;")
	unsafe := blackfriday.Run([]byte(fileContentSlice[1]))
	headerSlice := strings.Split(fileContentSlice[0], "\n")
	headerSlice[2] = linkTags(headerSlice[2])
	m := minify.New()
	m.AddFunc("text/html", html.Minify)
	content, err := m.String("text/html", string(bluemonday.UGCPolicy().SanitizeBytes(unsafe)))
	util.ErrIt(err, "")
	fileContent := fmt.Sprintf(
		"%s%s%s%s%s",
		applyStyle(layout.GenerateHeader()),
		fmt.Sprintf("<small>%s</small><hr />", headerSlice[1]),
		content,
		fmt.Sprintf("<hr /><p>Written by %s ( %s )</p><p>%s</p>", cfg.Author, strings.Replace(cfg.Email, "@", "[_AT_]", -1), headerSlice[2]),
		layout.GenerateFooter(),
	)
	util.GenerateFile(filePath, fileContent)
}

func buildTagIndex(tags map[string][]string, file os.FileInfo) map[string][]string {
	cfg := config.GetCfg()
	fileContentSlice := strings.Split(post.ReadMDFile(filepath.Join(config.GetFolders().PostsFolder, file.Name())), ";;;;;;;")
	headerSlice := strings.Split(fileContentSlice[0], "\n")
	tagSlice := strings.Split(strings.Replace(strings.Replace(headerSlice[2], " ", "", -1), "Tags:", "", -1), ",")
	for _, tag := range tagSlice {
		tags[tag] = append(tags[tag], fmt.Sprintf("%s://%s/entries/%s", cfg.Protocol, cfg.Domain, strings.Replace(file.Name(), ".md", ".html", -1)))
	}
	return tags
}

func writeTagFiles(tags map[string][]string) {
	for tag, posts := range tags {
		filePath := filepath.Join(config.GetFolders().TagsFolder, fmt.Sprintf("%s.html", tag))
		util.DelFileIfExists(filePath)
		for i, p := range posts {
			t := strings.Split(p, "__")
			title := strings.Title(strings.Replace(strings.Replace(t[1], ".html", "", -1), "-", " ", -1))
			posts[i] = fmt.Sprintf("<li><a href='%s'>%s</a></li>", p, title)
		}
		unsafe := blackfriday.Run([]byte(fmt.Sprintf("<ul>%s</ul>", strings.Join(posts, ""))))
		fileContent := fmt.Sprintf(
			"%s%s%s",
			applyStyle(layout.GenerateHeader()),
			bluemonday.UGCPolicy().SanitizeBytes(unsafe),
			layout.GenerateFooter(),
		)
		util.GenerateFile(filePath, fileContent)
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
