package ghosts

type Revenant struct {
}

func (r Revenant) Name() string {
	return "Revenant"
}

func (r Revenant) Evidence() [3]string {
	return [3]string{"EMF 5", "Writing", "Fingerprints"}
}
