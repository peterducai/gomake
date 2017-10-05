package main

import (
	"fmt"
)

func main() {
	major := 0
	minor := 0
	build := 0

	var version = fmt.Printf("%d.%d.%d", major, minor, build)

	fmt.Printf("gomake version %s \n", version)
}
