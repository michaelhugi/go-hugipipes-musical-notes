package go_hugipipes_musical_notes

import (
	"fmt"
	"math"
)

//MOctaveJust represents a musical octave for a MTemperamentJust with a specific concert pitch
type MOctaveJust struct {
	lowerFrequency float64
	upperFrequency float64
	octave         MGOctave
	allNotes       []MNote
}

//newMOctaveJust is the constructor for MOctaveJust used by MTemperamentJust
func newMOctaveJust(octave MGOctave, freqC0 float64, lowerFrequency float64) *MOctaveJust {
	freqC := freqC0 * math.Pow(2, float64(octave))
	midiNoteC := uint8(12 + (12 * octave))
	self := &MOctaveJust{
		lowerFrequency: lowerFrequency,
		upperFrequency: AddCents(freqC, 1150),
		octave:         octave,
	}
	self.allNotes = calcAllNotesJust(self, freqC, midiNoteC)
	return self
}

//calcAllNotesJust calculates all notes for the Just temperament for the octave
func calcAllNotesJust(octave *MOctaveJust, freqC float64, midiNoteC uint8) []MNote {
	notes := make([]MNote, 12)
	notes[0] = newMNoteJust(freqC, C, octave, midiNoteC, "C", AddCents(freqC, -50))
	notes[1] = newMNoteJust(AddCents(freqC, 100), Cis, octave, midiNoteC+1, "C#", notes[0].UpperFrequency())
	notes[2] = newMNoteJust(AddCents(freqC, 200), D, octave, midiNoteC+2, "D", notes[1].UpperFrequency())
	notes[3] = newMNoteJust(AddCents(freqC, 300), Dis, octave, midiNoteC+3, "D#", notes[2].UpperFrequency())
	notes[4] = newMNoteJust(AddCents(freqC, 400), E, octave, midiNoteC+4, "E", notes[3].UpperFrequency())
	notes[5] = newMNoteJust(AddCents(freqC, 500), F, octave, midiNoteC+5, "F", notes[4].UpperFrequency())
	notes[6] = newMNoteJust(AddCents(freqC, 600), Fis, octave, midiNoteC+6, "F#", notes[5].UpperFrequency())
	notes[7] = newMNoteJust(AddCents(freqC, 700), G, octave, midiNoteC+7, "G", notes[6].UpperFrequency())
	notes[8] = newMNoteJust(AddCents(freqC, 800), Gis, octave, midiNoteC+8, "G#", notes[7].UpperFrequency())
	notes[9] = newMNoteJust(AddCents(freqC, 900), A, octave, midiNoteC+9, "A", notes[8].UpperFrequency())
	notes[10] = newMNoteJust(AddCents(freqC, 1000), Ais, octave, midiNoteC+10, "A#", notes[9].UpperFrequency())
	notes[11] = newMNoteJust(AddCents(freqC, 1100), B, octave, midiNoteC+11, "B", notes[10].UpperFrequency())
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

func (s MOctaveJust) Note(n MGNote) MNote {
	for _, nn := range s.allNotes {
		if nn.Note() == n {
			return nn
		}
	}
	panic(fmt.Sprintf("Note %d not found", n))
}

func (s MOctaveJust) String() string {
	return fmt.Sprintf("%d", s.octave)
}
