package main

import (
        "fmt"
        "os"
)

func main() {
        for index, arg := range os.Args[1:] {
                newIndex := index + 1
                fmt.Println("index = ", newIndex, "element = ", arg)
        }
}
