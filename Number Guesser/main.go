package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
	"unicode"
)

const RANGE_TO_BE_GUESSED_IN = 100
const TIME_FOR_EACH_GUESS = 10

func main() {
	fmt.Println("---------------------------------")
	fmt.Println("      GUESS THE NUMBER           ")
	fmt.Println("RULES: Time based Scoring, More  ")
	fmt.Println("time it takens, less score for   ")
	fmt.Println("each guess")
	fmt.Println("---------------------------------")

	//generate the number to guess
	numberToGuess := rand.Intn(RANGE_TO_BE_GUESSED_IN)
	fmt.Printf("Number generated. Guess the number in range of 1 to %d \n\n", RANGE_TO_BE_GUESSED_IN)

	fmt.Printf("Start Guessing...\n\n")

	attemptCount := 0
	currentGuess := RANGE_TO_BE_GUESSED_IN + 1
	breakFlag := false
	inputChan := make(chan int, 1)
	errChan := make(chan string, 1)
	replaySlice := []int{}
	score := 0

	for {
		fmt.Printf("Current Attempt: %d \n", attemptCount+1)
		fmt.Printf("Your Guess: ")
		var currentGuessT time.Duration

		go func() {
			startT := time.Now()
			reader := bufio.NewReader(os.Stdin)
			userNumberStr, err := reader.ReadString('\n')
			if err != nil {
				errChan <- err.Error()
				return
			}
			userNumberCleaned := cleanString(userNumberStr)
			if !unicode.IsNumber([]rune(userNumberCleaned)[0]) {
				fmt.Printf("\n%s is not a number. Try again \n\n", userNumberCleaned)
				errChan <- "NOT A NUMBER"
				return
			}
			userNumber, err := strconv.Atoi(userNumberCleaned)
			if err != nil {
				errChan <- err.Error()
				return
			}
			currentGuessT = time.Since(startT)
			inputChan <- userNumber
		}()

		timer := time.NewTimer(TIME_FOR_EACH_GUESS * time.Second)
		defer timer.Stop()

		select {
		case userNumber := <-inputChan:
			currentGuess = userNumber
		case <-timer.C:
			breakFlag = true
			fmt.Println("\n\nTimeout! No input received.")
		case err := <-errChan:
			if err == "NOT A NUMBER" {
				continue
			} else {
				panic(err)
			}
		}

		if breakFlag {
			break
		}

		attemptCount++
		replaySlice = append(replaySlice, currentGuess)
		score += TIME_FOR_EACH_GUESS - int(currentGuessT.Seconds())
		if currentGuess == numberToGuess {
			break
		} else if currentGuess > numberToGuess {
			fmt.Printf("Current Guess is greater than number. \n\n")
		} else {
			fmt.Printf("Current Guess is lesser than number. \n\n")
		}
	}

	if currentGuess == numberToGuess {
		fmt.Printf("\nCorrect Guess. %d attempts taken.", attemptCount)
	} else {
		fmt.Println("\nGood game. Better luck next time.")
	}

	if len(replaySlice) != 0 {
		fmt.Println("\n\nREPLAY:")
		for i, guess := range replaySlice {
			fmt.Printf("Guess #%d: %d\n", i+1, guess)
		}
	}

	fmt.Printf("\nTOTAL SCORE: %d\n\n", score)

	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()
}

func cleanString(numStr string) string {
	cleanedString := strings.ReplaceAll(numStr, "\r\n", "")
	return cleanedString
}
