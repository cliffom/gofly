package main

import (
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/gdamore/tcell/v2"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	// How many flies do we want in our yard?
	numFlies := 50

	// Initialize a new screen
	s, err := tcell.NewScreen()
	if err != nil {
		log.Fatalf("%+v", err)
	}

	if err := s.Init(); err != nil {
		log.Fatalf("%+v", err)
	}

	// Set our default styles
	defStyle := tcell.StyleDefault.Background(tcell.ColorReset).Foreground(tcell.ColorReset)
	s.SetStyle(defStyle)

	// Release the flies!
	width, height := s.Size()
	flies := make([]*Fly, numFlies)
	for i := 0; i < numFlies; i++ {
		flies[i] = NewFly(width, height)
	}

	// Initialize our backyard, attract the flies
	backyard := &Backyard{
		Screen: s,
		Flies:  flies,
	}

	// Run the simulation
	go backyard.Simulate()

	// Enable updates on screen resizing as well as give us an
	// escape hatch to quit the simulation
	for {
		switch event := s.PollEvent().(type) {
		case *tcell.EventResize:
			s.Sync()
		case *tcell.EventKey:
			if event.Key() == tcell.KeyEscape || event.Key() == tcell.KeyCtrlC {
				s.Fini()
				os.Exit(0)
			}
		}
	}

}
