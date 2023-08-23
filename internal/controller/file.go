package controller

import (
	"os"
	"path"

	"github.com/zHenriqueGN/FolderOrganizer/internal/model"
)

func GetFolderFiles(folder string) ([]string, error) {
	var files []string
	dirEntrys, err := os.ReadDir(folder)
	if err != nil {
		return files, err
	}

	for _, entry := range dirEntrys {
		if !entry.IsDir() {
			files = append(files, entry.Name())
		}
	}
	return files, nil
}

func GenerateFoldersByExtension(dstFolder string, extensions []model.Extension) error {
	for _, extension := range extensions {
		folder := path.Join(dstFolder, extension.Name)

		err := os.Mkdir(folder, 0666)
		if err != nil {
			return err
		}
	}

	return nil
}