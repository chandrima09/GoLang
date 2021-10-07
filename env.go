package main

import (
	"os"
	"fmt"
	"strings"
)

func main() {
	for _, element := range os.Environ() {
        variable := strings.Split(element, "=")
        fmt.Println(variable[0],"=",variable[1])		
    }
}
