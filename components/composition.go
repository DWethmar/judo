package components

import "sync"

// Composition is a collection of components.
type Composition struct {
	queueMutex sync.Mutex
	queue      []Component
	components []Component
}

func NewComposition() *Composition {
	return &Composition{
		queueMutex: sync.Mutex{},
		queue:      []Component{},
		components: []Component{},
	}
}

func (c *Composition) InitializeQueue() error {
	if len(c.queue) == 0 {
		return nil
	}

	c.queueMutex.Lock()
	queue := c.queue
	c.queue = []Component{}
	c.queueMutex.Unlock()

	for _, p := range queue {
		if err := p.Init(); err == nil {
			// Note that appending to components array is thread-unsafe,
			// but in the current context of `InitializeQueue`, it's safe
			// because only one goroutine should be initializing the queue.
			c.components = append(c.components, p)
		} else {
			return err
		}
	}

	if len(c.queue) > 0 { // Check if new components were added while initializing.
		return c.InitializeQueue()
	}

	return nil
}

func (c *Composition) Add(component Component) {
	c.queueMutex.Lock()
	defer c.queueMutex.Unlock()

	c.queue = append(c.queue, component)
}

func (c *Composition) Components() []Component {
	return c.components
}
