package main

import (
	"fmt"
	"html/template"
	"log"
	"os"
	"path"
	"path/filepath"
	"roughpage/model"
	"roughpage/utils"
	"sort"
	"time"

	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/html"
	"github.com/gomarkdown/markdown/parser"
)

func Minus(a, b int) int {
	return a - b
}

func Add(a, b int) int {
	return a + b
}

func FormatCreateMonth(creattime time.Time) string {
	return creattime.Format("January 2006")
}

var partialHtml string

func RenderNew() {
	themePath := filepath.Join(model.RootPath, GlobalConfig.Site.Theme)
	partialstpl, err := template.ParseGlob(filepath.Join(themePath, "partials", "_*.html"))
	if err != nil {
		log.Fatal("ParseGlob() error when reading and parse _*.html")
	}

	Pages := len(Articles) / GlobalConfig.Site.Limit
	PageList := make([]string, Pages+1)
	restList := len(Articles) % GlobalConfig.Site.Limit

	indexByte, err := os.ReadFile("./themes/index.html")
	if err != nil {
		fmt.Println(err)
	}
	indextpl, err := template.New("index").Funcs(template.FuncMap{
		"Add":   Add,
		"Minus": Minus,
	}).Parse(string(indexByte))

	if err != nil {
		log.Fatal("ParseFiles() error when reading and parse index.html\n", err)
	}

	completeIndexTpl, err := partialstpl.AddParseTree("restitem", indextpl.Tree)
	if err != nil {
		log.Fatal("ParseFiles() error when reading and parse restitem.html\n", err)
	}

	completeIndexTpl.Funcs(template.FuncMap{
		"Add":   Add,
		"Minus": Minus,
	})
	var pagePath string
	for i := 0; i <= Pages; i++ {
		var Articlestemp []model.Article
		if i == Pages {
			Articlestemp = Articles[i*GlobalConfig.Site.Limit : i*GlobalConfig.Site.Limit+restList]
		} else {
			Articlestemp = Articles[i*GlobalConfig.Site.Limit : (i+1)*GlobalConfig.Site.Limit]
		}
		if i == 0 {
			pagePath = "./public/index.html"
		} else {
			pagePath = fmt.Sprint("./public/page_", i, ".html")
		}
		outFile, err := os.Create(pagePath)
		if err != nil {
			fmt.Println(err)
		}
		err = completeIndexTpl.Execute(outFile, map[string]interface{}{
			"globalconfig": GlobalConfig,
			"articles":     Articlestemp,
			"categories":   Categories,
			"tags":         UniqueTags,
			"pagination":   PageList,
			"pageNumber":   i,
		})
		if err != nil {
			fmt.Println(err)
		}
	}
}

func RenderCategories() {

	themePath := filepath.Join(model.RootPath, GlobalConfig.Site.Theme)
	partialstpl, err := template.ParseGlob(filepath.Join(themePath, "partials", "_*.html"))
	if err != nil {
		log.Fatal("ParseGlob() error when reading and parse _*.html")
	}

	indexByte, err := os.ReadFile("./themes/category.html")
	if err != nil {
		fmt.Println(err)
	}
	indextpl, err := template.New("index").Funcs(template.FuncMap{
		"Add":   Add,
		"Minus": Minus,
	}).Parse(string(indexByte))

	if err != nil {
		log.Fatal("ParseFiles() error when reading and parse category.html\n", err)
	}

	completeIndexTpl, err := partialstpl.AddParseTree("archieveCategory", indextpl.Tree)
	if err != nil {
		log.Fatal("ParseFiles() error when adding archieveCategory \n", err)
	}

	completeIndexTpl.Funcs(template.FuncMap{
		"Add":   Add,
		"Minus": Minus,
	})

	utils.CreateDirectoryIfnotexist(Categories, "./public/category")

	var pagePath string

	for _, archiveCategory := range Categories {
		var archiveCategoryList []model.Article
		for _, Article := range Articles {
			if archiveCategory == Article.ArticleConfig.Category {
				archiveCategoryList = append(archiveCategoryList, Article)
			}
		}

		Pages := len(archiveCategoryList) / GlobalConfig.Site.Limit
		restList := len(archiveCategoryList) % GlobalConfig.Site.Limit
		var PageList []string
		var totalPages int
		if restList == 0 {
			totalPages = Pages
			PageList = make([]string, totalPages)
		} else {
			totalPages = Pages + 1
			PageList = make([]string, totalPages)
		}
		for i := 0; i < totalPages; i++ {
			var Articlestemp []model.Article
			if len(archiveCategoryList) < GlobalConfig.Site.Limit {
				Articlestemp = archiveCategoryList
			} else {
				if i == totalPages-1 {
					Articlestemp = archiveCategoryList[i*GlobalConfig.Site.Limit : i*GlobalConfig.Site.Limit+restList]
				} else {
					Articlestemp = archiveCategoryList[i*GlobalConfig.Site.Limit : (i+1)*GlobalConfig.Site.Limit]

				}
			}
			if i == 0 {
				pagePath = fmt.Sprint("./public/category/", archiveCategory, "/index.html")
			} else {
				pagePath = fmt.Sprint("./public/category/", archiveCategory, "/page_", i, ".html")
			}
			outFile, err := os.Create(pagePath)
			if err != nil {
				log.Fatal("Create() error when create page.html")
			}
			err = completeIndexTpl.Execute(outFile, map[string]interface{}{
				"globalconfig": GlobalConfig,
				"articles":     Articlestemp,
				"categories":   Categories,
				"Category":     archiveCategoryList[0].ArticleConfig.Category,
				"tags":         UniqueTags,
				"pagination":   PageList,
				"pageNumber":   i,
			})
			if err != nil {
				fmt.Println(err)
			}
		}

	}

}

func RenderTags() {

	themePath := filepath.Join(model.RootPath, GlobalConfig.Site.Theme)
	partialstpl, err := template.ParseGlob(filepath.Join(themePath, "partials", "_*.html"))
	if err != nil {
		log.Fatal("ParseGlob() error when reading and parse _*.html")
	}

	indexByte, err := os.ReadFile("./themes/tag.html")
	if err != nil {
		fmt.Println(err)
	}
	indextpl, err := template.New("index").Funcs(template.FuncMap{
		"Add":   Add,
		"Minus": Minus,
	}).Parse(string(indexByte))

	if err != nil {
		log.Fatal("ParseFiles() error when reading and parse tag.html\n", err)
	}

	completeIndexTpl, err := partialstpl.AddParseTree("archiveTag", indextpl.Tree)
	if err != nil {
		log.Fatal("ParseFiles() error when adding archiveTag \n", err)
	}

	completeIndexTpl.Funcs(template.FuncMap{
		"Add":   Add,
		"Minus": Minus,
	})

	utils.CreateDirectoryIfnotexist(UniqueTags, "./public/tag")

	var pagePath string

	for _, archiveTag := range UniqueTags {
		var archiveTagList []model.Article

		for _, Article := range Articles {
			for _, articleTag := range Article.ArticleConfig.Tags {
				if archiveTag == articleTag {
					archiveTagList = append(archiveTagList, Article)
				}
			}
		}

		Pages := len(archiveTagList) / GlobalConfig.Site.Limit
		restList := len(archiveTagList) % GlobalConfig.Site.Limit
		var PageList []string
		var totalPages int
		if restList == 0 {
			totalPages = Pages
			PageList = make([]string, totalPages)
		} else {
			totalPages = Pages + 1
			PageList = make([]string, totalPages)
		}
		for i := 0; i < totalPages; i++ {
			var Articlestemp []model.Article
			if len(archiveTagList) < GlobalConfig.Site.Limit {
				Articlestemp = archiveTagList
			} else {
				if (i == totalPages-1) && (restList != 0) {
					Articlestemp = archiveTagList[i*GlobalConfig.Site.Limit : i*GlobalConfig.Site.Limit+restList]
				} else {
					Articlestemp = archiveTagList[i*GlobalConfig.Site.Limit : (i+1)*GlobalConfig.Site.Limit]

				}
			}
			if i == 0 {
				pagePath = fmt.Sprint("./public/tag/", archiveTag, "/index.html")
			} else {
				pagePath = fmt.Sprint("./public/tag/", archiveTag, "/page_", i, ".html")
			}
			outFile, err := os.Create(pagePath)
			if err != nil {
				log.Fatal("Create() error when create page.html")
			}
			err = completeIndexTpl.Execute(outFile, map[string]interface{}{
				"globalconfig": GlobalConfig,
				"articles":     Articlestemp,
				"categories":   Categories,
				"tag":          archiveTag,
				"tags":         UniqueTags,
				"pagination":   PageList,
				"pageNumber":   i,
			})
			if err != nil {
				fmt.Println(err)
			}
		}

	}

}

func RenderArchive() {
	themePath := filepath.Join(model.RootPath, GlobalConfig.Site.Theme)
	partialstpl, err := template.ParseGlob(filepath.Join(themePath, "partials", "_*.html"))
	if err != nil {
		log.Fatal("ParseGlob() error when reading and parse _*.html")
	}

	archiveByte, err := os.ReadFile("./themes/archive.html")
	if err != nil {
		fmt.Println(err)
	}
	archivetpl, err := template.New("archive").Funcs(template.FuncMap{
		"FormatDate": FormatCreateMonth,
		"add":        Add,
	}).Parse(string(archiveByte))

	if err != nil {
		log.Fatal("ParseFiles() error when reading and parse archive.html\n", err)
	}

	completeIndexTpl, err := partialstpl.AddParseTree("restitem", archivetpl.Tree)
	if err != nil {
		log.Fatal("ParseFiles() error when reading and parse archive.html\n", err)
	}

	completeIndexTpl.Funcs(template.FuncMap{
		"FormatDate": FormatCreateMonth,
		"add":        Add,
	})

	outFile, err := os.Create("./public/archive.html")
	if err != nil {
		fmt.Println(err)
	}

	sortedByCreatetimeArticles := make([]model.Article, len(Articles))
	copy(sortedByCreatetimeArticles, Articles)
	sort.Slice(sortedByCreatetimeArticles, func(i, j int) bool {
		return sortedByCreatetimeArticles[i].CreateTime.After(sortedByCreatetimeArticles[j].CreateTime)
	})

	err = completeIndexTpl.Execute(outFile, map[string]interface{}{
		"globalconfig": GlobalConfig,
		"articles":     sortedByCreatetimeArticles,
		"categories":   Categories,
		"tags":         UniqueTags,
	})
	if err != nil {
		fmt.Println(err)
	}

}

func RenderArticles() {
	themePath := filepath.Join(model.RootPath, GlobalConfig.Site.Theme)
	partialstpl, err := template.ParseGlob(filepath.Join(themePath, "partials", "_*.html"))
	if err != nil {
		log.Fatal("ParseGlob() error when reading and parse _*.html")
	}

	articleByte, err := os.ReadFile("./themes/article.html")
	if err != nil {
		fmt.Println(err)
	}
	articleTpl, err := template.New("article").Funcs(template.FuncMap{
		"FormatDate": FormatCreateMonth,
	}).Parse(string(articleByte))

	if err != nil {
		log.Fatal("ParseFiles() error when reading and parse article.html\n", err)
	}

	completeIndexTpl, err := partialstpl.AddParseTree("article", articleTpl.Tree)
	if err != nil {
		log.Fatal("ParseFiles() error when reading and parse article.html\n", err)
	}

	completeIndexTpl.Funcs(template.FuncMap{
		"FormatDate": FormatCreateMonth,
	})

	if _, err := os.Stat("./public/article"); os.IsNotExist(err) {
		err := os.Mkdir("./public/article", 0755)
		if err != nil {
			fmt.Println(err)
		}
	}

	for _, Article := range Articles {
		articlepath := filepath.Join("public", "article", path.Base(Article.Filename)+".html")
		outFile, err := os.Create(articlepath)
		if err != nil {
			fmt.Println(err)
		}
		Article.ContentHTML = template.HTML(string(ParseMDtoHTML([]byte(Article.ArticleConfig.Markdown))))
		err = completeIndexTpl.Execute(outFile, map[string]interface{}{
			"globalconfig": GlobalConfig,
			"article":      Article,
		})
		if err != nil {
			fmt.Println(err)
		}
	}
}

func ParseMDtoHTML(md []byte) []byte {
	// create markdown parser with extensions
	extensions := parser.CommonExtensions | parser.Footnotes | parser.HeadingIDs
	p := parser.NewWithExtensions(extensions)
	// create HTML renderer with extensions
	htmlFlags := html.CommonFlags | html.TOC | html.HrefTargetBlank | html.FootnoteReturnLinks
	opts := html.RendererOptions{Flags: htmlFlags}
	renderer := html.NewRenderer(opts)

	// result := markdown.Render(doc, renderer)
	result := markdown.ToHTML(md, p, renderer)
	return result
}
