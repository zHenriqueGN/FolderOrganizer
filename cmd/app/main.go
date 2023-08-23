package main

import (
	"log"
	"path/filepath"

	"github.com/zHenriqueGN/FolderOrganizer/internal/controller"
	"github.com/zHenriqueGN/FolderOrganizer/internal/model"
)

const (
	dstFolder = "/workspaces/FolderOrganizer/test"
)

func main() {
	files, err := controller.GetFolderFiles(dstFolder)
	if err != nil {
		log.Fatal(err)
	}

	var extensions = []model.Extension{}

	for _, file := range files {
		fileExtension := filepath.Ext(file)
		extension := model.Extension{Name: fileExtension}
		extension.RemoveDot()
		extensions = append(extensions, model.Extension{Name: extension.Name})
	}
	
	err = controller.GenerateFoldersByExtension(dstFolder, extensions)
	if err != nil {
		log.Fatal(err)
	}
}