package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Product struct {
	name     string
	price    float64
	quantity int
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

	// TODO: Create map to store product pointers

	// TODO: Create slice to maintain product order (to ensure consistent output)
	
productData = pharse(productDataStr)
operationsData= pharse(operationsStr)
fmt.Println(operationsData)
	// TODO: Parse product data and populate the map
	// Remember to:
	// - Split productDataStr by commas to get individual entries
	// - For each entry, split by colons to get name, price, quantity
	// - Convert price and quantity strings to appropriate types
	// - Store pointer to Product struct in map
	// - Add product name to order slice
	for _, product := range productData {
		inventory[product.name] = &Product{price: product.price,quantity: product.quantity}
	}
	fmt.Println(inventory["laptop"])
	fmt.Println("Initial Inventory:")
	for i, product := range inventory {
		fmt.Printf("%s: $%.2f (Stock: %d)\n",i ,product.price,product.quantity)
	}
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