package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
)

/*
Description:
You have JSON string which describes a menu. Calculate the SUM of
IDs of all "items" in the case a "label" exists for an item.

Input:
Your program should accept as its first argument a path to a filename.
All IDs are integers between 0 and 100. It can be 10 items maximum for a menu.

Approach:
We can only add the ID if there is a label for it. We can use Regex since it
can easily pick out a label with an ID.
*/
func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	//We are going to assume I know how to write an regex
	reg, _ := regexp.Compile("Label [0-9]+")
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(reg.Split(line, -1))
	}
}
