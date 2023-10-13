package systems

// System is a system that can be updated.
type System interface {
	Type() string
	Update() error
}
