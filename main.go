package main

import (
	"fmt"
	"os"
	"roughpage/utils"
)

func main() {

	directoryPath := "./public"

	os.RemoveAll(directoryPath)

	err := os.Mkdir(directoryPath, 0755)
	if err != nil {
		fmt.Println("Failed to create directory:", err)
	}

	ParseGlobalConfig()
	ParseArticles()
	RenderNew()
	RenderCategories()
	RenderTags()
	RenderArchive()
	RenderArticles()
	utils.Copy([]string{"js", "css", "fonts", "favicon.ico", "tools.html", "cheatsheets.html"})

}
