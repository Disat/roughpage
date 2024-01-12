package main

import (
	"fmt"
	"os"
	"roughpage/utils"
)

func main() {

	directoryPath := "./public"
	if _, err := os.Stat(directoryPath); os.IsNotExist(err) {
		fmt.Println("Directory does not exist.")
		return
	}

	// Remove the directory and its contents
	err := os.RemoveAll(directoryPath)
	if err != nil {
		fmt.Printf("Failed to remove directory: %v\n", err)
		return
	}

	fmt.Println("Directory removed successfully.")

	// Check if the directory exists
	if _, err := os.Stat(directoryPath); os.IsNotExist(err) {
		// Directory doesn't exist
		fmt.Printf("Directory %s does not exist\n", directoryPath)

		// Create the directory
		err := os.Mkdir(directoryPath, 0755)
		if err != nil {
			fmt.Println("Failed to create directory:", err)
			return
		}

		fmt.Println("Directory created successfully")
	} else {
		// Directory already exists
		fmt.Printf("Directory %s already exists\n", directoryPath)
	}
	ParseGlobalConfig()
	ParseArticles()
	RenderNew()
	RenderCategories()
	RenderTags()
	RenderArchive()
	RenderArticles()
	utils.Copy([]string{"js", "css", "fonts"})

}
