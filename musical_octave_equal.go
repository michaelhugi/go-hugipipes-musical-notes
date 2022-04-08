package go_hugipipes_musical_notes

import (
	"fmt"
	"math"
)

//MOctaveEqual represents a musical octave for a MTemperamentEqual with a specific concert pitch
type MOctaveEqual struct {
	lowerFrequency float64
	upperFrequency float64
	octave         MGOctave
	allNotes       []MNote
}

//newMOctaveEqual is the constructor for MOctaveEqual used by MTemperamentEqual
func newMOctaveEqual(octave MGOctave, freqC0 float64, lowerFrequency float64) *MOctaveEqual {
	freqC := freqC0 * math.Pow(2, float64(octave))
	midiNoteC := uint8(12 + (12 * octave))
	self := &MOctaveEqual{
		lowerFrequency: lowerFrequency,
		upperFrequency: AddCents(freqC, 1150),
		octave:         octave,
	}
	self.allNotes = calcAllNotesEqual(self, freqC, midiNoteC)
	return self
}

//calcAllNotesEqual calculates all notes for the equal temperament for the octave
func calcAllNotesEqual(octave *MOctaveEqual, freqC float64, midiNoteC uint8) []MNote {
	notes := make([]MNote, 12)
	notes[0] = newMNoteEqual(freqC, C, octave, midiNoteC, "C", AddCents(freqC, -50))
	notes[1] = newMNoteEqual(AddCents(freqC, 100), Cis, octave, midiNoteC+1, "C#", notes[0].UpperFrequency())
	notes[2] = newMNoteEqual(AddCents(freqC, 200), D, octave, midiNoteC+2, "D", notes[1].UpperFrequency())
	notes[3] = newMNoteEqual(AddCents(freqC, 300), Dis, octave, midiNoteC+3, "D#", notes[2].UpperFrequency())
	notes[4] = newMNoteEqual(AddCents(freqC, 400), E, octave, midiNoteC+4, "E", notes[3].UpperFrequency())
	notes[5] = newMNoteEqual(AddCents(freqC, 500), F, octave, midiNoteC+5, "F", notes[4].UpperFrequency())
	notes[6] = newMNoteEqual(AddCents(freqC, 600), Fis, octave, midiNoteC+6, "F#", notes[5].UpperFrequency())
	notes[7] = newMNoteEqual(AddCents(freqC, 700), G, octave, midiNoteC+7, "G", notes[6].UpperFrequency())
	notes[8] = newMNoteEqual(AddCents(freqC, 800), Gis, octave, midiNoteC+8, "G#", notes[7].UpperFrequency())
	notes[9] = newMNoteEqual(AddCents(freqC, 900), A, octave, midiNoteC+9, "A", notes[8].UpperFrequency())
	notes[10] = newMNoteEqual(AddCents(freqC, 1000), Ais, octave, midiNoteC+10, "A#", notes[9].UpperFrequency())
	notes[11] = newMNoteEqual(AddCents(freqC, 1100), B, octave, midiNoteC+11, "B", notes[10].UpperFrequency())
	return notes
}

//AllNotes returns all notes that belong to the octave
func (s MOctaveEqual) AllNotes() []MNote {
	return s.allNotes
}

//FrequencyBelongsToOctave returns, if a tested frequency belongs to the octave or not
func (s MOctaveEqual) FrequencyBelongsToOctave(frequency float64) bool {
	return s.lowerFrequency <= frequency && s.upperFrequency > frequency
}

//LowerFrequency returns the lowest frequency that belongs to the octave
func (s MOctaveEqual) LowerFrequency() float64 {
	return s.lowerFrequency
}

//UpperFrequency returns the highest frequency that belongs to the octave
func (s MOctaveEqual) UpperFrequency() float64 {
	return s.upperFrequency
}

//Octave returns the readable octave name of the note
func (s MOctaveEqual) Octave() MGOctave {
	return s.octave
}

//Note returns the generic Note enum for the musical note
func (s MOctaveEqual) Note(n MGNote) MNote {
	for _, nn := range s.allNotes {
		if nn.Note() == n {
			return nn
		}
	}
	panic(fmt.Sprintf("Note %d not found", n))
}

//MOctaveEqual is a human-readable representation of the note including the octave
func (s MOctaveEqual) String() string {
	return fmt.Sprintf("%d", s.octave)
}
