package ghosts

type Banshee struct {
}

func (b Banshee) Name() string {
	return "Banshee"
}

func (b Banshee) Evidence() [3]string {
	return [3]string{"Freezing", "EMF 5", "Fingerprints"}
}
