package main

import (
	"fmt"
	"math/rand"
	"time"
)
func getChance(limt int)int{
	chance :=0
	for limt>0{
		limt/=2
		chance++
	}
	return chance-1
}

func main() {
	rand.Seed(time.Now().UnixNano())
  var (
	limt  = 100
    random_number = rand.Intn(limt)
    input         int
    chance        int 
  )

 
  chance=getChance(limt)
  fmt.Println("Chance: ", chance)
  for {
    fmt.Printf("Input number: ")
    fmt.Scan(&input)
	if random_number>input{
		fmt.Println("tepada")
	}else if random_number<input{
		fmt.Println("pastda")
	}

    if input == random_number {
      fmt.Println("Find number 🎉🎉🎉")
      break
    } else {
      chance--
    }

    fmt.Println("Chance: ", chance)

    if chance <= 0 {
      fmt.Println("You lose 😂")
      break
    }
  }
  fmt.Println("random_number:", random_number)
}

// 1. Top or Botttom
// 2. Limit => chance