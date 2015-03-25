package main

import "fmt"
import "os"
import "bufio"
import "log"
import "strconv"

/*
Description:
Given a positive integer, find the sum of its constituent digits.

Input:
The first argument will be a path to a filename containing positive integers, 
one per line

Approach:
I take each line and convert it to an Int. Then I pass it to a function that
is tail recursive(language doesn't support that but it was an exercise) that will
take each digit and add it to a sum and pass the remaining digits back to itself.
This appraoch starts from the ones place and works to the left using mod and divide.
*/
func adder(input int, sum int) int {
     if input <= 0 {
     	  return sum
     } else {
     	  return adder(input / 10, sum + input % 10)
     }
}

func main() {
     file, err := os.Open(os.Args[1])
     if err != nil {
     	log.Fatal(err)
     }
     defer file.Close()
     
     scanner := bufio.NewScanner(file)
     for scanner.Scan() {
     	 number,err := strconv.Atoi(scanner.Text())
	 if err != nil {
	      log.Fatal(err)
	 } 
	 fmt.Println(adder(number,0))
     }
}