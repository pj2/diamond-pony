package renderer

import (
	"github.com/pj2/diamond-pony/raymarcher"
	"github.com/veandco/go-sdl2/sdl"
	"time"
)

// Renderer draws pixels to the screen.
type Renderer struct {
	Window  *sdl.Window
	Surface *sdl.Surface
}

// New creates a Renderer (and also initializes SDL) with a default window.
func New() (Renderer, error) {
	var window *sdl.Window
	var surface *sdl.Surface
	var err error

	sdl.Do(func() {
		sdl.Init(sdl.INIT_EVERYTHING)

		window, err = sdl.CreateWindow("Diamond Pony", sdl.WINDOWPOS_UNDEFINED,
			sdl.WINDOWPOS_UNDEFINED, 300, 300, sdl.WINDOW_SHOWN)
	})

	if err != nil {
		return Renderer{}, err
	}

	sdl.Do(func() {
		surface, err = window.GetSurface()
	})
	if err != nil {
		return Renderer{}, err
	}

	return Renderer{
		Window:  window,
		Surface: surface,
	}, nil
}

// Render draws pixels to the screen until c is closed.
func (r *Renderer) RenderLoop(c chan raymarcher.Pixel, tick <-chan time.Time) {
	var err error
	for {
		select {
		case px := <-c:
			sq := sdl.Rect{
				int32(px.X),
				int32(px.Y),
				1, // TODO Scale properly from worldspace to screenspace
				1,
			}
			sdl.Do(func() { // FIXME This is probably slow
				if !px.Empty() {
					r.Surface.FillRect(&sq, 0xFFFFFFFF)
				} else {
					r.Surface.FillRect(&sq, 0x00000000)
				}
			})
		case <-tick:
			sdl.Do(func() {
				err = r.Window.UpdateSurface()

				// Clear the screen
				// r.Surface.FillRect(&sdl.Rect{0, 0, r.Surface.W, r.Surface.H},
				// 	0)
			})
			if err != nil {
				panic(err) // FIXME
			}
		}
	}
}

// Destroy closes all resources safely.
func (r *Renderer) Destroy() {
	sdl.Do(func() {
		r.Window.Destroy()
		sdl.Quit()
	})
}
