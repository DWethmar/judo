package components

import (
	"errors"
	"reflect"
	"sync"
	"testing"
)

func TestNewComposition(t *testing.T) {
	tests := []struct {
		name string
		want *Composition
	}{
		{
			name: "NewComposition",
			want: &Composition{
				queueMutex: sync.Mutex{},
				queue:      []Component{},
				components: []Component{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewComposition(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewComposition() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestComposition_InitializeQueue(t *testing.T) {
	type fields struct {
		queueMutex sync.Mutex
		queue      []Component
		components []Component
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "initialize without components should not return an error",
			fields: fields{
				queueMutex: sync.Mutex{},
				queue:      []Component{},
				components: []Component{},
			},
			wantErr: false,
		},
		{
			name: "initialize with components should not return an error",
			fields: fields{
				queueMutex: sync.Mutex{},
				queue: []Component{&MockComponent{
					TypeFunc: func() string { return "mock" },
					InitFunc: func() error { return nil },
				}},
				components: []Component{},
			},
			wantErr: false,
		},
		{
			name: "should return an error when component init fails",
			fields: fields{
				queueMutex: sync.Mutex{},
				queue: []Component{&MockComponent{
					TypeFunc: func() string { return "mock" },
					InitFunc: func() error { return errors.New("mock error") },
				}},
				components: []Component{},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Composition{
				queueMutex: tt.fields.queueMutex,
				queue:      tt.fields.queue,
				components: tt.fields.components,
			}
			if err := c.InitializeQueue(); (err != nil) != tt.wantErr {
				t.Errorf("Composition.InitializeQueue() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestComposition_Add(t *testing.T) {
	t.Run("should add component to queue", func(t *testing.T) {
		c := NewComposition()
		c.Add(&MockComponent{
			TypeFunc: func() string { return "mock" },
			InitFunc: func() error { return nil },
		})

		if len(c.queue) != 1 {
			t.Errorf("Composition.Add() = %v, want %v", len(c.queue), 1)
		}

		if c.queue[0].Type() != "mock" {
			t.Errorf("Composition.Add() = %v, want %v", c.queue[0].Type(), "mock")
		}
	})
}

func TestComposition_Components(t *testing.T) {
	t.Run("should return components", func(t *testing.T) {
		c := NewComposition()
		c.Add(&MockComponent{
			TypeFunc: func() string { return "mock" },
			InitFunc: func() error { return nil },
		})

		if len(c.Components()) != 0 {
			t.Errorf("Composition.Components() = %v, want %v", len(c.Components()), 0)
		}

		if err := c.InitializeQueue(); err != nil {
			t.Errorf("Composition.InitializeQueue() error = %v", err)
		}

		if len(c.Components()) != 1 {
			t.Errorf("Composition.Components() = %v, want %v", len(c.Components()), 1)
		}

		if c.Components()[0].Type() != "mock" {
			t.Errorf("Composition.Components() = %v, want %v", c.Components()[0].Type(), "mock")
		}

		if len(c.queue) != 0 {
			t.Errorf("Composition.Components() = %v, want %v", len(c.queue), 0)
		}
	})
}
