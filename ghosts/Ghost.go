package ghosts

type Ghost interface {
	Name() string
	Evidence() [3]string
}
