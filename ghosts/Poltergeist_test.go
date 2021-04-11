package ghosts

import "testing"

func TestPoltergeist(t *testing.T) {

	result := new(Poltergeist)
	expectedName := "Poltergeist"
	expectedEvidence := [3]string{"Orbs", "Spirit Box", "Fingerprints"}

	t.Run("Poltergeist.Name()", func(t *testing.T) {
		if result.Name() != expectedName {
			t.Errorf("Poltergeist.Name() should equal %v, but instead got %v", expectedName, result.Name())
		}
	})

	t.Run("Poltergeist.Evidence()", func(t *testing.T) {
		if result.Evidence() != expectedEvidence {
			t.Errorf("Poltergeist.Evidence() should equal %v, but instead got %v", expectedEvidence, result.Evidence())
		}
	})

}
