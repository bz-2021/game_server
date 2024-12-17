package pg

import (
	"fmt"

	"golang.org/x/exp/rand"
)

// Maze represents a 2D maze that is generated using a backtracking algorithm.
// It consists of a grid defined by its Width and Height. The Walls field is a 3D slice
// where each element indicates the presence of obstacles:
// - Walls[0][i][j] indicates the presence of a horizontal barrier at position (i, j).
// - Walls[1][i][j] indicates the presence of a vertical barrier at position (i, j).
type Maze struct {
    Width  int         // Width of the maze
    Height int         // Height of the maze
    Walls  [][][]bool  // store the presence of walls
}

// NewMaze returns a new maze with specific width and height
func NewMaze(width, height int) *Maze {
	if width <= 0 || width > 10 {
		width = 7
	}
	if height <= 0 || height > 10 {
		height = 5
	}
	m := &Maze{
		Width: width,
		Height: height,
		Walls: make([][][]bool, 2),
	}
	m.Resize(width, height)
	return m
}

// Resize will resize the maze with specific width and height 
// and randomly reset the blocks.
func (m *Maze) Resize(width, height int) {
	for i := 0; i < len(m.Walls); i++ {
		rows := height + i
		cols := width + 1 - i;
		m.Walls[i] = make([][]bool, rows)
		for j := 0; j < rows; j++ {
			m.Walls[i][j] = make([]bool, cols)
			for k := 0; k < cols; k++ {
				m.Walls[i][j][k] = true
			}
		}
	}
	m.backtracking()
}

func (m *Maze) backtracking() {
	n := m.Width * m.Height
	current := rand.Intn(n - 1)
	visited := make([]bool, n)
	visited[current] = true
	var stack []int

	for {
		neighbors := m.Neighbors(current)
		var notVisited []int
		for _, n := range neighbors {
			if !visited[n] {
				notVisited = append(notVisited, n)
			}
		}
		if len(notVisited) > 0 {
			next := notVisited[rand.Int() % len(notVisited)]
			*m.Wall(current, next) = false
			visited[next] = true
			stack = append(stack, current)
			current = next
		} else if len(stack) > 0 {
			next := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			current = next
		} else {
			break
		}
	}
}

func (m *Maze) toIndex(x, y int) (i int) {
	return x + y * m.Width
}

func (m *Maze) fromIndex(i int) (x, y int) {
	x = i % m.Width
	y = i / m.Width
	return
}

// Neighbors returns the neighboring cells of cell i.
func (m *Maze) Neighbors(i int) (result []int) {
	x, y := m.fromIndex(i)
	if y > 0 {
		result = append(result, m.toIndex(x, y-1))
	}
	if x > 0 {
		result = append(result, m.toIndex(x-1, y))
	}
	if x+1 < m.Width {
		result = append(result, m.toIndex(x+1, y))
	}
	if y+1 < m.Height {
		result = append(result, m.toIndex(x, y+1))
	}
	return
}

func (m *Maze) WallAbove(x, y int) *bool {
	return &m.Walls[0][x][y]
}

func (m *Maze) WallLeftOf(x, y int) *bool {
	return &m.Walls[1][x][y]
}

func (m *Maze) Wall(i, j int) *bool {
	ix, iy := m.fromIndex(i)
	jx, jy := m.fromIndex(j)
	if iy == jy {
		if ix == jx+1 {
			return m.WallLeftOf(ix, iy)
		}
		if jx == ix+1 {
			return m.WallLeftOf(jx, jy)
		}
	}
	if ix == jx {
		if iy == jy+1 {
			return m.WallAbove(ix, iy)
		}
		if jy == iy+1 {
			return m.WallAbove(jx, jy)
		}
	}
	panic(fmt.Sprintf("Cannot get wall for nodes %v, %v", i, j))
}