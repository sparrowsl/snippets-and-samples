package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/microcosm-cc/bluemonday"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/parser"
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
	footer = `</body>
    </html>`
)

func main() {
	filename := flag.String("file", "", "Markdown file to preview")
	flag.Parse()

	if *filename == "" {
		flag.Usage()
		os.Exit(1)
	}

	if err := run(*filename, os.Stdout); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run(filename string, out io.Writer) error {
	input, err := os.ReadFile(filename)
	if err != nil {
		return err
	}

	htmlData := parsedContent(input)

	temp, err := os.CreateTemp("", "mdp*.html")
	if err != nil {
		return err
	}

	if err := temp.Close(); err != nil {
		return err
	}

	outName := temp.Name()
	fmt.Fprintln(out, outName)

	return saveHTML(outName, string(htmlData))
}

func parsedContent(content []byte) []byte {
	var buf bytes.Buffer
	goldmark.Convert(content, &buf, parser.WithContext(parser.NewContext()))
	body := bluemonday.UGCPolicy().SanitizeBytes(buf.Bytes())

	joined := append([]byte(header), body...)
	joined = append(joined, []byte(footer)...)

	return joined
}

func saveHTML(dest string, data string) error {
	return os.WriteFile(dest, []byte(data), 0644)
}
