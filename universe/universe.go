package universe

import "github.com/yonivy/go-life/grid"

const (
  DEAD byte = iota
  ALIVE
)

type universe struct {
  present *grid.Grid
  future *grid.Grid
}

// create a new universe with a given `seed`
func NewUniverse(seed grid.Grid) *universe {
  empty, _ := grid.NewGrid(seed.Size())

  return &universe {
    present: &seed,
    future: empty,
  }
}

// create one future generation of a universe
func (u *universe) Evolve() grid.Grid {
  size := u.present.Size()

  // create a new generation based on present one
  for x := 0; x < size; x++ {
    for y := 0; y < size; y++ {
      state := u.getNewState(x, y)
      u.future.Set(x, y, state)
    }
  }

  // swap the grids
  u.present, u.future = u.future, u.present

  return (*u.present)
}

// decide the next state of a cell given its present
func (u *universe) getNewState(row, col int) byte {
  sum := u.present.SumBoundingRect(row, col)

  newState := DEAD

  // a closer observation enables us to have a minimal decision tree.
  // see: https://en.wikipedia.org/wiki/Conway%27s_Game_of_Life#Algorithms
  if sum == 3 {
    newState = ALIVE
  } else if sum == 4 {
    oldState, _ := u.present.Get(row, col)
    newState = oldState
  }

  return newState
}
