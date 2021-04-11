package ghosts

import "testing"

func TestWraith(t *testing.T) {

	result := new(Wraith)
	expectedName := "Wraith"
	expectedEvidence := [3]string{"Freezing", "Spirit Box", "Fingerprints"}

	t.Run("Wraith.Name()", func(t *testing.T) {
		if result.Name() != expectedName {
			t.Errorf("Wraith.Name() should equal %v, but instead got %v", expectedName, result.Name())
		}
	})

	t.Run("Wraith.Evidence()", func(t *testing.T) {
		if result.Evidence() != expectedEvidence {
			t.Errorf("Wraith.Evidence() should equal %v, but instead got %v", expectedEvidence, result.Evidence())
		}
	})

}
