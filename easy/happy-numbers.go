package main

import (
	"fmt"
	"os"
	"bufio"
	"log"
	"strconv"
)

/*
sumSquareDigits takes in an input and a sum. It takes the ones place
and squares it to add to the sum and calls its self again until there
are no more digits left.
*/
func sumSquareDigits(input int, sum int) int {
	if input <= 0 {
		return sum
	} else {
		ones := input % 10
		squared := ones * ones
		return sumSquareDigits(input / 10, sum + squared)
	}
}

/*
This is the meat of the challenge. We simply code the property of a
happy number. This is the beauty of recursion.
*/
func happyChecker(input int,list map[int] int) bool {
	_, present := list[input]
	if input == 1 {
		return true
	} else if present {
		return false
	} else {
		list[input] = input	
		newInput := sumSquareDigits(input,0)
		return happyChecker(newInput, list)
	}
}

/*
Description:
A happy number is defined by the following process. Starting with 
any positive integer, replace the number by the sum of the squares
of its digits, and repeat the process until the number equals 1 
(where it will stay), or it loops endlessly in a cycle which does 
not include 1. Those numbers for which this process ends in 1 are 
happy numbers, while those that do not end in 1 are unhappy numbers.

Input:
The first argument is the pathname to a file which contains test data, 
one test case per line. Each line contains a positive integer.

Approach:
This is a bit more straight forward. We need to convert our string 
from the file into an int to pass it along to the happychecker.
Since, Go doesn't support nested functions (LAME) so we have to 
pass the accumulator to the functions.
*/
func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		num, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		result := happyChecker(num, make(map[int]int))
		if result {
			fmt.Println(1)
		} else {
			fmt.Println(0)
		}
	}
}
