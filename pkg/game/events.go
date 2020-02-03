package game

import "github.com/veandco/go-sdl2/sdl"

//HandleKeyboardEvents is a event handler
func HandleKeyboardEvents(t *sdl.KeyboardEvent) {
	handleLocalPlayersKeys(t)
}

func handleLocalPlayersKeys(t *sdl.KeyboardEvent) {

	var step int32 = 10

	switch t.Keysym.Sym {
	// Players One
	case sdl.K_a:
		if t.State == sdl.RELEASED {
			PlayersLocal[0].Position.X -= step
		}
	case sdl.K_s:
		if t.State == sdl.RELEASED {
			PlayersLocal[0].Position.Y += step
		}
	case sdl.K_d:
		if t.State == sdl.RELEASED {
			PlayersLocal[0].Position.X += step
		}
	case sdl.K_w:
		if t.State == sdl.RELEASED {
			PlayersLocal[0].Position.Y -= step
		}

	//Player TWO
	case sdl.K_LEFT:
		if t.State == sdl.RELEASED {
			PlayersLocal[1].Position.X -= step
		}
	case sdl.K_DOWN:
		if t.State == sdl.RELEASED {
			PlayersLocal[1].Position.Y += step
		}
	case sdl.K_RIGHT:
		if t.State == sdl.RELEASED {
			PlayersLocal[1].Position.X += step
		}
	case sdl.K_UP:
		if t.State == sdl.RELEASED {
			PlayersLocal[1].Position.Y -= step
		}
	}
}
