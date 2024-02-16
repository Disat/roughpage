package main

import (
	"encoding/json"
	"fmt"
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
		var article model.Article
		articleByte, err := os.ReadFile(file)
		if err != nil {
			log.Fatal("ReadFile() error when reading article file")
		}
		markdownList := strings.SplitN(string(articleByte), "---", 2)
		err = json.Unmarshal([]byte(markdownList[0]), &article.ArticleConfig)
		if err != nil {
			fmt.Println(file+"occured an error")
			log.Fatal("Unmarshel error when parse ArticleConfig\n", err)
		}
		article.ArticleConfig.Markdown = markdownList[1]
		filename := filepath.Base(file)
		filenamewithoutExt := filename[:len(filename)-len(path.Ext(filename))]
		article.Filename = filenamewithoutExt
		article.CreateTime, article.UpdateTime = utils.GetFileCreationAndModifiyTime(file)
		article.CreateYear = article.CreateTime.Year()
		article.CreateMonth = int(article.CreateTime.Month())
		article.CreateTimeString = article.CreateTime.Format("2006-01-02 15:04:05")
		article.UpdateTimeString = article.UpdateTime.Format("2006-01-02 15:04:05")
		UniqueTags = append(UniqueTags, article.ArticleConfig.Tags...)
		Categories = append(Categories, article.ArticleConfig.Category)
		Articles = append(Articles, article)
	}
	utils.RemoveDuplicates(&UniqueTags)
	utils.RemoveDuplicates(&Categories)
	sort.Slice(Articles, func(i, j int) bool {
		return Articles[i].UpdateTime.After(Articles[j].UpdateTime)
	})

}
