package main

import "fmt"

/*
Description:
Print the odd numbers from 1 to 99.

Input:
None

Approach:
Obviously, we only want odd numbers. To test if they are even,
we mod 2. If that fails to return 0, its an odd number.
*/
func main() {
     for i := 1; i < 100; i++ {
     	 if i % 2 != 0 {
	    fmt.Println(i)
	 }
     }
}