package main

import (
  "fmt"
  "math"
  "io/ioutil"
  "strconv"
  "strings"
)

// Coordinates - an x and y value corresponding to a given location on the grid
type Coordinates struct {
  x, y int
}

// distance - calculates distance for coordinates structure
func (currentCoords Coordinates) distance() int {
  return int(math.Abs(float64(currentCoords.x)) + math.Abs(float64(currentCoords.y)))
}

// calcCoord - calculates the addition or subtraction of 2 sets of coord values
func (currentCoords Coordinates) calcCoord(destinationCoords Coordinates, operator string) Coordinates {
  result := Coordinates{}
  switch operator {
  case "add":
    result = Coordinates{
      x: currentCoords.x + destinationCoords.x,
      y: currentCoords.y + destinationCoords.y,
    }
  case "subtract":
    result = Coordinates{
      x: currentCoords.x - destinationCoords.x,
      y: currentCoords.y - destinationCoords.y,
    }
  }
  
  return result
}

func toInt(s string) int {
  result, err := strconv.Atoi(s)
  check(err)
  return result
}

func calcMin(x, y int) int {
  if x < y {
      return x
  }
  return y
}

func check(err error) {
  if err != nil {
    panic(err)
  }
}

func parseDirection(direction string) (coord Coordinates, length int) {
  switch direction[0] {
  case 'R':
    coord.x++
  case 'L':
    coord.x--
  case 'U':
    coord.y++
  case 'D':
    coord.y--
  }

  return coord, toInt(direction[1:])
}

func main() {
  data, err := ioutil.ReadFile("./input.txt")

  if err != nil {
    fmt.Println("File reading error", err)
    check(err)
  }

  coordsArray := strings.Split(string(data), "\n")

  firstWire := strings.Split(coordsArray[0], ",")
  secondWire := strings.Split(coordsArray[1], ",")

  grid := make(map[Coordinates]int)

  pos, steps := Coordinates{}, 0
  for _, direction := range firstWire {
    coord, length := parseDirection(direction)
    for i := 0; i < length; i++ {
      pos = pos.calcCoord(coord, "add")
      steps++
      grid[pos] = steps
    }
  }

  pt1, pt2 := math.MaxInt32, math.MaxInt32

  pos, steps = Coordinates{}, 0
  for _, direction := range secondWire {
    coord, length := parseDirection(direction)

    for i := 0; i < length; i++ {
      pos = pos.calcCoord(coord, "add")
      steps++
      if grid[pos] != 0 {
        distance := pos.distance()
        pt1 = calcMin(pt1, distance)
        totalSteps := grid[pos] + steps
        pt2 = calcMin(pt2, totalSteps)
      }
    }
  }

    fmt.Println("Closest Manhattan (pt1):", pt1)
    fmt.Println("Closest Steps (pt2):", pt2)
}