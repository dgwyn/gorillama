package main

import (
	"fmt"

	"github.com/defn/shlib/v2"
	"github.com/dgwyn/gorillama"
)

func main() {
	fmt.Printf("1 + 2 = %d\n", gorillama.Sum(1, 2))
	fmt.Println(shlib.Hello())
}
