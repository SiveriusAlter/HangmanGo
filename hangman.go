package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	exe, err := os.Executable()
	if err != nil {
		panic(err)
	}

	exeDir := filepath.Dir(exe)
	dictPath := filepath.Join(exeDir, "src", "dictionary.txt")

	dictionary := getDictionary(dictPath)

	game(dictionary)

}

func getDictionary(dictPath string) []string {
	file, err := os.Open(dictPath)
	if err != nil {
		fmt.Printf("Ошибка при открытии файла: %v", err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	var dictionary []string
	for scanner.Scan() {
		dictionary = append(dictionary, scanner.Text())
		if scanner.Err() != nil {
			fmt.Printf("При загрузке %s из словаря упала ошибка %v!",
				scanner.Text(),
				scanner.Err())
		}
	}

	return dictionary
}

func game(dictionary []string) {
	fmt.Println("Привет! Хороший денёк для повешания!")
	word := randomWord(dictionary)

	reader := bufio.NewReader(os.Stdin)

	fmt.Printf("Слово %s уже загадано. Введи первую букву:\n", word)

	char, _, err := reader.ReadRune()
	if err != nil {
		fmt.Println("Ошибка ввода!", err)
	}

	hasRune := strings.ContainsRune(word, char)

	if hasRune {
		fmt.Printf("Верно! Такая буква есть в слове %s!\n", word)
	} else {
		fmt.Printf("Такой буквы нет в слове %s!", word)
	}

}

func randomWord(dictionary []string) string {
	len := len(dictionary)
	if len == 0 {
		panic("Не загружен словарь!")
	}
	num := rand.Intn(len)

	word := dictionary[num]
	return word
}
