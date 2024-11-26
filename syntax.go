package main

import (
	"fmt"
	"math/rand"
	"time"
)

func syntax() {
  // for i := 1; i <= 100; i++ {
  //   switch {
  //   case i % 2 == 1:
  //     fmt.Println(i, "-", "奇数")
  //   default:
  //     fmt.Println(i, "-", "偶数")
  //   }
  // }
  t := time.Now().UnixNano()
  rand.Seed(t)
  s := rand.Intn(100)
  fmt.Println(s)

  switch {
  case s == 6 :
    fmt.Println("大吉")
  case s == 4 || s == 5:
    fmt.Println("中吉")
  case s == 2 || s == 3:
    fmt.Println("小吉")
  case s == 1:
    fmt.Println("凶")
  default:
		syntax()
  }
}