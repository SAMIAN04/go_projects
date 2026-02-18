package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)
 type ProductInfo struct{
	Name string
	price float64
	quantity int
 }
func main()  {
	var initialInventoryStr string
	var operationsStr string
	var parametersStr string
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	initialInventoryStr = scanner.Text()
	scanner.Scan()
	operationsStr= scanner.Text()
    scanner.Scan()
	parametersStr = scanner.Text()
 
	initialInventoryParts := strings.Split(initialInventoryStr, ",")
	parametersParts := strings.Split(parametersStr,"|")
fmt.Println(operationsStr)
var initialInventory []ProductInfo

for i := 0; i < len(initialInventoryParts); i++ {
	
	productDetails:= strings.Split(initialInventoryParts[i],":")
	name := productDetails[0]
	price ,_:= strconv.ParseFloat(productDetails[1],64)
	quantity,_ := strconv.Atoi(productDetails[2])
	Products := ProductInfo{Name: name,price: float64(price),quantity: quantity}
	initialInventory = append(initialInventory, Products)

}
for i := 0; i < len(parametersParts); i++ {
	switch len(parametersParts[i]) {
	case :
		
	}
}
 displayheading(initialInventory)
}

//all functions
func displayheading(initialInventory []ProductInfo) {
	var totalProducts int
	for i := 0; i < len(initialInventory); i++ {
		totalProducts += initialInventory[i].quantity
	}
	fmt.Println("=== INVENTORY MANAGEMENT SYSTEM ===")
	fmt.Printf("System initialized with %d products\n", totalProducts)
}

