package main

import (
  "fmt"
  "strconv"
  "strings"
)

func countPossiblePasswords(lower int, upper int) []int {
  possiblePasswords := []int{}

  for currentPassword := lower; currentPassword <= upper; currentPassword++ {
    passwordArray := strconv.Itoa(currentPassword)
    prevPass := toInt(string(passwordArray[len(passwordArray)-1]))
		numberPairs := 1
		// set hasPair to false for pt2
    hasPair, isPotentialPassword := true, false

    for i := len(passwordArray) - 2; i >= 0; i-- {
      current := toInt(string(passwordArray[i]))
			if current == prevPass {
				numberPairs++
				isPotentialPassword = true
				// uncomment this section for pt2
				// if i == 0 {
				// 	if numberPairs == 2 && current == prevPass {
				// 		hasPair = true
				// 	}
				// } else {
				// 	next := toInt(string(passwordArray[i-1]))
				// 	if numberPairs == 2 && current != next {
				// 		hasPair = true
				// 	}
				// }
			} else if current > prevPass {
				isPotentialPassword = false
				break
			} else {
				numberPairs = 1
			}
			prevPass = current
    }

    if isPotentialPassword && hasPair {
      possiblePasswords = append(possiblePasswords, currentPassword)
		}
  }

  return possiblePasswords
}

func toInt(s string) int {
  result, err := strconv.Atoi(s)
  check(err)
  return result
}

func check(err error) {
  if err != nil {
    panic(err)
  }
}

func main() {
  input := "387638-919123"
  stringSplit := strings.Split(input, "-")
  lowerRange := toInt(stringSplit[0])
  upperRange := toInt(stringSplit[1])
  passwordArray := countPossiblePasswords(lowerRange, upperRange)

  fmt.Println("Number of possible passwords:", len(passwordArray))
}