package main

import (
	"fmt"
	"os"
)

func main() {
	os.Link("oldfile.txt", "newfile.txt")
	os.Symlink("oldfile.txt", "newfile-symlink.txt")

	link, err := os.Readlink("newfile-symlink.txt")
	if err != nil {
		panic(err)
	}
	fmt.Print("%v\n", link)
}
