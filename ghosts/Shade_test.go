package ghosts

import "testing"

func TestShade(t *testing.T) {

	result := new(Shade)
	expectedName := "Shade"
	expectedEvidence := [3]string{"EMF 5", "Orbs", "Writing"}

	t.Run("Shade.Name()", func(t *testing.T) {
		if result.Name() != expectedName {
			t.Errorf("Shade.Name() should equal %v, but instead got %v", expectedName, result.Name())
		}
	})

	t.Run("Shade.Evidence()", func(t *testing.T) {
		if result.Evidence() != expectedEvidence {
			t.Errorf("Shade.Evidence() should equal %v, but instead got %v", expectedEvidence, result.Evidence())
		}
	})

}
