package components

type Component interface {
	Type() string
	Init() error // Called once when the component is added to an entity.
}

// MockComponent is a mock component.
type MockComponent struct {
	TypeFunc func() string
	InitFunc func() error
}

func (m *MockComponent) Type() string { return m.TypeFunc() }
func (m *MockComponent) Init() error  { return m.InitFunc() }
