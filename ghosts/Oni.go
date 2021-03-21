package ghosts

type Oni struct {
}

func (o Oni) Name() string {
	return "Oni"
}

func (o Oni) Evidence() [3]string {
	return [3]string{"EMF 5", "Spirit Box", "Writing"}
}
