package ghosts

import "testing"

func TestYurei(t *testing.T) {

	result := new(Yurei)
	expectedName := "Yurei"
	expectedEvidence := [3]string{"Freezing", "Orbs", "Writing"}

	t.Run("Yurei.Name()", func(t *testing.T) {
		if result.Name() != expectedName {
			t.Errorf("Yurei.Name() should equal %v, but instead got %v", expectedName, result.Name())
		}
	})

	t.Run("Yurei.Evidence()", func(t *testing.T) {
		if result.Evidence() != expectedEvidence {
			t.Errorf("Yurei.Evidence() should equal %v, but instead got %v", expectedEvidence, result.Evidence())
		}
	})

}
