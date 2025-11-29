package main

import "fmt"

func multiply(a int, b int) int {
	return a * b

}
func greet(name string) string {
	return fmt.Sprintf("hello i am your sweet student %s", name)
}
  

func main() {
       age := 17
       if age >= 18 {
		fmt.Println("you can have a girlfriend")
	   }else{
		fmt.Println("still babyboy go study stupid")
	   }
}
