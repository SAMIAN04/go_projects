package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main()  {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("enter your girlfriends name :")

	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	fmt.Printf("wow i fucked %s last night, sorry bro \n", input)
}
