package seeds

import (
  "github.com/yonivy/go-life/grid"
  "github.com/yonivy/go-life/universe"
)

func Blinker() *grid.Grid {
  g, _ := grid.NewGrid(5)

  g.Set(2, 1, universe.ALIVE)
  g.Set(2, 2, universe.ALIVE)
  g.Set(2, 3, universe.ALIVE)

  return g
}

func Beacon() *grid.Grid {
  g, _ := grid.NewGrid(6)

  g.Set(1, 1, universe.ALIVE)
  g.Set(1, 2, universe.ALIVE)
  g.Set(2, 1, universe.ALIVE)
  g.Set(3, 4, universe.ALIVE)
  g.Set(4, 3, universe.ALIVE)
  g.Set(4, 4, universe.ALIVE)

  return g
}

var Seeds map[string]func() *grid.Grid = map[string]func() *grid.Grid {
  "beacon": Beacon,
  "blinker": Blinker,
}
