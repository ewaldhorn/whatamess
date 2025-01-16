package main

import (
  "fmt"
  "math/rand"
  "time"
)

func main() {
  rand.Seed(time.Now().UnixNano())

  isHeistOn := true
  eludedGuards := rand.Intn(100)

  if eludedGuards >= 50 {
    fmt.Println("Looks like you've managed to make it past the guards. Good job, but remember, this is the first step.")
  } else {
    isHeistOn = false
    fmt.Println("Plan a better disguise next time?")
  }

  openedVault := rand.Intn(100)

  if isHeistOn && openedVault >= 70 {
    fmt.Println("Grab and GO!")
  } else if isHeistOn {
    isHeistOn = false
    fmt.Println("The vault didn't open!")
  }

  leftSafely := rand.Intn(5)

  if isHeistOn {
    switch leftSafely {
      case 0: isHeistOn=false
      fmt.Println("Heist failed!")
      case 1: isHeistOn=false
      fmt.Println("You got caught!")
      case 2: isHeistOn=false
      fmt.Println("There's no money in the vault!")
      case 3: isHeistOn=false
      fmt.Println("The coppers nabbed you!")
      default: fmt.Println("Start the getaway car!")
    }
  }

  if isHeistOn {
    amtStolen := 10000 + rand.Intn(1000000)
    fmt.Printf("You got %d Zim Dollars!!\n", amtStolen)
  }

  fmt.Println(isHeistOn)
}

