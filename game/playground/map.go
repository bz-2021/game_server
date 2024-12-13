package pg

import (
	"time"

	"github.com/ByteArena/box2d"
	"github.com/bz-2021/game_server/conf"
	"golang.org/x/exp/rand"
)

// Map represents a game map that includes a physics world, a maze, a list of players, etc..
type Map struct {
	world      box2d.B2World // the Box2D world which handles all physical simulations.
	maze       *Maze         // Pointer to the Maze struct which represents the 2D maze in the game.
	playerList []Player      // Slice of Player structs representing the players in the game.

	cfg conf.MapConf
}

// NewMap creates a new Map instance using the provided configuration settings.
func NewMap(cfg conf.MapConf) *Map {
	rand.Seed(uint64(time.Now().UnixNano()))
	m := &Map{
		world: box2d.MakeB2World(box2d.MakeB2Vec2(0, 0)),
		cfg:   cfg,
	}
	if cfg.RandomSize {
		m.maze = NewMaze(rand.Int()%10, rand.Int()%10)
	} else if cfg.Width <= 0 || cfg.Width > 10 || cfg.Height <= 0 || cfg.Height > 10 {
		m.maze = NewDefaultMaze()
	} else {
		m.maze = NewMaze(cfg.Width, cfg.Height)
	}
	return m
}

// AddPlayer adds a player to the map if the maximum number of players has not been reached.
// It returns a boolean value indicating whether the player was successfully added.
func (m *Map) AddPlayer(p Player) bool {
	if len(m.playerList) > m.cfg.MaxPlayerNum {
		return false
	}
	m.playerList = append(m.playerList, p)
	return true
}
