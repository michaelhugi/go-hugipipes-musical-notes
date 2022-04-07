package go_hugipipes_musical_notes

import "testing"

func TestMNoteEqualInterface(t *testing.T) {
	n := MNoteEqual{}
	checkMNInterface(n)
}
