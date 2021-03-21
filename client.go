package main

import (
	"fmt"
	. "ghosts"
	"reflect"
)

func main() {
	ghosts := []Ghost{new(Banshee), new(Demon), new(Jinn), new(Mare), new(Oni), new(Poltergeist), new(Revenant), new(Shade), new(Spirit), new(Yurei), new(Wraith), new(Phantom)}

	for _, g := range ghosts {
		fmt.Println(reflect.TypeOf(g))
	}

}
