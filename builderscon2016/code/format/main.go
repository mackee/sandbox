package main

import (
	"fmt"
	"go/format"
)

var bs = []byte(`
func main() {
			fmt  . Println("aaaaaa"           ) }
`)

func main() {
	data, _ := format.Source(bs)
	fmt.Println(string(data))
}
