package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"

	"github.com/quantth/my-app/utils"
	"github.com/russross/blackfriday"
)

var tagMap = make(map[string]*Tag)

type Tag struct {
	Name  string
	Count int
	Path  string
	Posts []Post
}

type Post struct {
	Path  string
	Date  string
	Title string
	Tags  []string
}

func handleTag(tags []string, path, date, title string) {
	for _, tag := range tags {
		if val, ok := tagMap[tag]; ok {
			val.Name = tag
			val.Count++

			val.Posts = append(val.Posts, Post{
				Path:  path,
				Date:  date,
				Title: title,
				Tags:  tags,
			})
		} else {
			tagPath := fmt.Sprintf("tags/%s", tag)
			os.MkdirAll(tagPath, 0755)
			tagMap[tag] = &Tag{
				Path:  tagPath,
				Name:  tag,
				Count: 1,
				Posts: []Post{
					{
						Path:  path,
						Date:  date,
						Title: title,
						Tags:  tags,
					},
				},
			}
		}
	}
}

func getAboutMeta(file os.FileInfo) (string, string) {
	id := file.Name()[:len(file.Name())-3]
	title := strings.Split(string(utils.GetFile("_about/"+file.Name())), "\n")[0][2:]

	return id, title
}

func writeAbout() {
	pages := utils.GetDir("_about")

	for i := 0; i < len(pages); i++ {
		_, _ = getAboutMeta(pages[i])

		var b bytes.Buffer
		b.WriteString("const About = () => {\n")
		b.WriteString("return (\n")
		b.WriteString("<div className=\"about\">\n")
		b.Write(blackfriday.MarkdownCommon(utils.GetFile("_about/" + pages[i].Name())))
		b.WriteString("</div>\n")
		b.WriteString(")\n")
		b.WriteString("}\n")
		b.WriteString("export default About;")

		utils.WriteFile("../src/app/components/about/About.js", b)
	}
}

func getPostMeta(fi os.FileInfo) ([]string, string, string) {
	date := fi.Name()[0:10]
	title := strings.Split(string(utils.GetFile("_posts/"+fi.Name())), "\n")[0]
	title = strings.ReplaceAll(title, "[comment]: <> (", "")
	title = strings.ReplaceAll(title, ")", "")

	tags := strings.Split(string(utils.GetFile("_posts/"+fi.Name())), "\n")[1]
	tags = strings.ReplaceAll(tags, "[Comment]: <> (", "")
	tags = strings.ReplaceAll(tags, ")", "")
	return strings.Split(tags, ","), date, title
}

func writePosts() {
	var b bytes.Buffer
	posts := utils.GetDir("_posts")
	b.WriteString("const Posts = () => {\n")
	b.WriteString("return (\n")
	b.WriteString("<div className=\"posts\">\n")
	b.WriteString("<ul class=\"posts--list\">\n")
	for i := len(posts) - 1; i >= 0; i-- {
		tags, date, title := getPostMeta(posts[i])
		dateFolder := strings.ReplaceAll(date, "-", "/")
		path := "/posts/" + dateFolder + "/" + strings.ReplaceAll(title, " ", "-")
		handleTag(tags, path, date, title)
		b.WriteString("<li>\n" +
			"<a class=\"posts--link\" href=\"" + path + "\">" + title + "</a> \n" +
			"<time datetime=\"2017-01-15T00:00:00+00:00\">" + utils.ConvertDate(date) + "</time>\n" +
			"<p>" + parseTagElement(tags) + "</p>\n" +
			"</li>\n")
	}
	b.WriteString("</ul>\n")
	b.WriteString("</div>\n")
	b.WriteString(")\n")
	b.WriteString("}\n")
	b.WriteString("export default Posts;")
	utils.WriteFile("../src/app/components/posts/Posts.js", b)
}

// func writePost() {
// 	posts := utils.GetDir("_posts")

// 	for i := 0; i < len(posts); i++ {
// 		tags, date, title := getPostMeta(posts[i])
// 		var b bytes.Buffer
// 		b.WriteString(getPostPage(title, convertDate(date), tags))
// 		b.Write(blackfriday.MarkdownCommon(getFile("_posts/" + posts[i].Name())))
// 		b.WriteString(getPostEnd())

// 		dateFolder := strings.ReplaceAll(date, "-", "/")
// 		dir := "Posts/" + dateFolder + "/" + strings.ReplaceAll(title, " ", "-")
// 		os.MkdirAll(dir, 0755)
// 		writeFile(dir+"/index", b)
// 	}
// }

func parseTagElement(tags []string) string {
	var b bytes.Buffer
	for _, tag := range tags {
		b.WriteString("<a class=\"posts--tags\" href= \"" + fmt.Sprintf("/tags/%s", tag) + "\">" + fmt.Sprintf("#%s", tag) + "</a>")
		b.WriteString(" ")
	}
	return b.String()
}

func createDirs() {
	os.MkdirAll("_sections", 0755)
	os.MkdirAll("_posts", 0755)
	os.MkdirAll("_about", 0755)
	os.MkdirAll("../src/app/components/posts", 0755)
	os.MkdirAll("../src/app/components/about", 0755)
}

func main() {
	createDirs()
	writeAbout()
	writePosts()
}
