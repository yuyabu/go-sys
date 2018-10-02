package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	fmt.Printf("Temp File Path: %s\n", filepath.Join(os.TempDir(), "temp.txt"))
}
