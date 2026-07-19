package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Game(dictionary []string) {
	fmt.Println("Привет! Хороший денёк для повешания!")
	word := RandomWord(dictionary)

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
