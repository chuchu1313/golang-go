package main

import (
	"fmt"
	"os"
	"bytes"
)

func main() {
	for i := 1; i < len(os.Args); i++ {
		fmt.Printf("  %s\n", comma(os.Args[i]))
	}
}

func comma(s string) string {
	var buf bytes.Buffer
	n := len(s)

	for i := 0; i < n; i++ {
		if i % 3 == 0 && i != 0 && i != n-1{
			buf.WriteString(",")
		} else {
			fmt.Fprintf(&buf, "%s", string(s[i]))
		}
	}
	return buf.String() 
}
