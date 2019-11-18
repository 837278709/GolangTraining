package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now().Unix()
	fmt.Printf("%T", t)
}
