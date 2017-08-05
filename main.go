package main

import (
  "fmt"
  "os"
  "os/signal"
  "syscall"
  "flag"
  "time"
  "github.com/yonivy/go-life/seeds"
  "github.com/yonivy/go-life/universe"
  console "github.com/yonivy/go-life/display"
)

func main() {
  handleSigterm()

  seedFn := seeds.Seeds[getSeedName()]

  seed := seedFn()

  u := universe.NewUniverse(*seed)

  // showing the first generation e.g. seed
  console.Display(*seed)
  time.Sleep(time.Second)

  // create future generations indefinitly
  for {
    future := u.Evolve()
    console.Display(future)
    time.Sleep(time.Second)
  }
}

func getSeedName() string {
  const DEFAULT_SEED string = "beacon"

  inputSeed := flag.String("seed", DEFAULT_SEED, "a seed (pattern) to start with")

  flag.Parse()

  if _, ok := seeds.Seeds[*inputSeed]; ok {
    return *inputSeed
  }

  return DEFAULT_SEED
}

func handleSigterm() {
  c := make(chan os.Signal, 2)
  signal.Notify(c, os.Interrupt, syscall.SIGTERM)

  go func() {
    <-c
    fmt.Println("\n\nEnding all life... Bye Bye.")
    os.Exit(0)
  }()
}
