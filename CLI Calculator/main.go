package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

const (
	ADD      = "+"
	DIFF     = "-"
	MULITPLY = "*"
	DIVIDE   = "/"
)

func performOperation(operation string, numbers ...int) (float64, error) {
	value := numbers[0]
	switch {
	case operation == ADD:
		for i := 1; i < len(numbers); i++ {
			value += numbers[i]
		}
	case operation == DIFF:
		for i := 1; i < len(numbers); i++ {
			value -= numbers[i]
		}
	case operation == MULITPLY:
		for i := 1; i < len(numbers); i++ {
			value *= numbers[i]
		}
	case operation == DIVIDE:
		{
			if len(numbers) != 2 {
				return 0, errors.New("Cannot perform divison for these many values")
			}
			if numbers[1] == 0 {
				panic("Error: Divde by Zero")
			}
			value = numbers[0] / numbers[1]
		}
	}
	return float64(value), nil
}

func convertStrToInt(numbersStr []string) ([]int, error) {
	numbersInt := []int{}
	for _, num := range numbersStr {
		numInt, err := strconv.Atoi(num)
		if err != nil {
			return numbersInt, errors.New("Input is not a number")
		}
		numbersInt = append(numbersInt, numInt)
	}
	return numbersInt, nil
}

func main() {
	fmt.Println("---------------------------------------")
	fmt.Println("          CLI-Calculator               ")
	fmt.Println("---------------------------------------")
	args := os.Args
	noOfArgs := len(args)
	operation := args[noOfArgs-1]

	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()

	//convert to int
	numbersInt, err := convertStrToInt(args[1 : noOfArgs-1])
	if err != nil {
		panic(err)
	}

	opResult, err := performOperation(operation, numbersInt...)
	if err != nil {
		panic(err)
	}

	//format result
	for i := 0; i < len(numbersInt); i++ {
		if i == len(numbersInt)-1 {
			fmt.Printf(" %d = %.1f", numbersInt[i], opResult)
			break
		}
		fmt.Printf(" %d %s ", numbersInt[i], operation)
	}
}
