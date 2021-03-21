package ghosts

type Phantom struct {
}

func (p Phantom) Name() string {
	return "Phantom"
}

func (p Phantom) Evidence() [3]string {
	return [3]string{"Freezing", "EMF 5", "Orbs"}
}
