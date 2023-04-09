package gevli

import (
	"sync"
)

// Listener is a function that will be called when an event is emitted.
type Listener func(Event)

// Event is a struct that represents an event.
type Event struct {
	Type string
	Data interface{}
}

// EventEmitter is a struct that manages event listeners and emits events.
type EventEmitter struct {
	mu        sync.Mutex
	listeners map[string][]Listener
}

// NewEventEmitter creates a new EventEmitter.
func NewEventEmitter() *EventEmitter {
	return &EventEmitter{
		listeners: make(map[string][]Listener),
	}
}

// AddListener adds a new listener for a specific event type.
func (e *EventEmitter) AddListener(eventType string, listener Listener) {
	e.mu.Lock()
	defer e.mu.Unlock()
	e.listeners[eventType] = append(e.listeners[eventType], listener)
}

// RemoveListener removes a listener for a specific event type.
func (e *EventEmitter) RemoveListener(eventType string, listener Listener) {
	if _, ok := e.listeners[eventType]; !ok {
		return
	}

	listeners := e.listeners[eventType]

	for i := len(listeners) - 1; i >= 0; i-- {
		if &listeners[i] != &listener {
			continue
		}
		e.listeners[eventType] = append(listeners[:i], listeners[i+1:]...)
		break
	}
}

// Emit emits an event to all registered listeners asynchronously.
func (e *EventEmitter) Emit(eventType string, data interface{}) {
	e.mu.Lock()
	defer e.mu.Unlock()
	event := Event{
		Type: eventType,
		Data: data,
	}
	for _, listener := range e.listeners[eventType] {
		go listener(event)
	}
}
