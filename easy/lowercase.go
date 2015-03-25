package main

import "strings"
import "os"
import "fmt"
import "log"
import "bufio"

/*
Description:
Given a string write a program to convert it into lowercase.

Input:
The first argument will be a path to a filename containing sentences,
one per line. You can assume all characters are from the english language.

Approach:
I decided to use what Go has to offer and just use ToLower from strings.
*/
func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(strings.ToLower(scanner.Text()))
	}
}
