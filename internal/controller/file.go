package controller

import (
	"os"
	"path"

	"github.com/zHenriqueGN/FolderOrganizer/internal/model"
)

func GetFolderFiles(folder string) ([]model.File,  error) {
	var files []model.File
	dirEntrys, err := os.ReadDir(folder)
	if err != nil {
		return files, err
	}

	for _, entry := range dirEntrys {
		if !entry.IsDir() {
			var file model.File
			file.Name = entry.Name()
			file.Extension.Name = file.GetExtensionName()
			files = append(files, file)
		}
	}
	return files, nil
}

func GenerateFoldersByExtension(dstFolder string, extensions []model.Extension) error {
	auxExtMap := make(map[string]bool, len(extensions))
	for _, extension := range extensions {
		if _, ok := auxExtMap[extension.Name]; !ok {
			folder := path.Join(dstFolder, extension.Name)

			err := os.Mkdir(folder, 0777)
			if err != nil && !os.IsExist(err) {
				return err
			}

			auxExtMap[extension.Name] = true
		}
	}

	return nil
}

