package main

import (
	"fmt"

	"github.com/openingo/godkits/gox/stringsx"
)

func main() {
	helloworld := stringsx.DefaultIfEmpty("", "hello world")
	fmt.Println(helloworld)
}
