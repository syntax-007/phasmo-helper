package ghosts

type Shade struct {
}

func (s Shade) Name() string {
	return "Shade"
}

func (s Shade) Evidence() [3]string {
	return [3]string{"EMF 5", "Orbs", "Writing"}
}
