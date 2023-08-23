package main

import (
	"log"
	"os"
	"path"

	"github.com/zHenriqueGN/FolderOrganizer/internal/config"
	"github.com/zHenriqueGN/FolderOrganizer/internal/controller"
	"github.com/zHenriqueGN/FolderOrganizer/internal/model"
)

func init() {
	config.LoadEnv()
}

func main() {
	files, err := controller.GetFolderFiles(config.RootFolder)
	if err != nil {
		log.Fatal(err)
	}

	var extensions = []model.Extension{}

	for _, file := range files {
		extensions = append(extensions, file.Extension)
	}
	
	err = controller.GenerateFoldersByExtension(config.RootFolder, extensions)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		oldFilePath := path.Join(config.RootFolder, file.Name)
		newFilePath := path.Join(config.RootFolder, file.Extension.Name, file.Name)
		err = os.Rename(oldFilePath, newFilePath)
		if err != nil {
			log.Fatal(err)
		}
	}
}