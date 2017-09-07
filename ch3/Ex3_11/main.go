package main

import (
	"fmt"
	"os"
	"bytes"
	"strings"
)

func main() {
	for i := 1; i < len(os.Args); i++ {
		fmt.Printf("  %s\n", comma(os.Args[i]))
	}
}

func comma(s string) string {
	var buf bytes.Buffer
	n := len(s)
	preNum := strings.Index(s, ".") % 3
	count := 0
	for i := 0; i < n; i++ {
		fmt.Println(count, string(s[i]))
		if (string(s[i]) == ".") {
			fmt.Fprintf(&buf, "%s", string(s[i]))
			count = 0
			preNum = 0
			continue
		} else if (i == preNum && preNum != 0) || ((count-preNum) % 3 == 0  && count != n-1 && count != 0) {
			buf.WriteString(",")
		}
		count++
		fmt.Fprintf(&buf, "%s", string(s[i]))
	}
	return buf.String() 
}
