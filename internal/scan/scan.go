package scan

import (
	"fmt"
	"os"
	"path/filepath"
)

type Scanner struct {
	extensions []string
}

func (s *Scanner) Walk(root string) error {
	archives := map[string]string{}
	err := filepath.Walk(root, func(path string, entry os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if entry.IsDir() {
			for _, ext := range s.extensions {
				arc := path + ext
				archives[arc] = path
			}
			return nil
		}

		if dir, ok := archives[path]; ok {
			name := entry.Name()
			fmt.Printf("Unpacked archive %s (%s)\n", name, dir)
		}

		return nil
	})
	return err
}

func New(extensions []string) *Scanner {
	return &Scanner{extensions}
}
