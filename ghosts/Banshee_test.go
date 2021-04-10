package ghosts

import "testing"

func TestBanshee(t *testing.T) {

	result := new(Banshee)
	expectedName := "Banshee"
	expectedEvidence := [3]string{"Freezing", "EMF 5", "Fingerprints"}

	t.Run("Banshee.Name()", func(t *testing.T) {
		if result.Name() != expectedName {
			t.Errorf("Banshee.Name() should equal %v, but instead got %v", expectedName, result.Name())
		}
	})

	t.Run("Banshee.Evidence()", func(t *testing.T) {
		if result.Evidence() != expectedEvidence {
			t.Errorf("Banshee.Evidence() should equal %v, but instead got %v", expectedEvidence, result.Evidence())
		}
	})

}
