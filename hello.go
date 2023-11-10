package main

import (
	"fmt"
	"math"

	"github.com/rccsdev/ops/go/repo"
)

func main() {
	fmt.Println("hello")
	fmt.Println(math.Pi)
	fmt.Println(repo.SayHi())
}

func SayHi() string {
	return string("hello from GitHub")
}
