package go_hugipipes_musical_notes

import (
	"fmt"
)

type MNoteJust struct {
	octave         *MOctaveJust
	exactFrequency float64
	lowerFrequency float64
	upperFrequency float64
	midiNoteNumber uint8
	noteName       string
	note           MGNote
}

func newMNoteJust(exactFrequency float64, note MGNote, octave *MOctaveJust, midiNoteNumber uint8, noteName string, lowerFrequency float64) *MNoteJust {

	return &MNoteJust{
		octave:         octave,
		exactFrequency: exactFrequency,
		lowerFrequency: lowerFrequency,
		upperFrequency: AddCents(exactFrequency, 50),
		midiNoteNumber: midiNoteNumber,
		noteName:       noteName,
		note:           note,
	}
}

func (s MNoteJust) ExactFrequency() float64 {
	return s.exactFrequency
}
func (s MNoteJust) FrequencyBelongsToNote(frequency float64) bool {
	return frequency >= s.lowerFrequency && frequency < s.upperFrequency
}
func (s MNoteJust) LowerFrequency() float64 {
	return s.lowerFrequency
}
func (s MNoteJust) UpperFrequency() float64 {
	return s.upperFrequency
}
func (s MNoteJust) MidiNoteNumber() uint8 {
	return s.midiNoteNumber
}
func (s MNoteJust) Octave() MOctave {
	return s.octave
}
func (s MNoteJust) OffsetInCents(compareFrequency float64) float64 {
	return Cents(s.exactFrequency, compareFrequency)
}
func (s MNoteJust) String() string {
	return fmt.Sprintf("%s%d", s.noteName, s.octave.octave)
}

func (s MNoteJust) Note() MGNote {
	return s.note
}
