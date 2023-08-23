package controller

import (
	"os"
	"path"
)

func GenerateFoldersByExtension(dstFolder string, extensions []string) error {
	for _, extension := range extensions {
		folder := path.Join(dstFolder, extension)

		err := os.Mkdir(folder, 0666)
		if err != nil {
			return err
		}
	}

	return nil
}