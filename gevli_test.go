package gevli

import (
	"sync"
	"testing"
)

func TestNewEventEmitter(t *testing.T) {
	emitter := NewEventEmitter()
	if emitter == nil {
		t.Errorf("NewEventEmitter returned nil")
	}
	if emitter.listeners == nil {
		t.Errorf("NewEventEmitter did not initialize listeners")
	}
}

func TestAddListener(t *testing.T) {
	emitter := NewEventEmitter()
	eventType := "test"
	listener := func(event Event) {}
	emitter.AddListener(eventType, listener)
	listeners := emitter.listeners[eventType]
	if len(listeners) != 1 {
		t.Errorf("AddListener did not add the listener correctly")
	}
}

func TestEmit(t *testing.T) {
	emitter := NewEventEmitter()
	eventType := "test"
	var wg sync.WaitGroup
	wg.Add(1)
	listener := func(event Event) {
		if event.Type != eventType {
			t.Errorf("Emit sent the wrong event type")
		}
		wg.Done()
	}
	emitter.AddListener(eventType, listener)
	emitter.Emit(eventType, nil)
	wg.Wait()
}

func TestEmitSync(t *testing.T) {
	emitter := NewEventEmitter()

	var wg sync.WaitGroup
	wg.Add(2)

	emitter.AddListener("event", func(event Event) {
		defer wg.Done()
		if event.Data != "test" {
			t.Errorf("expected 'test', got %v", event.Data)
		}
	})

	emitter.AddListener("event", func(event Event) {
		defer wg.Done()
		if event.Data != "test" {
			t.Errorf("expected 'test', got %v", event.Data)
		}
	})

	emitter.EmitSync("event", "test")

	wg.Wait()
}
