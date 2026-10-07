package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/tdewolff/minify/v2"
	"github.com/tdewolff/minify/v2/css"
	"github.com/tdewolff/minify/v2/html"
	"github.com/tdewolff/minify/v2/js"
)

const limit = 3072

func build(folder string) error {
	// get bundle
	rawBytes, err := bundle(folder)
	if err != nil {
		return fmt.Errorf("bundle failed: %w", err)
	}

	// minify using lib
	m := minify.New()
	m.AddFunc("text/css", css.Minify)
	m.AddFunc("text/javascript", js.Minify)
	m.Add("text/html", &html.Minifier{
		KeepDocumentTags: true,
		KeepEndTags:      true,
	})

	htmlOut, err := m.String("text/html", string(rawBytes))
	if err != nil {
		return fmt.Errorf("minify failed: %w", err)
	}

	// setup for uri
	replacer := strings.NewReplacer("%", "%25", "#", "%23", "\n", "%0A")
	uri := "data:text/html," + replacer.Replace(htmlOut)

	// write uri
	if err := os.WriteFile("uri.txt", []byte(uri), 0644); err != nil {
		return fmt.Errorf("write uri failed: %w", err)
	}

	fmt.Printf("uri written to uri.txt\n")

	// info about uri limit
	bytesLen := len([]byte(uri))
	if bytesLen > limit {
		fmt.Printf("%d / %d bytes, %d over\n", bytesLen, limit, bytesLen-limit)
	} else {
		fmt.Printf("%d / %d bytes, %d left\n", bytesLen, limit, limit-bytesLen)
	}

	return nil
}
