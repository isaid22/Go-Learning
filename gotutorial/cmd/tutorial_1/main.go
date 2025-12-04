package main

import (
	"errors"
	"fmt"
)

func main() {

	var intNum int = 32
	intNum = intNum + 1
	fmt.Println(intNum)

	var myString string = "Hello, World!"
	fmt.Println(len(myString))

	var myBoolean bool = false
	fmt.Println(myBoolean)

	const myConst string = "const value"
	fmt.Println(myConst)
	printMe(myConst)

	var numerator int = 10
	var denominator int = 0
	var result, remainder, err = intDivision(numerator, denominator)
	if err != nil {
		fmt.Printf("%s\n", err.Error())
	} else {
		fmt.Printf("The result of the integer division is %v with remainder %v\n", result, remainder)
	}

}

func printMe(myConst string) {

	fmt.Println("Hello from printMe!")
	fmt.Println(myConst)
}

func intDivision(numerator int, denominator int) (int, int, error) {

	var err error = nil
	if denominator == 0 {
		err = errors.New("Denominator cannot be zero")
		return 0, 0, err
	}

	var result int = numerator / denominator
	var remainder int = numerator % denominator

	return result, remainder, err
}
