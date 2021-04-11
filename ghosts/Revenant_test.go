package ghosts

import "testing"

func TestRevenant(t *testing.T) {

	result := new(Revenant)
	expectedName := "Revenant"
	expectedEvidence := [3]string{"EMF 5", "Writing", "Fingerprints"}

	t.Run("Revenant.Name()", func(t *testing.T) {
		if result.Name() != expectedName {
			t.Errorf("Revenant.Name() should equal %v, but instead got %v", expectedName, result.Name())
		}
	})

	t.Run("Revenant.Evidence()", func(t *testing.T) {
		if result.Evidence() != expectedEvidence {
			t.Errorf("Revenant.Evidence() should equal %v, but instead got %v", expectedEvidence, result.Evidence())
		}
	})

}
