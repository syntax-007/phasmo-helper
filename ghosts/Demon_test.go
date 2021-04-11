package ghosts

import "testing"

func TestDemon(t *testing.T) {

	result := new(Demon)
	expectedName := "Demon"
	expectedEvidence := [3]string{"Freezing", "Spirit Box", "Writing"}

	t.Run("Demon.Name()", func(t *testing.T) {
		if result.Name() != expectedName {
			t.Errorf("Demon.Name() should equal %v, but instead got %v", expectedName, result.Name())
		}
	})

	t.Run("Demon.Evidence()", func(t *testing.T) {
		if result.Evidence() != expectedEvidence {
			t.Errorf("Demon.Evidence() should equal %v, but instead got %v", expectedEvidence, result.Evidence())
		}
	})

}
