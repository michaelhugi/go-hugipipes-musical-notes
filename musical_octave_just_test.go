package go_hugipipes_musical_notes

import (
	tu "github.com/informaticon/lib.go.base.test-utils"
	"testing"
)

func TestMOctaveJust(t *testing.T) {
	tj := NewMTemperamentJust(C, A, Octave4, 440)
	o5 := tj.Octave(5)
	notes := o5.AllNotes()
	tu.AssertEq(len(notes), 12, t)
	tu.AssertEq(notes[0].String(), "C5", t)
	tu.AssertEq(notes[1].String(), "C#5", t)
	tu.Assert(!o5.FrequencyBelongsToOctave(510), t)
	tu.Assert(o5.FrequencyBelongsToOctave(520), t)
	tu.Assert(o5.FrequencyBelongsToOctave(1020), t)
	tu.Assert(!o5.FrequencyBelongsToOctave(1030), t)
	tu.AssertEq(notes[0].UpperFrequency(), notes[1].LowerFrequency(), t)
	tu.AssertEq(roundFloat(o5.LowerFrequency(), 2), 512.97, t)
	tu.AssertEq(roundFloat(o5.UpperFrequency(), 2), 1025.94, t)
	tu.AssertEq(o5.Octave(), Octave5, t)
	dis5 := o5.Note(Dis)
	tu.AssertEq(dis5.Note(), Dis, t)
	tu.AssertEq(dis5.String(), "D#5", t)
	tu.AssertEq(o5.String(), "5", t)
}
