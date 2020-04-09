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

  totalModules := strings.Split(string(data), "\n")[:]
  totalModuleFuels, totalFuels := 0, 0

  for _, module := range totalModules {
    mass, err := strconv.ParseFloat(module, 32)
    if err != nil {
      panic(err)
    }
    
    fuelConsumption := fuelConsumption(mass)
    totalModuleFuels += fuelConsumption
    totalFuels += fuelForFuelConsumption(fuelConsumption)
  }
}

func fuelConsumption(n float64) int {
  return int(n/3 - 2)
}

func fuelForFuelConsumption(n int) int {
  fuelConsumption := fuelConsumption(float64(n))

  if (fuelConsumption <= 0) {
    return n
  }

  return n + fuelForFuelConsumption(fuelConsumption)
}