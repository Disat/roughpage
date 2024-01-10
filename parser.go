package main

import (
	"encoding/json"
	"log"
	"os"
	"path"
	"path/filepath"
	"roughpage/model"
	"roughpage/utils"
	"sort"
	"strings"

	"gopkg.in/yaml.v2"
)

var GlobalConfig model.GlobalConfig
var Article model.Article
var Articles []model.Article
var Categories []string
var UniqueTags []string

func ParseGlobalConfig() {
	dir, _ := filepath.Split(os.Args[0])
	globalconfigfilepath := filepath.Join(dir, "themes", "config.yml")
	globalconfigfile, err := os.ReadFile(globalconfigfilepath)
	if err != nil {
		log.Fatalln("ReadFile() error when reading globalconfigfile .", err)
	}

	err = yaml.Unmarshal(globalconfigfile, &GlobalConfig)
	if err != nil {
		log.Fatalln("yaml.Unmarshal() error when Unmarshaling globalconfigfile ", err)
	}
}

func ParseArticles() {
	files, err := filepath.Glob("./markdown/*.md")
	if err != nil {
		log.Fatal("Glob() error when reading files")
	}
	for _, file := range files {
		articleByte, err := os.ReadFile(file)
		if err != nil {
			log.Fatal("ReadFile() error when reading article file")
		}
		markdownList := strings.SplitN(string(articleByte), "---", 2)
		err = json.Unmarshal([]byte(markdownList[0]), &Article.ArticleConfig)
		if err != nil {
			log.Fatal("Unmarshel error when parse ArticleConfig\n", err)
		}
		Article.ArticleConfig.Markdown = markdownList[1]
		filename := filepath.Base(file)
		filenamewithoutExt := filename[:len(filename)-len(path.Ext(filename))]
		Article.Filename = filenamewithoutExt
		Article.CreateTime, Article.UpdateTime = utils.GetFileCreationAndModifiyTime(file)
		Article.CreateYear = Article.CreateTime.Year()
		Article.CreateMonth = int(Article.CreateTime.Month())
		Article.CreateTimeString = Article.CreateTime.Format("2006-01-02 15:04:05")
		Article.UpdateTimeString = Article.UpdateTime.Format("2006-01-02 15:04:05")
		UniqueTags = append(UniqueTags, Article.ArticleConfig.Tags...)
		Categories = append(Categories, Article.ArticleConfig.Category)
		Articles = append(Articles, Article)
	}
	utils.RemoveDuplicates(&UniqueTags)
	utils.RemoveDuplicates(&Categories)
	sort.Slice(Articles, func(i, j int) bool {
		return Articles[i].UpdateTime.After(Articles[j].UpdateTime)
	})

}
