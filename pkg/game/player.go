package game

import (
	"math/rand"
	"time"

	"github.com/google/uuid"
	"github.com/veandco/go-sdl2/sdl"
)

//Player struct
type Player struct {
	UUID       uuid.UUID
	Color      uint32
	X          int32
	Y          int32
	SizeX      int32
	SizeY      int32
	Nanosecond int64
}

//Players is a Map of all players into this map
var Players = make(map[uuid.UUID]*Player)

//PlayersOneUUID are the UUID of local user
var PlayersOneUUID uuid.UUID = uuid.New()

//PlayersTwoUUID are the UUID of local user
var PlayersTwoUUID uuid.UUID = uuid.New()

//CreateLocalPlayer create a local player
func CreateLocalPlayer() {
	createLocalPlayer(PlayersOneUUID)
	createLocalPlayer(PlayersTwoUUID)
}

func createLocalPlayer(UUID uuid.UUID) {
	var min int32 = 10

	Players[UUID] = &Player{
		UUID:       UUID,
		Color:      rand.Uint32(),
		X:          rand.Int31n(ScreamMaxX-min) + min,
		Y:          rand.Int31n(ScreamMaxY-min) + min,
		SizeX:      20,
		SizeY:      20,
		Nanosecond: time.Now().UnixNano(),
	}
}

//DrawPlayers will draw all players
func DrawPlayers(surface *sdl.Surface) {
	for _, p := range Players {
		ps := sdl.Rect{p.X, p.Y, p.SizeX, p.SizeY}
		surface.FillRect(&ps, p.Color)
	}
}
