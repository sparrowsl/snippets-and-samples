package main

import (
	"bytes"
	"flag"
	"fmt"
	"os"
	"path/filepath"

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
        <body>`
	footer = `
	</body>
    </html>`
)

func main() {
	filename := flag.String("file", "", "Markdown file to preview")
	flag.Parse()

	if *filename == "" {
		flag.Usage()
		os.Exit(1)
	}

	if err := run(*filename); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run(filename string) error {
	input, err := os.ReadFile(filename)
	if err != nil {
		return err
	}

	htmlData := parsedContent(input)

	outName := fmt.Sprintf("%s.html", filepath.Base(filename))
	fmt.Println(outName)

	return saveHTML(outName, string(htmlData))
}

func parsedContent(content []byte) []byte {
	var buf bytes.Buffer
	goldmark.Convert(content, &buf)
	body := bluemonday.UGCPolicy().SanitizeBytes(buf.Bytes())

	joined := append([]byte(header), body...)
	joined = append(joined, []byte(footer)...)

	return joined
}

func saveHTML(dest string, data string) error {
	return os.WriteFile(dest, []byte(data), 0644)
}
