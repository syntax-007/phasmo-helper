package ghosts

type Poltergeist struct {
}

func (p Poltergeist) Name() string {
	return "Poltergeist"
}

func (p Poltergeist) Evidence() [3]string {
	return [3]string{"Orbs", "Spirit Box", "Fingerprints"}
}
