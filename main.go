package main

import (
	"os"
	"path/filepath"
)

func main() {
	exe, err := os.Executable()
	if err != nil {
		panic(err)
	}

	exeDir := filepath.Dir(exe)
	dictPath := filepath.Join(exeDir, "src", "dictionary.txt")

	dictionary := GetDictionary(dictPath)

	Game(dictionary)

}
