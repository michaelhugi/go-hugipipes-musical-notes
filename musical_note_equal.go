package go_hugipipes_musical_notes

import (
	"fmt"
)

type MNoteEqual struct {
	octave         *MOctaveEqual
	exactFrequency float64
	lowerFrequency float64
	upperFrequency float64
	midiNoteNumber uint8
	noteName       string
	note           MGNote
}

func newMNoteEqual(exactFrequency float64, note MGNote, octave *MOctaveEqual, midiNoteNumber uint8, noteName string, lowerFrequency float64) *MNoteEqual {

	return &MNoteEqual{
		octave:         octave,
		exactFrequency: exactFrequency,
		lowerFrequency: lowerFrequency,
		upperFrequency: AddCents(exactFrequency, 50),
		midiNoteNumber: midiNoteNumber,
		noteName:       noteName,
		note:           note,
	}
}

func (s MNoteEqual) ExactFrequency() float64 {
	return s.exactFrequency
}
func (s MNoteEqual) FrequencyBelongsToNote(frequency float64) bool {
	return frequency >= s.lowerFrequency && frequency < s.upperFrequency
}
func (s MNoteEqual) LowerFrequency() float64 {
	return s.lowerFrequency
}
func (s MNoteEqual) UpperFrequency() float64 {
	return s.upperFrequency
}
func (s MNoteEqual) MidiNoteNumber() uint8 {
	return s.midiNoteNumber
}
func (s MNoteEqual) Octave() MOctave {
	return s.octave
}
func (s MNoteEqual) OffsetInCents(compareFrequency float64) float64 {
	return Cents(s.exactFrequency, compareFrequency)
}
func (s MNoteEqual) String() string {
	return fmt.Sprintf("%s%d", s.noteName, s.octave.octave)
}

func (s MNoteEqual) Note() MGNote {
	return s.note
}
