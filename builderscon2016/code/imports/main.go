package main

import (
	"fmt"

	"golang.org/x/tools/imports"
)

var bs = []byte(`
package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("aaaaaa")
}
`)

func main() {
	result, _ := imports.Process("main.go", bs, nil)
	fmt.Println(string(result))
}
