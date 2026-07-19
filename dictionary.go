package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
)

func GetDictionary(dictPath string) []string {
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

func RandomWord(dictionary []string) string {
	len := len(dictionary)
	if len == 0 {
		panic("Не загружен словарь!")
	}
	num := rand.Intn(len)

	word := dictionary[num]
	return word
}
