package main

import (
	. "ghosts"
)

func main() {
	ghosts := []Ghost{new(Banshee), new(Demon), new(Jinn), new(Mare), new(Oni), new(Poltergeist), new(Revenant), new(Shade), new(Spirit), new(Yurei), new(Wraith), new(Phantom)}
	m := make(map[string][3]string)
	for _, g := range ghosts {
		m[g.Name()] = g.Evidence()
	}

}
