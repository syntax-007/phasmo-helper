package ghosts

type Mare struct {
}

func (m Mare) Name() string {
	return "Mare"
}

func (m Mare) Evidence() [3]string {
	return [3]string{"Freezing", "Spirit Box", "Orbs"}
}
