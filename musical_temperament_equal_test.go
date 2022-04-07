package go_hugipipes_musical_notes

import "testing"

func TestMTemperamentEqualInterface(t *testing.T) {
	te := NewMTemperamentEqual(440.0)
	checkTemperamentInterface(*te)
}
