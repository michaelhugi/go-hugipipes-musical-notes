package go_hugipipes_musical_notes

import (
	tu "github.com/informaticon/lib.go.base.test-utils"
	"testing"
)

func TestMNoteEqual(t *testing.T) {
	tj := NewMTemperamentJust(C, A, Octave4, 440)
	o3 := tj.Octave(3)
	dis3 := o3.Note(Dis)
	tu.AssertEq(roundFloat(dis3.LowerFrequency(), 2), 152.51, t)
	tu.AssertEq(roundFloat(dis3.ExactFrequency(), 2), 156.98, t)
	tu.AssertEq(roundFloat(dis3.UpperFrequency(), 2), 161.58, t)
	tu.Assert(!dis3.FrequencyBelongsToNote(152.50), t)
	tu.Assert(dis3.FrequencyBelongsToNote(152.51), t)
	tu.Assert(dis3.FrequencyBelongsToNote(161.57), t)
	tu.Assert(!dis3.FrequencyBelongsToNote(161.58), t)
	tu.AssertEq(dis3.MidiNoteNumber(), 51, t)
	tu.AssertEq(dis3.Octave().Octave(), Octave3, t)
	tu.AssertEq(roundFloat(dis3.OffsetInCents(152.51), 0), -50, t)
	tu.AssertEq(roundFloat(dis3.OffsetInCents(156.98), 0), 0, t)
	tu.AssertEq(roundFloat(dis3.OffsetInCents(161.58), 0), 50, t)
	tu.AssertEq(dis3.String(), "D#3", t)
	tu.AssertEq(dis3.Note(), Dis, t)
}
