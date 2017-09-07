package main

import ( 
	"fmt"
	"os"
)

func main() {
	fmt.Println(reverseString(os.Args[1], os.Args[2]))
}

func reverseString(a string, b string) bool {
	if len(a) == len(b) {
		for i := 0; i < len(a); i++ {
			if a[i] != b[len(b)-1-i] {
				return false
			}
		}
		return true
	}
	return false
}
