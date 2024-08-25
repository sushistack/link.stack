package utils

import (
	"fmt"
	"os"
	"path/filepath"
)

var ProjectRoot string

func InitProjectRoot() {
	_, err := os.Stat("../go.mod")
	if os.IsNotExist(err) {
		fmt.Println(err)
		ProjectRoot = "./"
	}

	ProjectRoot = filepath.Dir("../go.mod")
}
