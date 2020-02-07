package game

import (
	"github.com/veandco/go-sdl2/sdl"
)

var ScreamMaxX int32 = 1200
var ScreamMaxY int32 = 800

//Play metod start the Game Loop
func Play() {

	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		panic(err)
	}
	defer sdl.Quit()

	window, err := sdl.CreateWindow("Go Game GoGO",
		sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		ScreamMaxX, ScreamMaxY, sdl.WINDOW_SHOWN)

	if err != nil {
		panic(err)
	}

	defer window.Destroy()
	surface, err := window.GetSurface()
	if err != nil {
		panic(err)
	}

	running := true
	for running {

		surface.FillRect(nil, 0)
		DrawPlayers(surface)
		window.UpdateSurface()

		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch t := event.(type) {
			case *sdl.KeyboardEvent:
				HandleKeyboardEvents(t)
			case *sdl.QuitEvent:
				println("Quit")
				running = false
				break
			}
		}
	}
}
