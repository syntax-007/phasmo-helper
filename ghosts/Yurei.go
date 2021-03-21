package ghosts

type Yurei struct {
}

func (y Yurei) Name() string {
	return "Yurei"
}

func (y Yurei) Evidence() [3]string {
	return [3]string{"Freezing", "Orbs", "Writing"}
}
