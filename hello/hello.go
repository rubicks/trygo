/* github.com/rubicks/trygo/hello/hello.go */

package main

import (
	"fmt"
	"runtime"
)

func main() {
	pc, file, line, ok := runtime.Caller(0)
	fmt.Println("pc	 : ", pc)
	fmt.Println("file : ", file)
	fmt.Println("line : ", line)
	fmt.Println("ok	 : ", ok)
	fmt.Println("Hello, Go!")
}
