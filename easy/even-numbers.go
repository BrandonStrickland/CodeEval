package main

import (
	"fmt"
	"os"
	"log"
	"strconv"
	"bufio"
)

/*
Description:
Write a program which checks input numbers and determines 
whether a number is even or not.

Input:
Your program should accept as its first argument a path 
to a filename.

Approach:
I simply took the input and converted it to an int and
used mod 2 to see if it was even.
*/
func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		number, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		if number % 2 == 0 {
			fmt.Println(1)
		} else {
			fmt.Println(0)
		}
	}
}
