package components

type Updater interface {
	Component
	Update() error
}
