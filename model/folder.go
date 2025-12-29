package model

type Folder struct {
	Name       string
	SubFolders []*Folder
	Files      []*File
}
