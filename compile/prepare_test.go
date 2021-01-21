package compile

import (
	"fmt"
	"strings"
	"testing"
)

func TestNewWith(t *testing.T) {
	str := "a b  c   d    e     f      g       h        i"
	strs := strings.Split(str, " ")
	for _, s := range strs {
		fmt.Printf("s=%s\n", s)
	}
	for ; strings.Contains(str, "  "); str = strings.ReplaceAll(str, "  ", " ") {}
	fmt.Printf("str=%s\n", str)
	fmt.Printf("done")
}

