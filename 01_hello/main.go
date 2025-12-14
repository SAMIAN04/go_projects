package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main()  {
	content := " fck you bitch"
	file ,err := os.Create("./yo.txt")
	if err != nil {
		panic(err)
	}
	io.WriteString(file,content)
	defer fmt.Println("file created successfully")
    defer file.Close()
	defer fmt.Println("file readeed")
	data ,err := ioutil.ReadFile("./yo.txt")
	if err != nil {
        panic(err)
	}
	fmt.Println(string(data))
}