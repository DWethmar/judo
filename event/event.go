package event

import "fmt"

type Bus[E any] struct {
	subscriptions []func(event E) error
}

func (e *Bus[E]) Send(event E) error {
	for _, callback := range e.subscriptions {
		if err := callback(event); err != nil {
			return fmt.Errorf("error sending event %T: %w", event, err)
		}
	}

	return nil
}

func (e *Bus[E]) Subscribe(callback func(event E) error) {
	e.subscriptions = append(e.subscriptions, callback)
}

func New[E any]() *Bus[E] {
	return &Bus[E]{
		subscriptions: []func(event E) error{},
	}
}
