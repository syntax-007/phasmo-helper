package ghosts

type Jinn struct {
}

func (j Jinn) Name() string {
	return "Jinn"
}

func (j Jinn) Evidence() [3]string {
	return [3]string{"EMF 5", "Orbs", "Spirit Box"}
}
