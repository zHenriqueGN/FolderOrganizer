package main

import (
	"log"
	"os"
	"path"

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
		extensions = append(extensions, file.Extension)
	}
	
	err = controller.GenerateFoldersByExtension(dstFolder, extensions)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		oldFilePath := path.Join(dstFolder, file.Name)
		newFilePath := path.Join(dstFolder, file.Extension.Name, file.Name)
		err = os.Rename(oldFilePath, newFilePath)
		if err != nil {
			log.Fatal(err)
		}
	}
}