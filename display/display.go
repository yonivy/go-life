package display

import (
  terminal "github.com/buger/goterm"
  "github.com/yonivy/go-life/grid"
  "github.com/yonivy/go-life/universe"
)

func Display(g grid.Grid) {
  size := g.Size()
  consoleWidth := terminal.Width()

  offset := (consoleWidth/2) - (size/2) + 1

  terminal.Clear()

  for i := 0; i < size; i++ {
    for j := 0; j < size; j++ {
      terminal.MoveCursor(i+1, j+offset)

      if state, _ := g.Get(i, j); state == universe.ALIVE {
        terminal.Print("*")
      } else {
        terminal.Print(" ")
      }
    }
  }

  terminal.Flush()
}
