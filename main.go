//Name:  Vincent Gabb
//Date: 3/26/2020
//Descritpion: Random Feedback



package main

import (

  "fmt"
  "math/rand"
  "time"

)

// messages to be displayed if answer is correct.
 func Correct() {
  s1 := rand.NewSource(time.Now().UnixNano())
  r1 := rand.New(s1)
// declare array of 5 positive messages
var a [5]string
a[0] = "GoodJob!"
a[1] = "Correct!"
a[2] = "Right!"
a[3] = "That's Right!"
a[4] = "NiceJob!"
// print random item from array
fmt.Println(a[r1.Intn(5)])

}



// messages to be displayed if answer is incorrect
func Incorrect(){
  s1 := rand.NewSource(time.Now().UnixNano())
  r1 := rand.New(s1)
// declare array of 5 inncorrect messages
  var b [5]string
b[0] = "Not Quite."
b[1] = "Wrong!"
b[2] = "Incorrect."
b[3] = "Nope!"
b[4] = "Try Again."
// print random item from array
fmt.Println(b[r1.Intn(5)])
}

func main() {

  var input int
// ask question
  fmt.Println("What is 2 + 2 = ?")
  fmt.Scanln(&input)
// if correct
  if input == 4 {
    Correct()
// if incorrecrt
  } else {
    Incorrect()
  }
}