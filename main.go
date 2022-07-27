package main

import (
	"flag"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/gdamore/tcell/v2"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	// How many flies do we want in our yard?
	numFlies := flag.Int("flies", 50, "How many flies are you inviting?")
	frameTime := flag.Int("frametime", 250, "How fast do you want your flies to fly (must be at least 100)?")
	flag.Parse()

	if *frameTime < 100 {
		log.Fatalf("error: frametime must be at least 100")
	}

	// Initiatlize our screen
	s, err := getScreen()
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	// Release the flies!
	flies := make([]*Fly, *numFlies)
	for i := 0; i < *numFlies; i++ {
		flies[i] = NewFly(s.Size())
	}

	// Initialize our backyard, attract the flies
	backyard := NewBackyard(s, flies, time.Duration(*frameTime))

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

func getScreen() (tcell.Screen, error) {
	// Initialize a new screen
	s, err := tcell.NewScreen()
	if err != nil {
		return nil, err
	}

	if err := s.Init(); err != nil {
		return nil, err
	}

	// Set our default styles
	defStyle := tcell.StyleDefault.Background(tcell.ColorReset).Foreground(tcell.ColorReset)
	s.SetStyle(defStyle)

	return s, nil
}
