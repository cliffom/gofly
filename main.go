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

	rand.Seed(time.Now().UnixNano())

	// Listen for events
	event := make(chan tcell.Event)
	quit := make(chan struct{})
	go s.ChannelEvents(event, quit)

	// Release the flies!
	flies := make([]*Fly, *numFlies)
	for i := 0; i < *numFlies; i++ {
		flies[i] = NewFly(s.Size())
	}

	// Initialize our b, attract the flies
	b := NewBackyard(s, flies, time.Duration(*frameTime))

	// Run the simulation
	go b.Simulate()

	// Enable updates on screen resizing as well as give us an
	// escape hatch to quit the simulation
	for {
		eventHandler(event, quit, s, b)
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

func eventHandler(event chan tcell.Event, quit chan struct{}, s tcell.Screen, b *Backyard) {
	select {
	case ev := <-event:
		switch event := ev.(type) {
		case *tcell.EventResize:
			s.Sync()
		case *tcell.EventKey:
			handleKeyPress(event.Key(), s, b)
		}
	case <-quit:
		s.Fini()
		os.Exit(0)
	}
}

// handleKeyPress handles key press events from the user
func handleKeyPress(k tcell.Key, s tcell.Screen, b *Backyard) {
	switch k {
	case tcell.KeyUp:
		if !b.SpeedUp() {
			s.Beep()
		}
	case tcell.KeyDown:
		if !b.SpeedDown() {
			s.Beep()
		}
	case tcell.KeyRight:
		b.AddFly()
	case tcell.KeyLeft:
		if !b.RemoveFly() {
			s.Beep()
		}
	case tcell.KeyDelete:
		if !b.RemoveFlies() {
			s.Beep()
		}
	case tcell.KeyEscape:
		s.Clear()
		s.Fini()
		os.Exit(0)
	}
}
