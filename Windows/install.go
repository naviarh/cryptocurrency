package main

import (
	"os"
)

func main() {
	// for _, envir := range os.Environ() {
	// 	fmt.Println(envir)
	// }
	// println()
	println(os.Getenv("PATH"))
}
