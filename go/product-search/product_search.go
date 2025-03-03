package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const filename = "./products.json"

type Product struct {
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	Quantity int     `json:"quantity"`
}

func main() {
	file, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
		return
	}

	jsonData := make(map[string][]Product)
	if err := json.Unmarshal(file, &jsonData); err != nil {
		fmt.Println(err)
		return
	}

	products, ok := jsonData["products"]
	if !ok {
		fmt.Println("no products in json...")
		return
	}

	for {
		productName := scanInput("What is the product name? ")

		product, found := searchProduct(productName, products)
		if !found {
			fmt.Println("Sorry, that product was not found in our inventory.")
			response := scanInput("Would you like to add product? (y/n) ")

			if response == "y" {
				product := newProduct()
				product.Name = productName
				products = append(products, product)

				go saveProducts(products)
				fmt.Printf("%s has been added.\n\n", productName)
			} else {
				fmt.Printf("invalid choice!!\n\n")
			}
			continue
		}

		displayProduct(product)
		break
	}
}

func newProduct() Product {
	value := scanInput("Enter product price: ")
	price, _ := strconv.ParseFloat(value, 64)

	value = scanInput("Enter product quantity: ")
	quantity, _ := strconv.Atoi(value)

	return Product{
		Price:    price,
		Quantity: quantity,
	}
}

func saveProducts(products []Product) {
	value := map[string]any{
		"products": products,
	}

	marshalled, _ := json.Marshal(value)
	os.WriteFile(filename, marshalled, os.ModePerm)
}

func scanInput(prompt string) string {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print(prompt)
	scanner.Scan()

	return strings.TrimSpace(scanner.Text())
}

func searchProduct(name string, products []Product) (Product, bool) {
	for _, product := range products {
		if strings.ToLower(product.Name) == strings.ToLower(name) {
			return product, true
		}
	}

	return Product{}, false
}

func displayProduct(product Product) {
	fmt.Printf("Name: %s\n", product.Name)
	fmt.Printf("Price: $%.2f\n", product.Price)
	fmt.Printf("Quantity on hand: %d\n", product.Quantity)
}
