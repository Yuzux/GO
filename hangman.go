package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
)

func main() {

	file, err := os.Open("words.txt")
	if err != nil {
		fmt.Println("Error in opening file")
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var words []string
	for scanner.Scan() {
		words = append(words, scanner.Text())
	}

	if len(words) == 0 {
		fmt.Println("The word list is empty.")
		os.Exit(1)
	}

	word := words[randInt(0, len(words))]
	lowerWord := strings.ToLower(word)

	guessedWord := make([]string, len(word))
	for i := range guessedWord {
		guessedWord[i] = "_"
	}

	printWord(guessedWord)

	chances := 10
	for chances > 0 {
		fmt.Println("Enter a letter:")
		reader := bufio.NewReader(os.Stdin)
		letter, _ := reader.ReadString('\n')
		letter = strings.TrimSpace(letter)
		letter = strings.ToLower(letter)

		if len(letter) != 1 {
			fmt.Println("Please enter a single letter.")
			continue
		}

		found := false
		for i := 0; i < len(word); i++ {
			if string(lowerWord[i]) == letter {
				guessedWord[i] = string(word[i])
				found = true
			}
		}

		printWord(guessedWord)

		if isWordGuessed(guessedWord) {
			fmt.Println("Congratulations! You have guessed the word.")
			break
		}

		if !found {
			chances--
			fmt.Println("Chances left:", chances)
		}
	}

	if chances == 0 {
		fmt.Println("You have lost the game. The word was:", word)
	}
}

func printWord(guessedWord []string) {
	for _, value := range guessedWord {
		fmt.Print(value, " ")
	}
	fmt.Println()
}

func isWordGuessed(guessedWord []string) bool {
	for _, value := range guessedWord {
		if value == "_" {
			return false
		}
	}
	return true
}

func randInt(min, max int) int {
	return min + rand.Intn(max-min)
}
