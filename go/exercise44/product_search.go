package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type Product struct {
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	Quantity int     `json:"quantity"`
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	file, err := os.ReadFile("./products.json")
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
		fmt.Print("What is the product name? ")
		scanner.Scan()
		productName := (scanner.Text())

		product, found := searchProduct(productName, products)
		if !found {
			fmt.Println("Sorry, that product was not found in our inventory.")
			continue
		}

		displayProduct(product)
		break
	}
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
