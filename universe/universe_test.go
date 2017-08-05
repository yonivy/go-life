package universe

import (
  "testing"
  "github.com/yonivy/go-life/grid"
)

func TestGetNewState(t * testing.T) {
  g, _ := grid.NewGrid(3)
  g.Set(0, 0, ALIVE)
  g.Set(0, 1, DEAD)
  g.Set(0, 2, DEAD)

  g.Set(1, 0, ALIVE)
  g.Set(1, 1, ALIVE)
  g.Set(1, 2, DEAD)

  g.Set(2, 0, ALIVE)
  g.Set(2, 1, DEAD)
  g.Set(2, 2, ALIVE)

  u := NewUniverse(*g)

  cases := []struct {
    row int
    col int
    expected byte
  }{
    {0, 0, ALIVE},
    {0, 1, ALIVE},
    {1, 0, ALIVE},
    {1, 1, DEAD},
    {1, 2, DEAD},
    {2, 2, DEAD},
  }

  for i, c := range cases {
    actual := u.getNewState(c.row, c.col)

    if actual != c.expected {
      t.Error("case", i, "actual:", actual, "expected:", c.expected)
    }
  }
}

func TestEvolve(t * testing.T) {
  g, _ := grid.NewGrid(2)
  g.Set(0, 0, ALIVE)
  g.Set(0, 1, ALIVE)
  g.Set(1, 0, DEAD)
  g.Set(1, 1, ALIVE)

  u := NewUniverse(*g)

  g_future, _ := grid.NewGrid(2)
  g_future.Set(0, 0, ALIVE)
  g_future.Set(0, 1, ALIVE)
  g_future.Set(1, 0, ALIVE)
  g_future.Set(1, 1, ALIVE)

  cases := []struct {
    expected *grid.Grid
  }{
    {g_future},
    {g_future},
  }

  for i, c := range cases {
    actual := u.Evolve()

    if !actual.Equals(c.expected) {
      t.Error("case", i, "grids aren't equal. \nactual:", actual, "\nexpected:", c.expected)
    }
  }
}
