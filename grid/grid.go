package grid

import (
  "bytes"
  "errors"
)

type Grid struct {
  size int
  cells [][]byte
}

const (
  DEAD byte = iota
  ALIVE
)

func NewGrid(size int) (*Grid, error) {
  if size < 1 {
    return nil, errors.New("size must be a positive number")
  }

  return &Grid {
    size: size,
    cells: newCells(size),
  }, nil
}

func newCells(size int) [][]byte {
  cells := make([][]byte, size)

  for i := range cells {
    cells[i] = make([]byte, size)
  }

  return cells
}

func (g *Grid) Size() int {
  return g.size
}

func (g *Grid) Get(x, y int) (byte, error) {
  if err := g.checkBounderies(x, y); err != nil {
    return 0, err
  }

  return g.cells[x][y], nil
}

func (g *Grid) Set(x, y int, state byte) error {
  if err := g.checkBounderies(x, y); err != nil {
    return err
  }

  g.cells[x][y] = state

  return nil
}

func (g *Grid) checkBounderies(x, y int) error {
  if x < 0 || x >= g.Size() {
    return errors.New("`x` is out of range")
  }

  if y < 0 || y >= g.Size() {
    return errors.New("`y` is out of range")
  }

  return nil
}

// summation of all the values in a cell's (logical) bounding rectangle
func (g *Grid) SumBoundingRect(x, y int) byte {
  var sum byte = 0

  xStart, xFinish := x-1, x+2
  yStart, yFinish := y-1, y+2

  // verify x bounds are valid
  if xStart < 0 {
    xStart = 0
  }
  if xFinish > g.Size() {
    xFinish = g.Size()
  }

  // verify y bounds are valid
  if yStart < 0 {
    yStart = 0
  }
  if yFinish > g.Size() {
    yFinish = g.Size()
  }

  // and now sum all valid cells
  for i := xStart; i < xFinish; i++ {
    for j := yStart; j < yFinish; j++ {
      state, _ := g.Get(i, j)
      sum += state
    }
  }

  return sum
}

func (g *Grid) Equals(other *Grid) bool {
  for i, _ := range g.cells {
    if !bytes.Equal(g.cells[i], other.cells[i]) {
      return false
    }
  }

  return true
}
