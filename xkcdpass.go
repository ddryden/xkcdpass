package main

import (
	"bufio"
	"crypto/rand"
	"flag"
	"fmt"
	"log"
	"math/big"
	"os"
	"strings"
)

func randi(max int) int {
	randomNumber, err := rand.Int(rand.Reader, big.NewInt(int64(max)))
	if err != nil {
		log.Fatal(err)
	}
	return int(randomNumber.Int64())
}

func readWords(path string, min int, max int) []string {
	words := []string{}
	dict, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer dict.Close()
	scanner := bufio.NewScanner(dict)
	for scanner.Scan() {
		word := scanner.Text()
		if len(word) > min && len(word) < max {
			words = append(words, word)
		}
	}
	return words
}

// https://xkcd.com/936/
// Generates password from local dictionary
func main() {
	dictionaryPath := flag.String("dict", "/usr/share/dict/words", "Path to a custom dictionary")
	wordSeparation := flag.String("ws", " ", "String to separate words")
	min := flag.Int("min", 5, "minimum word length")
	max := flag.Int("max", 9, "maximum word length")
	flag.Parse()

	words := readWords(*dictionaryPath, *min, *max)
	if len(words) < 1000 {
		log.Fatal("Dictionary too small.")
	}

	pass := []string{}
	for n := 0; n < 4; n++ {
		pass = append(pass, words[randi(len(words))])
	}
	fmt.Println(strings.Join(pass, *wordSeparation))
}
