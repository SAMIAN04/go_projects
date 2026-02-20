package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type ProductInfo struct {
	Name     string
	price    float64
	quantity int
}


type ReportProduct struct {
	reportType string
	quantity   int
}


func main() {
	var initialInventoryStr string
	var operationsStr string
	var parametersStr string
	var initialInventory []ProductInfo
	var AddInventtory []ProductInfo
	var checkinventory []ProductInfo
	var updateInventory []ProductInfo
	var reportInventory []ReportProduct
	var checkop string
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	initialInventoryStr = scanner.Text()
	scanner.Scan()
	operationsStr = scanner.Text()
	scanner.Scan()
	parametersStr = scanner.Text()

	initialInventoryParts := strings.Split(initialInventoryStr, ",")
	operationsParts := strings.Split(operationsStr, ",")
	parametersParts := strings.Split(parametersStr, "|")

	//pharsing and making initialInvantory slice
	for i := 0; i < len(initialInventoryParts); i++ {

		productDetails := strings.Split(initialInventoryParts[i], ":")
		name := productDetails[0]
		price, _ := strconv.ParseFloat(productDetails[1], 64)
		quantity, _ := strconv.Atoi(productDetails[2])
		Products := ProductInfo{Name: name, price: float64(price), quantity: quantity}
		initialInventory = append(initialInventory, Products)

	}
	//pharsing and making opatration perameter slice
	for i := 0; i < len(parametersParts); i++ {
		productDetails := strings.Split(initialInventoryParts[i], ":")
		if len(productDetails[i]) == 3 {
			name := productDetails[0]
			price, _ := strconv.ParseFloat(productDetails[1], 64)
			quantity, _ := strconv.Atoi(productDetails[2])
			AddProductItem := ProductInfo{Name: name, price: float64(price), quantity: quantity}
			AddInventtory = append(AddInventtory, AddProductItem)
		} else if len(productDetails[i]) > 1 {

			checkinventory = append(checkinventory, ProductInfo{Name: productDetails[i]})
			fmt.Println(checkinventory)
		} else if len(productDetails[i]) == 2 && checkop == "update" {
			name := productDetails[0]
			quantity, _ := strconv.Atoi(productDetails[1])
			updateInventory = append(updateInventory, ProductInfo{Name: name, quantity: quantity})
		} else if len(productDetails[i]) == 2 && checkop != "update" {
			report := productDetails[0]
			quantity, _ := strconv.Atoi(productDetails[2])
			reportInventory = append(reportInventory, ReportProduct{reportType: report, quantity: quantity})
		}
	}
	displayheading(initialInventory)

	// doing operation
	for i := 0; i < len(operationsParts); i++ {
		switch operationsParts[i] {
		case "check":
			{
				check(checkinventory[0].Name, initialInventory)
			}
		case "add":
			{
				// Example: call Add function
				// Add()
			}
		case "report":
			{
				// Example: call report function
				// report()
			}
		case "update":
			{
				checkop = "update"
				// Example: call update function
				// update()
			}
		}
	}
}

// all functions
func displayheading(initialInventory []ProductInfo) {
	var totalProducts int
	for i := 0; i < len(initialInventory); i++ {
		totalProducts += initialInventory[i].quantity
	}
	fmt.Println("=== INVENTORY MANAGEMENT SYSTEM ===")
	fmt.Printf("System initialized with %d products\n", totalProducts)
}

func Add(NewProduct []ProductInfo , allProducts []ProductInfo) {
     allProducts = append(allProducts, ProductInfo(NewProduct[0]))
	 fmt.Println(allProducts)
}
func check(productName string , allProducts []ProductInfo) {
	fmt.Println("--- STOCK CHECK --")
	fmt.Printf("Checking stock for: %s\n", productName)
for i := 0; i < len(allProducts); i++ {
	if allProducts[i].Name == productName {
		
		fmt.Printf("Stock level: %d units\n", allProducts[i].quantity)
		break
	}else{
		fmt.Println("Product not found in inventory")
	}
}
}
func report(report []ReportProduct, allProduct []ProductInfo) {
	show := report[0].reportType
	q := report[0].quantity
      fmt.Println(show,q)
}
func update(updateProduct []ProductInfo , allProducts []ProductInfo) {
	found := false
for i := 0; i < len(allProducts); i++ {
	if allProducts[i].Name == updateProduct[0].Name {
		allProducts[i].quantity = updateProduct[0].quantity
		found = true 
		break
		
	} else {
		found = false
	}
}
if !found {
	fmt.Println("product not found")
}
}
