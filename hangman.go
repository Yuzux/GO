package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
)

func main() {

	positions, err := loadJosePositions("hangman.txt")
	if err != nil {
		fmt.Println("Error in loading José's positions:", err)
		os.Exit(1)
	}

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

	lettersToReveal := (len(word) / 2) - 1
	if lettersToReveal < 1 {
		lettersToReveal = 1
	}

	revealRandomLetters(word, guessedWord, lettersToReveal)

	chances := 10
	fmt.Printf("Good Luck, you have %d attempts.\n", chances)
	printWord(guessedWord)

	for chances > 0 {
		fmt.Print("\nChoose: ")
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
			displayJosePosition(positions, chances)
			break
		}

		if !found {
			chances--
			fmt.Printf("Not present in the word, %d attempts remaining\n", chances)
			displayJosePosition(positions, chances)
		}
	}

	if chances == 0 {
		fmt.Println("You have lost the game. The word was:", word)
		displayJosePosition(positions, 0)
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

func revealRandomLetters(word string, guessedWord []string, n int) {
	revealedIndices := map[int]bool{}

	for len(revealedIndices) < n {
		randomIndex := rand.Intn(len(word))
		if _, alreadyRevealed := revealedIndices[randomIndex]; !alreadyRevealed {
			guessedWord[randomIndex] = string(word[randomIndex])
			revealedIndices[randomIndex] = true
		}
	}
}

func loadJosePositions(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var positions []string
	var position string

	for scanner.Scan() {
		line := scanner.Text()
		if strings.TrimSpace(line) == ";" {
			positions = append(positions, position)
			position = ""
		} else {
			position += line + "\n"
		}
	}

	if strings.TrimSpace(position) != "" {
		positions = append(positions, position)
	}

	return positions, scanner.Err()
}

func displayJosePosition(positions []string, chances int) {
	index := 10 - chances
	if index >= 10 {
		index = 9
	} else if index < 0 {
		index = 0
	}

	fmt.Println(positions[index])
}
