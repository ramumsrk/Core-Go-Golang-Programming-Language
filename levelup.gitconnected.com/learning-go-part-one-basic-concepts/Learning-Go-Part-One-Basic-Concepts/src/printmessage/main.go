package main

import (
  "fmt"
  "math/rand"
)

func main() {
  sampleNumber := randomInteger(1, 6)
  fmt.Printf("Generated a number: '%d'\n", sampleNumber)
  printMessage(sampleNumber)
}

func randomInteger(minimum int, maximum int) int {
  return minimum + rand.Intn(maximum - minimum)
}

func printMessage(sampleNumber int) {
  if sampleNumber == 1 {
    fmt.Println("Printing this for a 1")
  } else if sampleNumber == 2 {
    fmt.Println("Printing this for a 2")
  } else if sampleNumber == 3 {
    fmt.Println("Printing this for a 3")
  } else if sampleNumber == 4 {
    fmt.Println("Printing this for a 4")
  } else if sampleNumber == 5 {
    fmt.Println("Printing this for a 5")
  } else {
    fmt.Println("Whoops!, didn't expect this!")
  }
}
