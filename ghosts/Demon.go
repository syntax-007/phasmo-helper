package ghosts

type Demon struct {
}

func (d Demon) Name() string {
	return "Demon"
}

func (d Demon) Evidence() [3]string {
	return [3]string{"Freezing", "Spirit Box", "Writing"}
}
