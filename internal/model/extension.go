package model

import "strings"

type Extension struct {
	Name string
}

func (ext *Extension) RemoveDot() {
	ext.Name = strings.TrimPrefix(ext.Name, ".")
}