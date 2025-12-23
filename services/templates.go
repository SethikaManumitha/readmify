package services

import (
	_ "embed"
)

//go:embed templates/init.md
var InitTemplate []byte

//go:embed templates/structure.md
var StructureTemplate []byte
