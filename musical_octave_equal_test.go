package go_hugipipes_musical_notes

import (
	tu "github.com/informaticon/lib.go.base.test-utils"
	"testing"
)

func TestMOctaveEqualInterface(t *testing.T) {
	o := MOctaveEqual{}
	checkMOctaveInterface(o)
}

func TestOctaveEqual440(t *testing.T) {
	te := NewMTemperamentEqual(440)
	o4 := te.Octave(Octave4)
	tu.AssertEq(roundFloat(o4.LowerFrequency(), 2), 254.18, t)
	tu.AssertEq(roundFloat(o4.UpperFrequency(), 2), 508.36, t)

	n4 := o4.AllNotes()
	tu.AssertEq(len(n4), 12, t)
	tu.AssertEq(roundFloat(n4[0].ExactFrequency(), 2), 261.63, t)
	tu.AssertEq(roundFloat(n4[1].ExactFrequency(), 2), 277.18, t)
	tu.AssertEq(roundFloat(n4[2].ExactFrequency(), 2), 293.66, t)
	tu.AssertEq(roundFloat(n4[3].ExactFrequency(), 2), 311.13, t)
	tu.AssertEq(roundFloat(n4[4].ExactFrequency(), 2), 329.63, t)
	tu.AssertEq(roundFloat(n4[5].ExactFrequency(), 2), 349.23, t)
	tu.AssertEq(roundFloat(n4[6].ExactFrequency(), 2), 369.99, t)
	tu.AssertEq(roundFloat(n4[7].ExactFrequency(), 2), 392.00, t)
	tu.AssertEq(roundFloat(n4[8].ExactFrequency(), 2), 415.30, t)
	tu.AssertEq(roundFloat(n4[9].ExactFrequency(), 2), 440.00, t)
	tu.AssertEq(roundFloat(n4[10].ExactFrequency(), 2), 466.16, t)
	tu.AssertEq(roundFloat(n4[11].ExactFrequency(), 2), 493.88, t)
	tu.AssertEq(n4[2].String(), "D4", t)
	tu.AssertEq(n4[3].String(), "D#4", t)
	tu.AssertEq(n4[4].MidiNoteNumber(), 64, t)

	tu.Assert(o4.FrequencyBelongsToOctave(349.23), t)
	tu.Assert(!o4.FrequencyBelongsToOctave(510), t)
	tu.Assert(!o4.FrequencyBelongsToOctave(250), t)

	tu.AssertEq(o4.Octave(), Octave4, t)

	tu.AssertEq(o4.Note(Fis).MidiNoteNumber(), 66, t)
	tu.AssertEq(roundFloat(o4.Note(G).ExactFrequency(), 2), 392.00, t)
	tu.AssertEq(roundFloat(o4.Note(C).LowerFrequency(), 2), 254.18, t)
	tu.AssertEq(o4.String(), "4", t)

	o5 := te.Octave(Octave5)
	tu.AssertEq(roundFloat(o5.LowerFrequency(), 2), 508.36, t)
	tu.AssertEq(roundFloat(o5.UpperFrequency(), 2), 1016.71, t)
	n5 := o5.AllNotes()
	tu.AssertEq(len(n5), 12, t)
	tu.AssertEq(roundFloat(n5[0].ExactFrequency(), 2), 523.25, t)
	tu.AssertEq(roundFloat(n5[1].ExactFrequency(), 2), 554.37, t)
	tu.AssertEq(roundFloat(n5[2].ExactFrequency(), 2), 587.33, t)
	tu.AssertEq(roundFloat(n5[3].ExactFrequency(), 2), 622.25, t)
	tu.AssertEq(roundFloat(n5[4].ExactFrequency(), 2), 659.26, t)
	tu.AssertEq(roundFloat(n5[5].ExactFrequency(), 2), 698.46, t)
	tu.AssertEq(roundFloat(n5[6].ExactFrequency(), 2), 739.99, t)
	tu.AssertEq(roundFloat(n5[7].ExactFrequency(), 2), 783.99, t)
	tu.AssertEq(roundFloat(n5[8].ExactFrequency(), 2), 830.61, t)
	tu.AssertEq(roundFloat(n5[9].ExactFrequency(), 2), 880, t)
	tu.AssertEq(roundFloat(n5[10].ExactFrequency(), 2), 932.33, t)
	tu.AssertEq(roundFloat(n5[11].ExactFrequency(), 2), 987.77, t)
	tu.AssertEq(n5[2].String(), "D5", t)
	tu.AssertEq(n5[3].String(), "D#5", t)
	tu.AssertEq(n5[4].MidiNoteNumber(), 76, t)

	tu.Assert(o5.FrequencyBelongsToOctave(698.46), t)
	tu.Assert(!o5.FrequencyBelongsToOctave(1020), t)
	tu.Assert(!o5.FrequencyBelongsToOctave(500), t)

	tu.AssertEq(o5.Octave(), Octave5, t)

	tu.AssertEq(o5.Note(Fis).MidiNoteNumber(), 78, t)
	tu.AssertEq(roundFloat(o5.Note(G).ExactFrequency(), 2), 783.99, t)
	tu.AssertEq(roundFloat(o5.Note(C).LowerFrequency(), 2), 508.36, t)
	tu.AssertEq(o5.String(), "5", t)

}

func TestOctaveEqual220(t *testing.T) {
	te := NewMTemperamentEqual(220)
	o4 := te.Octave(Octave4)
	tu.AssertEq(roundFloat(o4.LowerFrequency(), 2), 127.09, t)
	tu.AssertEq(roundFloat(o4.UpperFrequency(), 2), 254.18, t)

	n4 := o4.AllNotes()
	tu.AssertEq(len(n4), 12, t)
	tu.AssertEq(roundFloat(n4[0].ExactFrequency(), 2), 130.81, t)
	tu.AssertEq(roundFloat(n4[1].ExactFrequency(), 2), 138.59, t)
	tu.AssertEq(roundFloat(n4[2].ExactFrequency(), 2), 146.83, t)
	tu.AssertEq(roundFloat(n4[3].ExactFrequency(), 2), 155.56, t)
	tu.AssertEq(roundFloat(n4[4].ExactFrequency(), 2), 164.81, t)
	tu.AssertEq(roundFloat(n4[5].ExactFrequency(), 2), 174.61, t)
	tu.AssertEq(roundFloat(n4[6].ExactFrequency(), 2), 185, t)
	tu.AssertEq(roundFloat(n4[7].ExactFrequency(), 2), 196, t)
	tu.AssertEq(roundFloat(n4[8].ExactFrequency(), 2), 207.65, t)
	tu.AssertEq(roundFloat(n4[9].ExactFrequency(), 2), 220, t)
	tu.AssertEq(roundFloat(n4[10].ExactFrequency(), 2), 233.08, t)
	tu.AssertEq(roundFloat(n4[11].ExactFrequency(), 2), 246.94, t)
	tu.AssertEq(n4[2].String(), "D4", t)
	tu.AssertEq(n4[3].String(), "D#4", t)
	tu.AssertEq(n4[4].MidiNoteNumber(), 64, t)

	tu.Assert(o4.FrequencyBelongsToOctave(174.62), t)
	tu.Assert(!o4.FrequencyBelongsToOctave(255), t)
	tu.Assert(!o4.FrequencyBelongsToOctave(125), t)

	tu.AssertEq(o4.Octave(), Octave4, t)

	tu.AssertEq(o4.Note(Fis).MidiNoteNumber(), 66, t)
	tu.AssertEq(roundFloat(o4.Note(G).ExactFrequency(), 2), 196, t)
	tu.AssertEq(roundFloat(o4.Note(C).LowerFrequency(), 2), 127.09, t)
	tu.AssertEq(o4.String(), "4", t)

	o5 := te.Octave(Octave5)
	tu.AssertEq(roundFloat(o5.LowerFrequency(), 2), 254.18, t)
	tu.AssertEq(roundFloat(o5.UpperFrequency(), 2), 508.36, t)
	n5 := o5.AllNotes()
	tu.AssertEq(len(n5), 12, t)
	tu.AssertEq(roundFloat(n5[0].ExactFrequency(), 2), 261.63, t)
	tu.AssertEq(roundFloat(n5[1].ExactFrequency(), 2), 277.18, t)
	tu.AssertEq(roundFloat(n5[2].ExactFrequency(), 2), 293.66, t)
	tu.AssertEq(roundFloat(n5[3].ExactFrequency(), 2), 311.13, t)
	tu.AssertEq(roundFloat(n5[4].ExactFrequency(), 2), 329.63, t)
	tu.AssertEq(roundFloat(n5[5].ExactFrequency(), 2), 349.23, t)
	tu.AssertEq(roundFloat(n5[6].ExactFrequency(), 2), 369.99, t)
	tu.AssertEq(roundFloat(n5[7].ExactFrequency(), 2), 392, t)
	tu.AssertEq(roundFloat(n5[8].ExactFrequency(), 2), 415.3, t)
	tu.AssertEq(roundFloat(n5[9].ExactFrequency(), 2), 440, t)
	tu.AssertEq(roundFloat(n5[10].ExactFrequency(), 2), 466.16, t)
	tu.AssertEq(roundFloat(n5[11].ExactFrequency(), 2), 493.88, t)
	tu.AssertEq(n5[2].String(), "D5", t)
	tu.AssertEq(n5[3].String(), "D#5", t)
	tu.AssertEq(n5[4].MidiNoteNumber(), 76, t)

	tu.Assert(o5.FrequencyBelongsToOctave(349.23), t)
	tu.Assert(!o5.FrequencyBelongsToOctave(510), t)
	tu.Assert(!o5.FrequencyBelongsToOctave(250), t)

	tu.AssertEq(o5.Octave(), Octave5, t)

	tu.AssertEq(o5.Note(Fis).MidiNoteNumber(), 78, t)
	tu.AssertEq(roundFloat(o5.Note(G).ExactFrequency(), 2), 392, t)
	tu.AssertEq(roundFloat(o5.Note(C).LowerFrequency(), 2), 254.18, t)
	tu.AssertEq(o5.String(), "5", t)

}

func TestOctaveEqual435(t *testing.T) {
	te := NewMTemperamentEqual(435)
	o4 := te.Octave(Octave4)
	tu.AssertEq(roundFloat(o4.LowerFrequency(), 2), 251.29, t)
	tu.AssertEq(roundFloat(o4.UpperFrequency(), 2), 502.58, t)

	n4 := o4.AllNotes()
	tu.AssertEq(len(n4), 12, t)
	tu.AssertEq(roundFloat(n4[0].ExactFrequency(), 2), 258.65, t)
	tu.AssertEq(roundFloat(n4[1].ExactFrequency(), 2), 274.03, t)
	tu.AssertEq(roundFloat(n4[2].ExactFrequency(), 2), 290.33, t)
	tu.AssertEq(roundFloat(n4[3].ExactFrequency(), 2), 307.59, t)
	tu.AssertEq(roundFloat(n4[4].ExactFrequency(), 2), 325.88, t)
	tu.AssertEq(roundFloat(n4[5].ExactFrequency(), 2), 345.26, t)
	tu.AssertEq(roundFloat(n4[6].ExactFrequency(), 2), 365.79, t)
	tu.AssertEq(roundFloat(n4[7].ExactFrequency(), 2), 387.54, t)
	tu.AssertEq(roundFloat(n4[8].ExactFrequency(), 2), 410.59, t)
	tu.AssertEq(roundFloat(n4[9].ExactFrequency(), 2), 435, t)
	tu.AssertEq(roundFloat(n4[10].ExactFrequency(), 2), 460.87, t)
	tu.AssertEq(roundFloat(n4[11].ExactFrequency(), 2), 488.27, t)
	tu.AssertEq(n4[2].String(), "D4", t)
	tu.AssertEq(n4[3].String(), "D#4", t)
	tu.AssertEq(n4[4].MidiNoteNumber(), 64, t)

	tu.Assert(o4.FrequencyBelongsToOctave(349.23), t)
	tu.Assert(!o4.FrequencyBelongsToOctave(510), t)
	tu.Assert(!o4.FrequencyBelongsToOctave(250), t)

	tu.AssertEq(o4.Octave(), Octave4, t)

	tu.AssertEq(o4.Note(Fis).MidiNoteNumber(), 66, t)
	tu.AssertEq(roundFloat(o4.Note(G).ExactFrequency(), 2), 387.54, t)
	tu.AssertEq(roundFloat(o4.Note(C).LowerFrequency(), 2), 251.29, t)
	tu.AssertEq(o4.String(), "4", t)

	o5 := te.Octave(Octave5)
	tu.AssertEq(roundFloat(o5.LowerFrequency(), 2), 502.58, t)
	tu.AssertEq(roundFloat(o5.UpperFrequency(), 2), 1005.16, t)
	n5 := o5.AllNotes()
	tu.AssertEq(len(n5), 12, t)
	tu.AssertEq(roundFloat(n5[0].ExactFrequency(), 2), 517.31, t)
	tu.AssertEq(roundFloat(n5[1].ExactFrequency(), 2), 548.07, t)
	tu.AssertEq(roundFloat(n5[2].ExactFrequency(), 2), 580.66, t)
	tu.AssertEq(roundFloat(n5[3].ExactFrequency(), 2), 615.18, t)
	tu.AssertEq(roundFloat(n5[4].ExactFrequency(), 2), 651.76, t)
	tu.AssertEq(roundFloat(n5[5].ExactFrequency(), 2), 690.52, t)
	tu.AssertEq(roundFloat(n5[6].ExactFrequency(), 2), 731.58, t)
	tu.AssertEq(roundFloat(n5[7].ExactFrequency(), 2), 775.08, t)
	tu.AssertEq(roundFloat(n5[8].ExactFrequency(), 2), 821.17, t)
	tu.AssertEq(roundFloat(n5[9].ExactFrequency(), 2), 870, t)
	tu.AssertEq(roundFloat(n5[10].ExactFrequency(), 2), 921.73, t)
	tu.AssertEq(roundFloat(n5[11].ExactFrequency(), 2), 976.54, t)
	tu.AssertEq(n5[2].String(), "D5", t)
	tu.AssertEq(n5[3].String(), "D#5", t)
	tu.AssertEq(n5[4].MidiNoteNumber(), 76, t)

	tu.Assert(o5.FrequencyBelongsToOctave(698.46), t)
	tu.Assert(!o5.FrequencyBelongsToOctave(1020), t)
	tu.Assert(!o5.FrequencyBelongsToOctave(500), t)

	tu.AssertEq(o5.Octave(), Octave5, t)

	tu.AssertEq(o5.Note(Fis).MidiNoteNumber(), 78, t)
	tu.AssertEq(roundFloat(o5.Note(G).ExactFrequency(), 2), 775.08, t)
	tu.AssertEq(roundFloat(o5.Note(C).LowerFrequency(), 2), 502.58, t)
	tu.AssertEq(o5.String(), "5", t)

}
