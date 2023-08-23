package controller

import (
	"os"
	"path"

	"github.com/zHenriqueGN/FolderOrganizer/internal/model"
)

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