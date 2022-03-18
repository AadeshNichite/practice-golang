/* Note: to run command line argument code in Go use below commands
   $ go build command-line-arguments.go
   $ ./command-line-arguments a b c d */

package main

import (
	"fmt"
	"os"
)

func main() {
	ts := os.Args
	ps := os.Args[1:]
	ks := os.Args[2]

	fmt.Println(ks)
	fmt.Println(ps)
	fmt.Println(ts)
}
