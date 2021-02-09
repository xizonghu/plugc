package main

import (
	"fmt"
	"regexp"
)

func main() {
	//reg := regexp.Compile("[\(]{0,1}[a-z]{0,}[\)]{0,1}")
	reg, _ := regexp.Compile("[(]{0,1}[a-z]{0,}[)]{0,1}")
	fmt.Printf("%s\n", reg.FindString("(abc))"))
}
