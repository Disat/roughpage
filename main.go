package main

import (
	"roughpage/utils"
)

func main() {
	// TemplateFunc()
	utils.ClearDir()
	ParseGlobalConfig()
	ParseArticles()

	RenderNew()
	RenderCategories()
	RenderTags()
	RenderArchive()
	RenderArticles()
	utils.Copy([]string{"js", "css", "fonts"})
	// time.Sleep(3 * time.Second)
	// Build01()
	// Readingtpl()
	// ParseMD()

}
