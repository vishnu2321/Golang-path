package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const OUTPUT_FILE_EXISTING_ERROR = "The system cannot find the file specified."

func main() {
	fmt.Println("-----------------------------------------------------------------------------------")
	fmt.Println("      				     	FILE WORD COUNTER										")
	fmt.Println("-----------------------------------------------------------------------------------")
	reader := bufio.NewReader(os.Stdin)
	fmt.Println()

	fmt.Print("Enter the relative file path for input text(.txt) file: ")
	filePath, err := reader.ReadString('\n')
	if err != nil {
		panic(err.Error())
	}
	cleanFP := strings.TrimSpace(filePath)
	file, err := os.Open(cleanFP)
	if err != nil {
		panic(err.Error())
	}
	defer file.Close()

	lineCount := 0
	emptyLineCount := 0
	wordCount := 0
	characterCount := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		lineCount++
		if line == "" {
			emptyLineCount++
			continue
		}
		words := strings.Split(line, " ")
		wordCount += len(words)
		for _, word := range words {
			characterCount += len(strings.Split(word, ""))
		}
	}

	var outputFile *os.File
	if _, err := os.Stat("./output.txt"); err == nil {
		outputFile, err = os.OpenFile("./output.txt", os.O_WRONLY, 0644)
		if err != nil {
			panic(err.Error())
		}
	} else if os.IsNotExist(err) {
		outputFile, err = os.Create("./output.txt")
		if err != nil {
			panic(err.Error())
		}
	}
	defer outputFile.Close()
	writer := bufio.NewWriter(outputFile)
	defer writer.Flush()
	fmt.Fprintf(writer, "Line count: %d\n", lineCount)
	fmt.Fprintf(writer, "Empty line count: %d\n", emptyLineCount)
	fmt.Fprintf(writer, "Words count: %d\n", wordCount)
	fmt.Fprintf(writer, "Characters count: %d\n", characterCount)

	fmt.Println("Output file created successfully.")

	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()
}
