package c4f

import (
	"github.com/fatih/color"
)

func Println(str string) {
	if len(str) == 0 {
		return
	}

	color.Red(str)
}
