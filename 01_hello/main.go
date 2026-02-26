package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
"sort"
	
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
	initialInventory := []*ProductInfo{}
	var AddInventtory []ProductInfo
	var checkinventory []ProductInfo
	var updateInventory []ProductInfo
	var reportInventory []ReportProduct
	var products []string

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

		// Fix: append a pointer to ProductInfo struct to the slice
		initialInventory = append(initialInventory, &ProductInfo{Name: name, price: price, quantity: quantity})

	}
	//pharsing and making opatration perameter slice
	for i := 0; i < len(operationsParts); i++ {

		switch operationsParts[i] {
		case "add":
			products = strings.Split(parametersParts[i], ":")
			name := products[0]
			price, _ := strconv.ParseFloat(products[1], 64)
			quantity, _ := strconv.Atoi(products[2])
			AddProductItem := ProductInfo{Name: name, price: float64(price), quantity: quantity}
			AddInventtory = append(AddInventtory, AddProductItem)
		case "check":
			products = strings.Split(parametersParts[i], ":")
			checkinventory = append(checkinventory, ProductInfo{Name: parametersParts[i]})

		case "update":
			products = strings.Split(parametersParts[i], ":")
			name := products[0]
			quantity, _ := strconv.Atoi(products[1])
			updateInventory = append(updateInventory, ProductInfo{Name: name, quantity: quantity})
		case "report":
			products = strings.Split(parametersParts[i], ",")
			
			report := products[0]
			quantity, _ := strconv.Atoi(products[1])
			reportInventory = append(reportInventory, ReportProduct{reportType: report, quantity: quantity})
		}
	}
	// Convert []*ProductInfo to []ProductInfo for displayheading
	initialInventoryVals := make([]ProductInfo, len(initialInventory))
	for i, p := range initialInventory {
		initialInventoryVals[i] = *p
	}
	displayheading(initialInventoryVals)

	// doing operation
	for i := 0; i < len(operationsParts); i++ {
		switch operationsParts[i] {
		case "check":
			{
				// we need to add a pointer stuff in here becuse check sint working

				check(checkinventory[0].Name, initialInventoryVals)

				fmt.Println("Operation completed. Continuing to next operation...")
			}
		case "add":
			{
				Add(AddInventtory, &initialInventoryVals)
				fmt.Println("Operation completed. Continuing to next operation...")
			}
		case "report":
			{
				
				report(reportInventory, initialInventoryVals)
				fmt.Println("Operation completed. Continuing to next operation...")
			}
		case "update":
			{
				update(updateInventory[0].Name, updateInventory[0].quantity, &initialInventoryVals)

				fmt.Println("Operation completed. Continuing to next operation...")
			}
		case "exit":
			{
exit(&initialInventoryVals)
			}
		}
	}
}

// all functions
func displayheading(initialInventory []ProductInfo) {
	var totalProducts int
	for i := 0; i < len(initialInventory); i++ {
		totalProducts +=1
	}
	fmt.Println("=== INVENTORY MANAGEMENT SYSTEM ===")
	fmt.Printf("System initialized with %d products\n", totalProducts)
	fmt.Println("Starting interactive session...")
}

func Add(NewProduct []ProductInfo, allProducts *[]ProductInfo) {
	fmt.Println("--- ADD ITEM ---")
	fmt.Printf("Adding new product: %s\n", NewProduct[0].Name)

	// Fix: Dereference allProducts and update the slice
	*allProducts = append(*allProducts, ProductInfo(NewProduct[0]))
	fmt.Println("Product added successfully")
}

func check(productName string, allProducts []ProductInfo) {
	fmt.Println("--- STOCK CHECK ---")
	fmt.Printf("Checking stock for: %s\n", productName)
	found := false
	for i := 0; i < len(allProducts); i++ {
		if allProducts[i].Name == productName {

			fmt.Printf("Stock level: %d units\n", allProducts[i].quantity)
			found = true
			break
		}
	}
	if !found {
		fmt.Println("Product not found in inventory")
	}
}
func report(report []ReportProduct, allProduct []ProductInfo) {
	fmt.Println("--- GENERATE REPORT ---")
	found := false
	for i := 0; i < len(report); i++ {
		fmt.Printf("Generating %s report with threshold %d\n", report[i].reportType, report[i].quantity)
		switch report[i].reportType {
		case "low":
			{
				fmt.Println("=== LOW REPORT ===")
				fmt.Printf("Products with stock below %d:\n", report[i].quantity)
				for j := 0; j < len(allProduct); j++ {
					if report[i].quantity > allProduct[j].quantity {
						fmt.Printf("- %s: %d units (Price: $%v)\n", allProduct[j].Name, allProduct[j].quantity, allProduct[j].price)
						found = true
					}

				}
				if !found {
					fmt.Println("no product found")
				}
				// Complete inventory (minimum threshold: 12):

			}
		case "full":
			{
				fmt.Println("=== FULL REPORT ===")
                sort.Slice(allProduct,func(i, f int) bool {
				return allProduct[i].Name < allProduct[f].Name
			})
			
				fmt.Printf("Complete inventory (minimum threshold: %d):\n", report[i].quantity)
				for j := 0; j < len(allProduct); j++ {
					// we have problem of selecting producta accroding to thashold here
					if report[i].quantity <= allProduct[j].quantity {
						fmt.Printf("- %s: %d units @ $%.2f each [OK]\n", allProduct[j].Name, allProduct[j].quantity, allProduct[j].price)
						found = true
					}

				}
				if !found {
					fmt.Println("no product found")
				}
			}

		}
	}

}
func update(poductName string, stock int, allProducts *[]ProductInfo) {

	fmt.Println("--- UPDATE STOCK ---")
	fmt.Printf("Updating stock for: %s\n", poductName)
	// Loop through items to UPDATE (not all products!)
	for i := 0; i < len(*allProducts); i++ {

		// Find matching product in allProducts

		if (*allProducts)[i].Name == poductName {

			// Actually UPDATE the quantity! 🎯
			(*allProducts)[i].quantity += stock

			if stock >= 0 {
				fmt.Printf("Added %d units. New stock: %d\n",
					stock, (*allProducts)[i].quantity)

			} else if stock < 0 {
				fmt.Printf("Removed %d units. New stock: %d\n",
					-stock, (*allProducts)[i].quantity)

			} else {
				fmt.Println("Update failed: Product not found")
			}

		}

	}
}
func exit(allProducts *[]ProductInfo) {
	fmt.Println("--- SYSTEM EXIT ---")

	totalProducts := 0
	totalItems := 0
	totalPrice := 0.00

	for i := 0; i < len(*allProducts); i++ {
		totalProducts += 1
		totalItems += (*allProducts)[i].quantity
		totalPrice += float64((*allProducts)[i].quantity) * (*allProducts)[i].price
	}

	fmt.Println("Final inventory status:")
	fmt.Printf("Total products: %d\n", totalProducts)
	fmt.Printf("Total items: %d\n", totalItems)
	fmt.Printf("Total value: $%.2f\n", totalPrice)
	fmt.Println("Session completed successfully")
	fmt.Println("Thank you for using the Inventory Management System")
}