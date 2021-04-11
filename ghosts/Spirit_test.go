package ghosts

import "testing"

func TestSpirit(t *testing.T) {

	result := new(Spirit)
	expectedName := "Spirit"
	expectedEvidence := [3]string{"Spirit Box", "Writing", "Fingerprints"}

	t.Run("Spirit.Name()", func(t *testing.T) {
		if result.Name() != expectedName {
			t.Errorf("Spirit.Name() should equal %v, but instead got %v", expectedName, result.Name())
		}
	})

	t.Run("Spirit.Evidence()", func(t *testing.T) {
		if result.Evidence() != expectedEvidence {
			t.Errorf("Spirit.Evidence() should equal %v, but instead got %v", expectedEvidence, result.Evidence())
		}
	})

}
