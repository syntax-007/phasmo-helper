package ghosts

import "testing"

func TestPhantom(t *testing.T) {

	result := new(Phantom)
	expectedName := "Phantom"
	expectedEvidence := [3]string{"Freezing", "EMF 5", "Orbs"}

	t.Run("Phantom.Name()", func(t *testing.T) {
		if result.Name() != expectedName {
			t.Errorf("Phantom.Name() should equal %v, but instead got %v", expectedName, result.Name())
		}
	})

	t.Run("Phantom.Evidence()", func(t *testing.T) {
		if result.Evidence() != expectedEvidence {
			t.Errorf("Phantom.Evidence() should equal %v, but instead got %v", expectedEvidence, result.Evidence())
		}
	})

}
