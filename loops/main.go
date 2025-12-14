package main

import (
	"bufio"
	// "go/scanner"
	"os"
	"fmt"
)

func main()  {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		
		fmt.Println("You entered: \n", line)
	}
}
