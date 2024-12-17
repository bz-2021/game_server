package pg

import (
	"time"

	"github.com/ByteArena/box2d"
	"github.com/bz-2021/game_server/conf"
	"golang.org/x/exp/rand"
)

type Block = box2d.B2Body

// Map represents a game map that includes a physics world, a maze, a list of players, etc..
type Map struct {
	world      box2d.B2World // the Box2D world which handles all physical simulations.
	maze       *Maze         // Pointer to the Maze struct which represents the 2D maze in the game.
	playerList []Player      // Slice of Player structs representing the players in the game.
	blockList  []*Block      // Slice of the obstacle created in Box2D world

	cfg conf.MapConf
}

// NewMap creates a new Map instance using the provided configuration settings.
func NewMap(cfg conf.MapConf) *Map {
	m := &Map{
		world: box2d.MakeB2World(box2d.MakeB2Vec2(0, 0)),
		cfg:   cfg,
	}
	if cfg.RandomSize {
		rand.Seed(uint64(time.Now().UnixNano()))
		m.maze = NewMaze(rand.Int()%10, rand.Int()%10)
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

func (m *Map) createWorld() {
	worldGravity := box2d.MakeB2Vec2(0, 0)
	m.world = box2d.MakeB2World(worldGravity)
}

// CreateBlocks creates obstacles' box2d body
func (m *Map) CreateBlocks() {
	blockDef := box2d.MakeB2BodyDef()
	blockDef.Type = box2d.B2BodyType.B2_staticBody

	for x := 0; x < len(m.maze.Walls[0]); x++ {
		for y := 0; y < len(m.maze.Walls[0][x]); y++ {
			if m.maze.Walls[0][x][y] {
				blockDef.Position = box2d.MakeB2Vec2()
			}
		}
	}
	for x := 0; x < len(m.maze.Walls[1]); x++ {
		for y := 0; y < len(m.maze.Walls[1][x]); y++ {
			if m.maze.Walls[1][x][y] {

			}
		}
	}
}
