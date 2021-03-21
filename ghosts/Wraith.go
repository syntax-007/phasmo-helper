package ghosts

type Wraith struct {
}

func (w Wraith) Name() string {
	return "Wraith"
}

func (w Wraith) Evidence() [3]string {
	return [3]string{"Freezing", "Spirit Box", "Fingerprints"}
}
