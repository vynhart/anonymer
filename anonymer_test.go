package anonymer

import "testing"

func TestAnonymize(t *testing.T) {
	name := "Hancock"
	newName := Anonymize(name)
	expectValue := "H***k"
	if newName != expectValue {
		t.Errorf("Expect: %v, Got: %v", expectValue, newName)
	}
}
