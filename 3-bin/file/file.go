package file

import (
	"os"
	"path/filepath"
	"strings"
)

func ReadFile(fileName string) ([]byte, error) {
	data, err := os.ReadFile(fileName)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func IsJSONFile(fileName string) bool {
    return strings.EqualFold(filepath.Ext(fileName), ".json")
}
