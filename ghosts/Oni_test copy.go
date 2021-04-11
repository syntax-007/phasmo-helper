package ghosts

import "testing"

func TestOni(t *testing.T) {

	result := new(Oni)
	expectedName := "Oni"
	expectedEvidence := [3]string{"EMF 5", "Spirit Box", "Writing"}

	t.Run("Oni.Name()", func(t *testing.T) {
		if result.Name() != expectedName {
			t.Errorf("Oni.Name() should equal %v, but instead got %v", expectedName, result.Name())
		}
	})

	t.Run("Oni.Evidence()", func(t *testing.T) {
		if result.Evidence() != expectedEvidence {
			t.Errorf("Oni.Evidence() should equal %v, but instead got %v", expectedEvidence, result.Evidence())
		}
	})

}
