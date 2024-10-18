package utils

import (
	"github.com/rivo/uniseg"
)

func GetCharacterLength(character string) int {
	return uniseg.GraphemeClusterCount(character)
}
