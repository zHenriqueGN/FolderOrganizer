package model

import "path/filepath"

type File struct {
	Name string
	Extension Extension
}

func (f File) GetExtensionName() string{
	var extension Extension
	extension.Name = filepath.Ext(f.Name)
	extension.RemoveDot()
	return extension.Name
}