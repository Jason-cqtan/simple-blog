package utils

import (
	"fmt"
	"html/template"
	"os"
	"path/filepath"
	"strings"
)

// LoadTemplates loads all HTML templates from rootDir, using the relative path
// (e.g. "posts/list.html") as each template's name to avoid name conflicts.
func LoadTemplates(rootDir string) (*template.Template, error) {
	tmpl := template.New("")

	err := filepath.Walk(rootDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() || !strings.HasSuffix(path, ".html") {
			return nil
		}

		content, err := os.ReadFile(path)
		if err != nil {
			return err
		}

		relPath, err := filepath.Rel(rootDir, path)
		if err != nil {
			return err
		}

		// Normalise to forward slashes for Windows compatibility.
		relPath = strings.ReplaceAll(relPath, "\\", "/")

		if _, err = tmpl.New(relPath).Parse(string(content)); err != nil {
			return fmt.Errorf("failed to parse template %s: %w", relPath, err)
		}

		return nil
	})

	if err != nil {
		return nil, fmt.Errorf("failed to load templates: %w", err)
	}

	return tmpl, nil
}
