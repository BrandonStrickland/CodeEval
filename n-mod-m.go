package main

import "os"
import "fmt"
import "log"
import "bufio"
import "strings"
import "strconv"


/*
Challenge Description:
Given two integers N and M, calculate N Mod M (without 
using any inbuilt modulus operator).

Input:
Your program should accept as its first argument a path to a filename. 
Each line in this file contains two comma separated positive integers

Approach:
Mod give the remainer of a division. We can find this by finding many
time n will divide m evenly. If we subtract out that whole number from
m, it will give the remainder.
*/
func main() {
     file, err := os.Open(os.Args[1])
     if err != nil {
     	  log.Fatal(err)
     }
     defer file.Close()

     scanner := bufio.NewScanner(file)
     for scanner.Scan() {
     	  line :=  scanner.Text()
	  split := strings.Split(line, ",")
	  
	  mString := split[0]
	  nString := split[1]
	  
	  m, err := strconv.Atoi(mString)
	  if err != nil {
	       log.Fatal(err)
	  }
	  n, err := strconv.Atoi(nString)
	  if err != nil {
	       log.Fatal(err)
	  }

	  multiple := m / n
	  lesser := n * multiple
	  fmt.Println( m - lesser ) 
     } 
}