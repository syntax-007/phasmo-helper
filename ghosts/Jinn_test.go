package ghosts

import "testing"

func TestJinn(t *testing.T) {

	result := new(Jinn)
	expectedName := "Jinn"
	expectedEvidence := [3]string{"EMF 5", "Orbs", "Spirit Box"}

	t.Run("Jinn.Name()", func(t *testing.T) {
		if result.Name() != expectedName {
			t.Errorf("Jinn.Name() should equal %v, but instead got %v", expectedName, result.Name())
		}
	})

	t.Run("Jinn.Evidence()", func(t *testing.T) {
		if result.Evidence() != expectedEvidence {
			t.Errorf("Jinn.Evidence() should equal %v, but instead got %v", expectedEvidence, result.Evidence())
		}
	})

}
