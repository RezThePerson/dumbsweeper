package dev

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
)

func bundle(dir string) ([]byte, error) {
	var buf bytes.Buffer

	// css
	cssPath := filepath.Join(dir, "style.css")
	if cssData, err := os.ReadFile(cssPath); err == nil {
		buf.WriteString("<style>\n")
		buf.Write(cssData)
		buf.WriteString("\n</style>\n")
	} else if !os.IsNotExist(err) {
		return nil, fmt.Errorf("read %s: %w", cssPath, err)
	}

	// html
	htmlPath := filepath.Join(dir, "index.html")
	if htmlData, err := os.ReadFile(htmlPath); err == nil {
		buf.Write(htmlData)
		buf.WriteString("\n")
	} else if !os.IsNotExist(err) {
		return nil, fmt.Errorf("read %s: %w", htmlPath, err)
	}

	// js
	jsPath := filepath.Join(dir, "script.js")
	if jsData, err := os.ReadFile(jsPath); err == nil {
		buf.WriteString("<script>\n")
		buf.Write(jsData)
		buf.WriteString("\n</script>\n")
	} else if !os.IsNotExist(err) {
		return nil, fmt.Errorf("read %s: %w", jsPath, err)
	}

	return buf.Bytes(), nil
}
