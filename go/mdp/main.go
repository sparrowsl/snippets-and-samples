package main

import (
	"bytes"
	"flag"
	"fmt"
	"os"

	"github.com/microcosm-cc/bluemonday"
	"github.com/yuin/goldmark"
)

const (
	header = `<!DOCTYPE html>
    <html>
        <head>
            <meta http-equiv="content-type" content="text/html; charset=utf-8">
            <title>Markdown Preview Tool</title>
        </head>
        <body>
    `
	footer = `
	</body>
    </html>`
)

func main() {
	file := flag.String("file", "", "Markdown file to parse")
	flag.Parse()

	if *file == "" {
		fmt.Println("No file parsed!!!")
		return
	}

	fileContents, err := os.ReadFile(*file)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	p := bluemonday.UGCPolicy()
	sanitized := p.Sanitize(string(fileContents))

	var buf bytes.Buffer
	if err := goldmark.Convert([]byte(sanitized), &buf); err != nil {
		fmt.Println(err.Error())
		return
	}

	joined := header + buf.String() + footer
	os.WriteFile("./index.html", []byte(joined), os.ModePerm)
}
