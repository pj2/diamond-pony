package main

import (
	"fmt"
	"github.com/pj2/diamond-pony/raymarcher"
	"github.com/pj2/diamond-pony/renderer"
	"github.com/veandco/go-sdl2/sdl"
	"time"
)

func serveEvents(quit chan struct{}) {
	sdl.Do(func() {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				close(quit)
			}
		}
	})
}

func main() {
	sdl.Main(func() {
		vector := raymarcher.D3{}
		fmt.Printf("The X component is %f\n", vector.X)

		marcher := raymarcher.New()
		fmt.Printf("The value of epsilon is %f\n", marcher.Epsilon)

		marcher.World = func(position raymarcher.D3) float64 {
			return raymarcher.Sphere(position, 4.85)
		}

		renderer, err := renderer.New()
		if err != nil {
			panic(err)
		}
		defer renderer.Destroy()

		quit := make(chan struct{})
		pixels := make(chan raymarcher.Pixel)
		tick := time.Tick(time.Second / 100)

		go marcher.Pixels(pixels, tick, quit)
		go renderer.RenderLoop(pixels, tick)

		for {
			select {
			case <-quit:
				break
			default:
				serveEvents(quit)
			}
		}
	})
}
