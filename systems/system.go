package systems

// System is a system that can be updated.
type System interface {
	Update() error
}
