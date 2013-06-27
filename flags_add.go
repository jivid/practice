// Use a flag to take a string of two numbers and add them

package main

import (
    "fmt"
    "flag"
    "strings"
    "strconv"
)

func add (a,b int) (s int) {
  s = a+b
  return
}

func main() {
    // variable to store numbers in
    var input string

    // Create the flag and parse input
    flag.StringVar(&input, "numbers", "", "Numbers to be added (separate by comma)")
    flag.Parse()

    // If no input, add 2 and 3. If there are numbers, add them
    if (len(input) == 0) {
        fmt.Println("No input provided. Adding 2 and 3")
        fmt.Println(add(2,3))
    } else {
        numbers := strings.Split(input, ",")
        n1,_ := strconv.Atoi(numbers[0])
        n2,_ := strconv.Atoi(numbers[1])
        fmt.Printf("Adding %d and %d\n", n1, n2)
        fmt.Println(add(n1, n2))
    }
}
