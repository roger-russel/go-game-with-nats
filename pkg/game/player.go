package game

import "github.com/veandco/go-sdl2/sdl"

//Position of player
type Position struct {
	X int32
	Y int32
}

//Player struct
type Player struct {
	ID       string
	Color    uint32
	Position Position
}

//var PlayersRemote = map[string]*Player{}

//PlayersLocal are the players running locally
var PlayersLocal = []*Player{
	{
		ID:    "4cb1c289-c0d6-49c0-b225-78b22364ab0f",
		Color: 0xffeeff00,
		Position: Position{
			20,
			20,
		},
	},
	{
		ID:    "5cb1c289-c0d6-49c0-b225-78b22364ab0f",
		Color: 0xffffffff,
		Position: Position{
			120,
			120,
		},
	},
}

//DrawPlayers will draw all players
func DrawPlayers(surface *sdl.Surface) {

	for _, p := range PlayersLocal {
		pl := sdl.Rect{p.Position.X, p.Position.Y, 20, 20}
		surface.FillRect(&pl, p.Color)
	}

}
