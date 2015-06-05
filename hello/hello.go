/* github.com/rubicks/trygo/hello/hello.go */

package main

import (
	"fmt"
	"github.com/rubicks/trygo/stringutil"
	"runtime"
)

func main() {
	pc, file, line, ok := runtime.Caller(0)
	fmt.Println("pc	 : ", pc)
	fmt.Println("file : ", file)
	fmt.Println("line : ", line)
	fmt.Println("ok	 : ", ok)
	fmt.Println("Hello, Go!")
	fmt.Printf("stringutil.Reverse(\"race car\") == \"%s\"\n",
		stringutil.Reverse("race car"))
}
