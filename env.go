//1. env(1)
//   Create a program which lists environment variables as '=' delimited
//   key/value pairs:
//```
//PATH=/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin
//HOSTNAME=e6a0557d8a1b
//```

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
