package main

import (
	"roughpage/utils"
)

func main() {
	utils.CreateDirectoryIfnotexist([]string{"public"}, "./")

	utils.ClearDir()
	ParseGlobalConfig()
	ParseArticles()
	RenderNew()
	RenderCategories()
	RenderTags()
	RenderArchive()
	RenderArticles()

	utils.Copy([]string{"js", "css", "fonts"})

}
