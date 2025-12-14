package main

import (
	"encoding/json"
	"fmt"
) 
func main()  {
	type user struct {
		Name string
		Age  int
	}
	var u1 user
	u1.Name = "samian"
	u1.Age = 21

	check, err := json.Marshal(u1)
	if err != nil {
		fmt.Println("error occured :", err)
	} else {
		fmt.Println("json data is :", string(check))
	}
}
	
