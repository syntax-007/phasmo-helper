package ghosts

type Spirit struct {
}

func (s Spirit) Name() string {
	return "Spirit"
}

func (s Spirit) Evidence() [3]string {
	return [3]string{"Spirit Box", "Writing", "Fingerprints"}
}
