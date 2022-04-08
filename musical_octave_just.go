package go_hugipipes_musical_notes

import (
	"fmt"
	"math"
)

//MOctaveJust represents a musical octave for a MTemperamentJust with a specific concert pitch
type MOctaveJust struct {
	lowerFrequency float64
	upperFrequency float64

	octave   MGOctave
	allNotes []MNote
}

//newMOctaveJust is the constructor for MOctaveJust used by MTemperamentJust
func newMOctaveJust(octave MGOctave, freqC0 float64, freqCis0 float64, freqD0 float64, freqDis0 float64, freqE0 float64, freqF0 float64, freqFis0 float64, freqG0 float64, freqGis0 float64, freqA0 float64, freqAis0 float64, freqB0 float64, lowerFrequency float64) *MOctaveJust {
	octaveFactor := math.Pow(2, float64(octave))
	freqC := freqC0 * octaveFactor
	freqCis := freqCis0 * octaveFactor
	freqD := freqD0 * octaveFactor
	freqDis := freqDis0 * octaveFactor
	freqE := freqE0 * octaveFactor
	freqF := freqF0 * octaveFactor
	freqFis := freqFis0 * octaveFactor
	freqG := freqG0 * octaveFactor
	freqGis := freqGis0 * octaveFactor
	freqA := freqA0 * octaveFactor
	freqAis := freqAis0 * octaveFactor
	freqB := freqB0 * octaveFactor
	midiNoteC := uint8(12 + (12 * octave))
	self := &MOctaveJust{
		lowerFrequency: lowerFrequency,
		upperFrequency: AddCents(freqC, 1150),
		octave:         octave,
	}
	self.allNotes = calcAllNotesJust(self, freqC, freqCis, freqD, freqDis, freqE, freqF, freqFis, freqG, freqGis, freqA, freqAis, freqB, midiNoteC)
	return self
}

//calcAllNotesJust calculates all notes for the Just temperament for the octave
func calcAllNotesJust(octave *MOctaveJust, freqC float64, freqCis float64, freqD float64, freqDis float64, freqE float64, freqF float64, freqFis float64, freqG float64, freqGis float64, freqA float64, freqAis float64, freqB float64, midiNoteC uint8) []MNote {
	notes := make([]MNote, 12)
	notes[0] = newMNoteJust(freqC, C, octave, midiNoteC, "C", AddCents(freqC, -50))
	notes[1] = newMNoteJust(freqCis, Cis, octave, midiNoteC+1, "C#", notes[0].UpperFrequency())
	notes[2] = newMNoteJust(freqD, D, octave, midiNoteC+2, "D", notes[1].UpperFrequency())
	notes[3] = newMNoteJust(freqDis, Dis, octave, midiNoteC+3, "D#", notes[2].UpperFrequency())
	notes[4] = newMNoteJust(freqE, E, octave, midiNoteC+4, "E", notes[3].UpperFrequency())
	notes[5] = newMNoteJust(freqF, F, octave, midiNoteC+5, "F", notes[4].UpperFrequency())
	notes[6] = newMNoteJust(freqFis, Fis, octave, midiNoteC+6, "F#", notes[5].UpperFrequency())
	notes[7] = newMNoteJust(freqG, G, octave, midiNoteC+7, "G", notes[6].UpperFrequency())
	notes[8] = newMNoteJust(freqGis, Gis, octave, midiNoteC+8, "G#", notes[7].UpperFrequency())
	notes[9] = newMNoteJust(freqA, A, octave, midiNoteC+9, "A", notes[8].UpperFrequency())
	notes[10] = newMNoteJust(freqAis, Ais, octave, midiNoteC+10, "A#", notes[9].UpperFrequency())
	notes[11] = newMNoteJust(freqB, B, octave, midiNoteC+11, "B", notes[10].UpperFrequency())
	return notes
}

//AllNotes returns all notes that belong to the octave
func (s MOctaveJust) AllNotes() []MNote {
	return s.allNotes
}

//FrequencyBelongsToOctave returns, if a tested frequency belongs to the octave or not
func (s MOctaveJust) FrequencyBelongsToOctave(frequency float64) bool {
	return s.lowerFrequency <= frequency && s.upperFrequency > frequency
}

//LowerFrequency returns the lowest frequency that belongs to the octave
func (s MOctaveJust) LowerFrequency() float64 {
	return s.lowerFrequency
}

//UpperFrequency returns the highest frequency that belongs to the octave
func (s MOctaveJust) UpperFrequency() float64 {
	return s.upperFrequency
}

//Octave returns the readable octave name of the note
func (s MOctaveJust) Octave() MGOctave {
	return s.octave
}

//Note returns the generic Note enum for the musical note
func (s MOctaveJust) Note(n MGNote) MNote {
	for _, nn := range s.allNotes {
		if nn.Note() == n {
			return nn
		}
	}
	panic(fmt.Sprintf("Note %d not found", n))
}

//MOctaveEqual is a human-readable representation of the note including the octave
func (s MOctaveJust) String() string {
	return fmt.Sprintf("%d", s.octave)
}
