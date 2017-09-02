package main

import (
	"fmt"
	"github.com/pj2/diamond-pony/raymarcher"
)

func main() {
	vector := raymarcher.D3{}
	fmt.Printf("The X component is %f\n", vector.X)

	marcher := raymarcher.New()
	fmt.Printf("The value of epsilon is %f\n", marcher.Epsilon)
}
