package grid

import "testing"

func TestNewGrid(t *testing.T) {
  cases := []struct {
    size int
    valid bool
  }{
    {1, true},
    {0, false},
    {-1, false},
    {1.0, false},
  }

  for i, c := range cases {
    _, err := NewGrid(c.size)

    if err != nil && c.valid {
      t.Error("case", i, "errored with:", err)
    }
  }
}

func TestCheckBounderies(t *testing.T) {
  cases := []struct {
    x, y int
    valid bool
  }{
    {0, 0, true},
    {1, 1, true},
    {-1, 0, false},
    {0, -1, false},
  }

  grid, _ := NewGrid(2)

  for i, c := range cases {
    err := grid.Set(c.x, c.y, 1)

    if err != nil && c.valid {
      t.Error("case", i, "errored with:", err)
    }
  }
}

func TestSet(t *testing.T) {
  cases := []struct {
    x, y int
    valid bool
  }{
    {0, 0, true},
    {0, -1, false},
    {-1, 0, false},
  }

  grid, _ := NewGrid(2)

  for i, c := range cases {
    err := grid.Set(c.x, c.y, 1)

    if err != nil && c.valid {
      t.Error("case", i, "Set errored with:", err)
    }
  }
}

func TestGet(t *testing.T) {
  cases := []struct {
    x, y int
    state byte
  }{
    {0, 0, 1},
    {0, 1, 1},
    {1, 0, 0},
    {1, 1, 0},
  }

  grid, _ := NewGrid(2)
  grid.Set(0, 0, 1)
  grid.Set(0, 1, 1)
  grid.Set(1, 0, 0)
  grid.Set(1, 1, 0)

  for i, c := range cases {
    state, err := grid.Get(c.x, c.y)

    if err != nil {
      t.Error("case", i, "errored with:", err)
    }

    if state != c.state {
      t.Error("case", i, "state mismatch. actual:", state, "expected:", c.state)
    }
  }
}

func TestSumBoundingRect(t *testing.T) {
  const GRID_SIZE int = 2
  const EXPECTED_SUM byte = 2

  grid, _ := NewGrid(2)
  grid.Set(0, 0, 1)
  grid.Set(0, 1, 1)
  grid.Set(1, 0, 0)
  grid.Set(1, 1, 0)

  for i := 0; i < GRID_SIZE; i++ {
    for j := 0; j < GRID_SIZE; j++ {
      actual := grid.SumBoundingRect(i,  j)

      if actual != EXPECTED_SUM {
        t.Error("no good... actual:", actual, "expected:", EXPECTED_SUM)
      }
    }
  }
}

func TestEquals(t *testing.T) {
  grid1, _ := NewGrid(2)
  grid1.Set(0, 0, 1)
  grid1.Set(0, 1, 1)
  grid1.Set(1, 0, 0)
  grid1.Set(1, 1, 0)

  grid2, _ := NewGrid(2)
  grid2.Set(0, 0, 1)
  grid2.Set(0, 1, 1)
  grid2.Set(1, 0, 0)
  grid2.Set(1, 1, 0)

  if !grid1.Equals(grid2) {
    t.Error("no good... grids aren't equal")
  }
}
