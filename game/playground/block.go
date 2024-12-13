package pg

import "github.com/ByteArena/box2d"

// Block is a obstacle in the map
// it contains many physical definitions
type Block struct {
	body *box2d.B2Body
}