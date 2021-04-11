package ghosts

import "testing"

func TestMare(t *testing.T) {

	result := new(Mare)
	expectedName := "Mare"
	expectedEvidence := [3]string{"Freezing", "Spirit Box", "Orbs"}

	t.Run("Mare.Name()", func(t *testing.T) {
		if result.Name() != expectedName {
			t.Errorf("Mare.Name() should equal %v, but instead got %v", expectedName, result.Name())
		}
	})

	t.Run("Mare.Evidence()", func(t *testing.T) {
		if result.Evidence() != expectedEvidence {
			t.Errorf("Mare.Evidence() should equal %v, but instead got %v", expectedEvidence, result.Evidence())
		}
	})

}
