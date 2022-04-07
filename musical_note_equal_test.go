package go_hugipipes_musical_notes

import (
	"testing"
)

func TestMNoteEqualInterface(t *testing.T) {
	n := MNoteEqual{}
	checkMNInterface(n)
}

func TestMNoteEqualF2(t *testing.T) {
	/*o2 := NewMTemperamentEqual(440)
	n1, err := newMNoteEqual(o2, 87.31)
	tu.AssertNErr(err)

	tu.AssertEq(roundFloat(n1.LowerFrequency(), 2), 84.82, t)
	tu.AssertEq(roundFloat(n1.UpperFrequency(), 2), 89.87, t)
	tu.AssertEq(n1.ExactFrequency(), 87.31, t)
	tu.Assert(n1.FrequencyBelongsToNote(84.83), t)
	tu.Assert(!n1.FrequencyBelongsToNote(84.81), t)
	tu.Assert(!n1.FrequencyBelongsToNote(89.87), t)
	tu.Assert(n1.FrequencyBelongsToNote(89.86), t)
	tu.AssertEq(n1.MidiNoteNumber(), 41, t)
	tu.AssertEq(n1.Octave().Octave(), 2, t)
	tu.AssertEq(roundFloat(n1.OffsetInCents(87.32), 0), 0, t)
	tu.AssertEq(roundFloat(n1.OffsetInCents(77.78), 0), -200, t)
	tu.AssertEq(roundFloat(n1.OffsetInCents(92.50), 0), 100, t)
	tu.AssertEq(n1.String(), "F2", t)*/
}
