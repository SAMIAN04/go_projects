package main

import (
	"fmt"
	"math/rand"
// 	"time"
)


func main()  {
	// var fruitList =[]string{"Apple", "Banana", "Cherry"}
	// fruitList = append(fruitList, "Date", "Elderberry")
	// fmt.Printf("%v\n", fruitList)
	// fruitList = fruitList[0:2]
	// fmt.Printf("%v\n", fruitList)

// 	highScores := make([]int, 4)
// 	highScores[0] = 234
// 	highScores[1] = 567
// 	highScores[2] = 890
// 	highScores[3] = 123
// 	//highScores = append(highScores, 456, 789)
// 	if sort.IntsAreSorted(highScores) {
// 		fmt.Println("already sorted")
// 	} else {
		
// 		sort.Ints(highScores)
// 	}
// 	fmt.Printf("high scores: %v\n", highScores)


// languages := make(map[string]string)
// languages["en"] = "English"
// languages["es"] = "Spanish"
// languages["fr"] = "French"
// fmt.Println(languages["es"])
// delete(languages, "es")
// fmt.Println(languages)

// type user struct {
// 	name string
// 	email string
// 	status bool
// 	age int
// }
// samian := user{"Samian", "samiansikdar04@gmail.com", true, 20}
// fmt.Println(samian)
// fmt.Printf("%v\n", samian)


// rand.Seed(time.Now().UnixNano())
diceRoll := rand.Intn(6) + 1
fmt.Println("you rolled a", diceRoll)

}