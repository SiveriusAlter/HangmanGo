package hangmango

import (
	"bufio"
	"fmt"
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

	dictionary := GetDictionary(dictPath)

	game(dictionary)

}

func game(dictionary []string) {
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
