package main

import "strconv"
import "fmt"
import "log"
import "bufio"
import "os"

/*
Description:
Print out the sum of integers read from a file.

Input:
The first argument to the program will be a path to a filename containing
a positive integer, one per line

Approach:
We use scanner.Text() to convert scanner.Scan()'s token from a string
to an int. To finish up, we add the number to the running total (sum).
*/
func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var sum int
	for scanner.Scan() {
		temp, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		sum += temp
	}
	fmt.Println(sum)
}
