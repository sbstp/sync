package sync

import "sync"

// Event permits one goroutine to send a signal to other goroutines waiting for the event to be set.
type Event struct {
	ch   chan struct{}
	once *sync.Once
}

// NewEvent creates a new event.
func NewEvent() Event {
	return Event{
		ch:   make(chan struct{}),
		once: new(sync.Once),
	}
}

// IsSet checks if the event has been set.
func (e Event) IsSet() bool {
	select {
	case _, ok := <-e.ch:
		return !ok
	default:
		return false
	}
}

// Set sets this event, waking up goroutines waiting on this event.
func (e Event) Set() {
	e.once.Do(func() {
		close(e.ch)
	})
}

// Wait waits for a goroutine to set this event and returns once it's set.
func (e Event) Wait() {
	<-e.ch
}
