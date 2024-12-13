package conf

type MapConf struct {
	Width        int // The number of horizontal blocks.
	Height       int // The number of vertical blocks.
	MaxPlayerNum int // The maximum number of players that the map can contain.
	BlockLength  int // The length of each block.
	BlockWidth   int // The width of each block.

	RandomSize bool // If this value is true, the map will generate a maze with random width and height.
}
