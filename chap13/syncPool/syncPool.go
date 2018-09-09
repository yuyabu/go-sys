package main

import (
	"fmt"
	"sync"
)

func main() {
	var count int
	pool := sync.Pool{
		New: func() interface{} {
			count++
			return fmt.Sprintf("created:%d", count)
		},
	}
	pool.Put("manualy added: 1")
	pool.Put("manualy added: 2")
	fmt.Println(pool.Get()) //manualy added: 1
	fmt.Println(pool.Get()) //manualy added: 2
	fmt.Println(pool.Get()) //created:1

}
