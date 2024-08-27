package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const indexHtml = `<!DOCTYPE html>
<html lang="en">
  <head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta name="author" content="{{author}}">
    <title>{{sitename}}</title>
  </head>
  <body>
  </body>
</html>`

func main() {
	// currDir, _ := os.Getwd()

	siteName := readInput("Site name: ")
	if siteName == "" {
		fmt.Println("site name must not be empty")
		return
	}

	author := readInput("Author: ")
	addJS := readInput("Do you want a folder for JavaScript?: ")
	addCSS := readInput("Do you want a folder for CSS?: ")

	if err := os.Mkdir(siteName, os.ModePerm); err != nil {
		switch {
		case os.IsExist(err):
			fmt.Printf("%q already exists, skipping...\n", siteName)
		default:
			fmt.Println(err)
		}
	}

	os.Chdir(siteName)

	file, err := os.Create("index.html")
	if err != nil {
		fmt.Println(err)
	}

	nameReplaced := strings.ReplaceAll(indexHtml, "{{sitename}}", siteName)
	authorReplaced := strings.ReplaceAll(nameReplaced, "{{author}}", author)
	file.WriteString(authorReplaced)

	fmt.Printf("Created: ./%s\n", siteName)
	fmt.Printf("Created: ./%s/%s\n", siteName, "index.html")
	if addJS == "y" {
		os.Mkdir(siteName+"/js/", os.ModePerm)
		fmt.Printf("Created: ./%s/js/\n", siteName)
	}

	if addCSS == "y" {
		os.Mkdir(siteName+"/css/", os.ModePerm)
		fmt.Printf("Created: ./%s/css/\n", siteName)
	}
}

func readInput(prompt string) string {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print(prompt)
	scanner.Scan()

	return strings.TrimSpace(scanner.Text())
}
