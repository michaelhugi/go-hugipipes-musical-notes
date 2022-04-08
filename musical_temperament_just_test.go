package go_hugipipes_musical_notes

import (
	tu "github.com/informaticon/lib.go.base.test-utils"
	"testing"
)

func TestNewMTemperamentJustCA440(t *testing.T) {
	tj := NewMTemperamentJust(C, A, Octave4, 440)
	tu.AssertEq(tj.freqC0, 27.5/scaleJust9, t)
	tu.AssertEq(tj.baseNote, C, t)
	tu.AssertEq(roundFloat(AddCents(tj.freqC0, intervalToCents(scaleJust1)), 2), roundFloat(tj.freqCis0, 2), t)
	tu.AssertEq(roundFloat(AddCents(tj.freqC0, intervalToCents(scaleJust2)), 2), roundFloat(tj.freqD0, 2), t)
	tu.AssertEq(roundFloat(AddCents(tj.freqC0, intervalToCents(scaleJust3)), 2), roundFloat(tj.freqDis0, 2), t)
	tu.AssertEq(roundFloat(AddCents(tj.freqC0, intervalToCents(scaleJust4)), 2), roundFloat(tj.freqE0, 2), t)
	tu.AssertEq(roundFloat(AddCents(tj.freqC0, intervalToCents(scaleJust5)), 2), roundFloat(tj.freqF0, 2), t)
	tu.AssertEq(roundFloat(AddCents(tj.freqC0, intervalToCents(scaleJust6)), 2), roundFloat(tj.freqFis0, 2), t)
	tu.AssertEq(roundFloat(AddCents(tj.freqC0, intervalToCents(scaleJust7)), 2), roundFloat(tj.freqG0, 2), t)
	tu.AssertEq(roundFloat(AddCents(tj.freqC0, intervalToCents(scaleJust8)), 2), roundFloat(tj.freqGis0, 2), t)
	tu.AssertEq(roundFloat(AddCents(tj.freqC0, intervalToCents(scaleJust9)), 2), roundFloat(tj.freqA0, 2), t)
	tu.AssertEq(roundFloat(AddCents(tj.freqC0, intervalToCents(scaleJust10)), 2), roundFloat(tj.freqAis0, 2), t)
	tu.AssertEq(roundFloat(AddCents(tj.freqC0, intervalToCents(scaleJust11)), 2), roundFloat(tj.freqB0, 2), t)
	tu.AssertEq(roundFloat(tj.freqA0, 2), 27.50, t)
}

func TestNewMTemperamentJustCisA440(t *testing.T) {
	tj := NewMTemperamentJust(Cis, A, Octave4, 440)
	tu.AssertEq(tj.freqCis0, 27.5/scaleJust8, t)
	tu.AssertEq(tj.baseNote, Cis, t)
	tu.AssertEq(roundFloat(AddCents(tj.freqCis0, intervalToCents(scaleJust1)), 2), roundFloat(tj.freqD0, 2), t)
	tu.AssertEq(roundFloat(AddCents(tj.freqCis0, intervalToCents(scaleJust2)), 2), roundFloat(tj.freqDis0, 2), t)
	tu.AssertEq(roundFloat(AddCents(tj.freqCis0, intervalToCents(scaleJust3)), 2), roundFloat(tj.freqE0, 2), t)
	tu.AssertEq(roundFloat(AddCents(tj.freqCis0, intervalToCents(scaleJust4)), 2), roundFloat(tj.freqF0, 2), t)
	tu.AssertEq(roundFloat(AddCents(tj.freqCis0, intervalToCents(scaleJust5)), 2), roundFloat(tj.freqFis0, 2), t)
	tu.AssertEq(roundFloat(AddCents(tj.freqCis0, intervalToCents(scaleJust6)), 2), roundFloat(tj.freqG0, 2), t)
	tu.AssertEq(roundFloat(AddCents(tj.freqCis0, intervalToCents(scaleJust7)), 2), roundFloat(tj.freqGis0, 2), t)
	tu.AssertEq(roundFloat(AddCents(tj.freqCis0, intervalToCents(scaleJust8)), 2), roundFloat(tj.freqA0, 2), t)
	tu.AssertEq(roundFloat(AddCents(tj.freqCis0, intervalToCents(scaleJust9)), 2), roundFloat(tj.freqAis0, 2), t)
	tu.AssertEq(roundFloat(AddCents(tj.freqCis0, intervalToCents(scaleJust10)), 2), roundFloat(tj.freqB0, 2), t)
	tu.AssertEq(roundFloat(AddCents(tj.freqCis0, intervalToCents(scaleJust11/2)), 2), roundFloat(tj.freqC0, 2), t)
	tu.AssertEq(roundFloat(tj.freqA0, 2), 27.50, t)
}

func TestNewMTemperamentJustDA440(t *testing.T) {
	tj := NewMTemperamentJust(D, A, Octave4, 440)
	tu.AssertEq(tj.freqD0, 27.5/scaleJust7, t)
	tu.AssertEq(tj.baseNote, D, t)
	tu.AssertEq(roundFloat(AddCents(tj.freqD0, intervalToCents(scaleJust1)), 2), roundFloat(tj.freqDis0, 2), t)
	tu.AssertEq(roundFloat(AddCents(tj.freqD0, intervalToCents(scaleJust2)), 2), roundFloat(tj.freqE0, 2), t)
	tu.AssertEq(roundFloat(AddCents(tj.freqD0, intervalToCents(scaleJust3)), 2), roundFloat(tj.freqF0, 2), t)
	tu.AssertEq(roundFloat(AddCents(tj.freqD0, intervalToCents(scaleJust4)), 2), roundFloat(tj.freqFis0, 2), t)
	tu.AssertEq(roundFloat(AddCents(tj.freqD0, intervalToCents(scaleJust5)), 2), roundFloat(tj.freqG0, 2), t)
	tu.AssertEq(roundFloat(AddCents(tj.freqD0, intervalToCents(scaleJust6)), 2), roundFloat(tj.freqGis0, 2), t)
	tu.AssertEq(roundFloat(AddCents(tj.freqD0, intervalToCents(scaleJust7)), 2), roundFloat(tj.freqA0, 2), t)
	tu.AssertEq(roundFloat(AddCents(tj.freqD0, intervalToCents(scaleJust8)), 2), roundFloat(tj.freqAis0, 2), t)
	tu.AssertEq(roundFloat(AddCents(tj.freqD0, intervalToCents(scaleJust9)), 2), roundFloat(tj.freqB0, 2), t)
	tu.AssertEq(roundFloat(AddCents(tj.freqD0, intervalToCents(scaleJust10/2)), 2), roundFloat(tj.freqC0, 2), t)
	tu.AssertEq(roundFloat(AddCents(tj.freqD0, intervalToCents(scaleJust11/2)), 2), roundFloat(tj.freqCis0, 2), t)
}

func TestNewMTemperamentJustDisA440(t *testing.T) {
	tj := NewMTemperamentJust(Dis, A, Octave4, 440)
	tu.AssertEq(tj.freqDis0, 27.5/scaleJust6, t)
	tu.AssertEq(tj.baseNote, Dis, t)
	tu.AssertEq(roundFloat(AddCents(tj.freqDis0, intervalToCents(scaleJust1)), 2), roundFloat(tj.freqE0, 2), t)
	tu.AssertEq(roundFloat(AddCents(tj.freqDis0, intervalToCents(scaleJust2)), 2), roundFloat(tj.freqF0, 2), t)
	tu.AssertEq(roundFloat(AddCents(tj.freqDis0, intervalToCents(scaleJust3)), 2), roundFloat(tj.freqFis0, 2), t)
	tu.AssertEq(roundFloat(AddCents(tj.freqDis0, intervalToCents(scaleJust4)), 2), roundFloat(tj.freqG0, 2), t)
	tu.AssertEq(roundFloat(AddCents(tj.freqDis0, intervalToCents(scaleJust5)), 2), roundFloat(tj.freqGis0, 2), t)
	tu.AssertEq(roundFloat(AddCents(tj.freqDis0, intervalToCents(scaleJust6)), 2), roundFloat(tj.freqA0, 2), t)
	tu.AssertEq(roundFloat(AddCents(tj.freqDis0, intervalToCents(scaleJust7)), 2), roundFloat(tj.freqAis0, 2), t)
	tu.AssertEq(roundFloat(AddCents(tj.freqDis0, intervalToCents(scaleJust8)), 2), roundFloat(tj.freqB0, 2), t)
	tu.AssertEq(roundFloat(AddCents(tj.freqDis0, intervalToCents(scaleJust9/2)), 2), roundFloat(tj.freqC0, 2), t)
	tu.AssertEq(roundFloat(AddCents(tj.freqDis0, intervalToCents(scaleJust10/2)), 2), roundFloat(tj.freqCis0, 2), t)
	tu.AssertEq(roundFloat(AddCents(tj.freqDis0, intervalToCents(scaleJust11/2)), 2), roundFloat(tj.freqD0, 2), t)
}

func TestNewMTemperamentJustBA440(t *testing.T) {
	tj := NewMTemperamentJust(B, A, Octave4, 440)
	tu.AssertEq(roundFloat(tj.freqB0*scaleJust10/2, 2), roundFloat(27.5, 2), t)
	tu.AssertEq(tj.baseNote, B, t)
	tu.AssertEq(roundFloat(AddCents(tj.freqB0, intervalToCents(scaleJust1/2)), 2), roundFloat(tj.freqC0, 2), t)
	tu.AssertEq(roundFloat(AddCents(tj.freqB0, intervalToCents(scaleJust2/2)), 2), roundFloat(tj.freqCis0, 2), t)
	tu.AssertEq(roundFloat(AddCents(tj.freqB0, intervalToCents(scaleJust3/2)), 2), roundFloat(tj.freqD0, 2), t)
	tu.AssertEq(roundFloat(AddCents(tj.freqB0, intervalToCents(scaleJust4/2)), 2), roundFloat(tj.freqDis0, 2), t)
	tu.AssertEq(roundFloat(AddCents(tj.freqB0, intervalToCents(scaleJust5/2)), 2), roundFloat(tj.freqE0, 2), t)
	tu.AssertEq(roundFloat(AddCents(tj.freqB0, intervalToCents(scaleJust6/2)), 2), roundFloat(tj.freqF0, 2), t)
	tu.AssertEq(roundFloat(AddCents(tj.freqB0, intervalToCents(scaleJust7/2)), 2), roundFloat(tj.freqFis0, 2), t)
	tu.AssertEq(roundFloat(AddCents(tj.freqB0, intervalToCents(scaleJust8/2)), 2), roundFloat(tj.freqG0, 2), t)
	tu.AssertEq(roundFloat(AddCents(tj.freqB0, intervalToCents(scaleJust9/2)), 2), roundFloat(tj.freqGis0, 2), t)
	tu.AssertEq(roundFloat(AddCents(tj.freqB0, intervalToCents(scaleJust10/2)), 2), roundFloat(tj.freqA0, 2), t)
	tu.AssertEq(roundFloat(AddCents(tj.freqB0, intervalToCents(scaleJust11/2)), 2), roundFloat(tj.freqAis0, 2), t)
}

func TestNewMTemperamentJustBA220(t *testing.T) {
	tj := NewMTemperamentJust(B, A, Octave4, 220)
	tu.AssertEq(roundFloat(tj.freqB0*scaleJust10/2, 2), roundFloat(27.5/2, 2), t)
	tu.AssertEq(tj.baseNote, B, t)
	tu.AssertEq(roundFloat(AddCents(tj.freqB0, intervalToCents(scaleJust1/2)), 2), roundFloat(tj.freqC0, 2), t)
	tu.AssertEq(roundFloat(AddCents(tj.freqB0, intervalToCents(scaleJust2/2)), 2), roundFloat(tj.freqCis0, 2), t)
	tu.AssertEq(roundFloat(AddCents(tj.freqB0, intervalToCents(scaleJust3/2)), 2), roundFloat(tj.freqD0, 2), t)
	tu.AssertEq(roundFloat(AddCents(tj.freqB0, intervalToCents(scaleJust4/2)), 2), roundFloat(tj.freqDis0, 2), t)
	tu.AssertEq(roundFloat(AddCents(tj.freqB0, intervalToCents(scaleJust5/2)), 2), roundFloat(tj.freqE0, 2), t)
	tu.AssertEq(roundFloat(AddCents(tj.freqB0, intervalToCents(scaleJust6/2)), 2), roundFloat(tj.freqF0, 2), t)
	tu.AssertEq(roundFloat(AddCents(tj.freqB0, intervalToCents(scaleJust7/2)), 2), roundFloat(tj.freqFis0, 2), t)
	tu.AssertEq(roundFloat(AddCents(tj.freqB0, intervalToCents(scaleJust8/2)), 2), roundFloat(tj.freqG0, 2), t)
	tu.AssertEq(roundFloat(AddCents(tj.freqB0, intervalToCents(scaleJust9/2)), 2), roundFloat(tj.freqGis0, 2), t)
	tu.AssertEq(roundFloat(AddCents(tj.freqB0, intervalToCents(scaleJust10/2)), 2), roundFloat(tj.freqA0, 2), t)
	tu.AssertEq(roundFloat(AddCents(tj.freqB0, intervalToCents(scaleJust11/2)), 2), roundFloat(tj.freqAis0, 2), t)
}

func TestNewMTemperamentJust(t *testing.T) {
	tj := NewMTemperamentJust(B, A, Octave4, 440)
	tu.AssertEq(roundFloat(tj.freqB0*scaleJust10/2, 2), roundFloat(27.5, 2), t)
	oct := tj.AllOctaves()
	tu.AssertEq(oct[0].Octave(), OctaveMinus1, t)
	tu.AssertEq(oct[1].Octave(), Octave0, t)
	tu.AssertEq(oct[2].Octave(), Octave1, t)
	tu.AssertEq(oct[3].Octave(), Octave2, t)
	tu.AssertEq(oct[4].Octave(), Octave3, t)
	tu.AssertEq(oct[5].Octave(), Octave4, t)
	tu.Assert(oct[5].FrequencyBelongsToOctave(440.0), t)
	tu.AssertEq(oct[6].Octave(), Octave5, t)
	tu.AssertEq(oct[7].Octave(), Octave6, t)

	o3, err := tj.OctaveFromFrequency(220)
	tu.AssertNErr(err)
	tu.AssertEq(o3.Octave(), Octave3, t)

	o4 := tj.Octave(Octave4)
	tu.AssertEq(roundFloat(o4.LowerFrequency(), 2), 253.32, t)
	tu.AssertEq(roundFloat(o4.Note(C).ExactFrequency(), 2), 260.74, t)
	tu.AssertEq(roundFloat(o4.Note(Cis).ExactFrequency(), 2), 276.25, t)
	tu.AssertEq(roundFloat(o4.Note(D).ExactFrequency(), 2), 292.67, t)
	tu.AssertEq(roundFloat(o4.Note(Dis).ExactFrequency(), 2), 310.07, t)
	tu.AssertEq(roundFloat(o4.Note(E).ExactFrequency(), 2), 328.51, t)
	tu.AssertEq(roundFloat(o4.Note(F).ExactFrequency(), 2), 348.05, t)
	tu.AssertEq(roundFloat(o4.Note(Fis).ExactFrequency(), 2), 368.74, t)
	tu.AssertEq(roundFloat(o4.Note(G).ExactFrequency(), 2), 390.67, t)
	tu.AssertEq(roundFloat(o4.Note(Gis).ExactFrequency(), 2), 413.9, t)
	tu.AssertEq(roundFloat(o4.Note(A).ExactFrequency(), 2), 438.51, t)
	tu.AssertEq(roundFloat(o4.Note(Ais).ExactFrequency(), 2), 464.59, t)
	tu.AssertEq(roundFloat(o4.Note(B).ExactFrequency(), 2), 492.21, t)
	tu.AssertEq(roundFloat(o4.UpperFrequency(), 2), 506.64, t)

	fis3, err := tj.NoteFromFrequency(182)
	tu.AssertNErr(err)
	tu.AssertEq(fis3.Octave().Octave(), Octave3, t)
	tu.AssertEq(roundFloat(fis3.ExactFrequency(), 2), 184.37, t)
}

func TestMTemperamentJustAA440(t *testing.T) {
	tj := NewMTemperamentJust(A, A, Octave4, 440)
	oct4 := tj.Octave(Octave4)
	tu.AssertEq(roundFloat(oct4.Note(A).ExactFrequency(), 2), roundFloat(440, 2), t)
}

func TestMTemperamentJustCA440(t *testing.T) {
	tj := NewMTemperamentJust(C, A, Octave4, 440)
	oct4 := tj.Octave(Octave4)
	tu.AssertEq(roundFloat(oct4.Note(C).ExactFrequency(), 2), roundFloat(264, 2), t)
	tu.AssertEq(roundFloat(oct4.Note(Cis).ExactFrequency(), 2), roundFloat(281.6, 2), t)
	tu.AssertEq(roundFloat(oct4.Note(D).ExactFrequency(), 2), roundFloat(297, 2), t)
	tu.AssertEq(roundFloat(oct4.Note(Dis).ExactFrequency(), 2), roundFloat(316.8, 2), t)
	tu.AssertEq(roundFloat(oct4.Note(E).ExactFrequency(), 2), roundFloat(330.0, 2), t)
	tu.AssertEq(roundFloat(oct4.Note(F).ExactFrequency(), 2), roundFloat(352.0, 2), t)
	tu.AssertEq(roundFloat(oct4.Note(Fis).ExactFrequency(), 2), roundFloat(371.25, 2), t)
	tu.AssertEq(roundFloat(oct4.Note(G).ExactFrequency(), 2), roundFloat(396.0, 2), t)
	tu.AssertEq(roundFloat(oct4.Note(Gis).ExactFrequency(), 2), roundFloat(422.4, 2), t)
	tu.AssertEq(roundFloat(oct4.Note(A).ExactFrequency(), 2), roundFloat(440, 2), t)
	tu.AssertEq(roundFloat(oct4.Note(B).ExactFrequency(), 2), roundFloat(475.2, 2), t)
	tu.AssertEq(roundFloat(oct4.Note(C).ExactFrequency(), 2), roundFloat(495.0, 2), t)
}
