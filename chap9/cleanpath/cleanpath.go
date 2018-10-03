package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	//パスをそのままクリーンにする
	fmt.Println(filepath.Clean("./path/filepath/../path.go"))

}
