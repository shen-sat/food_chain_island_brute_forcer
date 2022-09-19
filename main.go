package main

import (
	"fmt"
	"food_chain_island_brute_forcer/brute"
)

func main() {
	names := brute.Force()
	for _, name := range names {
		fmt.Println(name)
	}
}