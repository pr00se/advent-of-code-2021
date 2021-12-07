package data

import (
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
)

// ParseCommaSeparatedInts converts a string of comma separated integers into a []int
func ParseCommaSeparatedInts(input string) ([]int, error) {
	var (
		vals []int
	)

	nums := strings.Split(strings.TrimSpace(input), ",")

	for _, n := range nums {
		i, err := strconv.Atoi(strings.TrimSpace(n))
		if err != nil {
			return nil, err
		}
		vals = append(vals, i)
	}

	return vals, nil
}

// ReadInput reads the contents of input.txt located in the same directory as the calling function
func ReadInput() (string, error) {
	_, path, _, _ := runtime.Caller(1)
	path = filepath.Join(filepath.Dir(path), "input.txt")

	content, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}

	return string(content), nil
}
