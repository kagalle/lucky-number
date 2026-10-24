package main

// https://www.alexedwards.net/blog/an-introduction-to-packages-imports-and-modules

/*
was $ run main.go
now $ run *.go
*/

import (
	"fmt"
)

func main() {

	fmt.Printf("Your lucky number is %d!\n", randomNumber())
}
