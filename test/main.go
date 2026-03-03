package main

import (
	"fmt"
	
	"strconv"
	"strings"
)

type Product struct {
	operation string
	name      string
	price     float64
	quantity  int
}

func main() {
	// Read input
	var numProductsStr string
	var productDataStr string
	var operationsStr string

	fmt.Scanln(&numProductsStr)
	fmt.Scanln(&productDataStr)
	fmt.Scanln(&operationsStr)
	var productData []Product
	var operationsData []Product
	inventory := make(map[string]*Product)
	// TODO: Define Product struct here
 keys := []string{}

	// TODO: Create map to store product pointers

	// TODO: Create slice to maintain product order (to ensure consistent output)

	productData = pharseProducData(productDataStr)
	operationsData = pharseOparationData(operationsStr)
   for _ , key := range productData{
	keys = append(keys, key.name)
   }
	for _, product := range productData {
		inventory[product.name] = &Product{price: product.price, quantity: product.quantity}
	}

	fmt.Println("Initial Inventory:")
	for _, key := range keys {
		fmt.Printf("%s: $%.2f (Stock: %d)\n", key, inventory[key].price, inventory[key].quantity)
	}
	update(operationsData, inventory,keys)

}

func pharse(productDataStr string) []Product {
	productData := []Product{}
	productDataParts := strings.Split(productDataStr, ",")
	for _, product := range productDataParts {
		parts := strings.Split(product, ":")
		price, _ := strconv.ParseFloat(parts[1], 64)
		quantity, _ := strconv.Atoi(parts[2])
		productData = append(productData, Product{name: parts[0], price: price, quantity: quantity})
	}
	return productData
}
func pharseProducData(productDataStr string) []Product {
	productData := []Product{}
	productDataParts := strings.Split(productDataStr, ",")
	for _, product := range productDataParts {
		parts := strings.Split(product, ":")
		price, _ := strconv.ParseFloat(parts[1], 64)
		quantity, _ := strconv.Atoi(parts[2])
		productData = append(productData, Product{name: parts[0], price: price, quantity: quantity})
	}
	return productData
}
func pharseOparationData(productDataStr string) []Product {
	operationsData := []Product{}
	operationsDataParts := strings.Split(productDataStr, ",")
	for _, operation := range operationsDataParts {
		parts := strings.Split(operation, ":")

		operationName := parts[0]
		productName := parts[1]
		price := 0.00
		quantity := 0
		switch operationName {
		case "price":
			{
				price, _ = strconv.ParseFloat(parts[2], 64)
				operationsData = append(operationsData, Product{operation: operationName, name: productName, price: price})
			}
		case "quantity":
			{
				quantity, _ = strconv.Atoi(parts[2])
				operationsData = append(operationsData, Product{operation: operationName, name: productName, quantity: quantity})
			}

		}

	}
	return operationsData
}
func update(operationsData []Product, inventory map[string]*Product, keys []string) {
	var total float64
	for _, opreoperation := range operationsData {
		switch opreoperation.operation {
		case "price":
			{
				inventory[opreoperation.name].price = opreoperation.price
				fmt.Printf("Updated %s: price changed to %.2f\n", opreoperation.name, opreoperation.price)
				

			}
		case "quantity":
			{
				inventory[opreoperation.name].quantity = opreoperation.quantity
				fmt.Printf("Updated %s: quantity changed to %d\n", opreoperation.name,opreoperation.quantity)
			}
		default:
			{
				fmt.Println("Enter a vailid operation")
			}

		}
	}
	fmt.Println("Updated Inventory:")
	
	// sort.Slice(keys,func(i, j int) bool {
	// 	return keys[i] > keys[j]
	// })
	for _,key := range keys{
		total += inventory[key].price * float64(inventory[key].quantity)
		fmt.Printf("%s: $%.2f (Stock: %d)\n",key,inventory[key].price,inventory[key].quantity)
	} 
	
	fmt.Printf("Total Inventory Value: $%.2f\n", total)

}
