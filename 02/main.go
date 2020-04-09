package main

import (
  "fmt"
  "io/ioutil"
  "strings"
  "strconv"
)

func main() {

  data, err := ioutil.ReadFile("./input.txt")

  if err != nil {
    fmt.Println("File reading error", err)
    panic(err)
  }

  codesArray := []int{}

  for _, v:= range strings.Split(string(data), ",") {
    opCode, err := strconv.Atoi(v)
    if err != nil {
      panic(err)
    }

    codesArray = append(codesArray, opCode)
  }

  // pt 1
  pt1Answer := runComputer(codesArray, 12, 2)
  fmt.Println("pt 1 answer:", pt1Answer)

  // pt 2
  for nounIndex := 0; nounIndex< 100; nounIndex++ {
    for verbIndex := 0; verbIndex < 100; verbIndex++ {
      res := runComputer(codesArray, nounIndex, verbIndex)

      if res == 19690720 {
        fmt.Println("pt 2 answer:", 100 * nounIndex + verbIndex)
      }
    }
  }

}

func runComputer(a []int, noun, verb int) int {
  cleanA := make([]int, len(a))
  copy(cleanA, a)

  // add inputs
  cleanA[1] = noun
  cleanA[2] = verb

  for i, l := 0, len(cleanA) - 1; i < l; i += 4 {
    if i < l-3 && cleanA[i+1] < l && cleanA[i+2] < l && cleanA[i+3] < l  {
      switch cleanA[i] {
      case 99:
        break
      case 1:
        cleanA[cleanA[i+3]] = cleanA[cleanA[i+1]] + cleanA[cleanA[i+2]]
      case 2:
        cleanA[cleanA[i+3]] = cleanA[cleanA[i+1]] * cleanA[cleanA[i+2]]
      default:
        panic("Input machine 🅱️ roke")
      }
    } else {
      panic("Input machine 🅱️ roke x2")
    }
  }

  return cleanA[0]
}