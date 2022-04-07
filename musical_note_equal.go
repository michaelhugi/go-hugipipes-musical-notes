package go_hugipipes_musical_notes

type MNoteEqual struct {
}

func (s MNoteEqual) ExactFrequency() float64 {
	panic("not implemented")
}
func (s MNoteEqual) FrequencyBelongsToNote(frequency float64) bool {
	panic("not implemented")
}
func (s MNoteEqual) LowerFrequency() float64 {
	panic("not implemented")
}
func (s MNoteEqual) UpperFrequency() float64 {
	panic("not implemented")
}
func (s MNoteEqual) MidiNoteNumber() uint8 {
	panic("not implemented")
}
func (s MNoteEqual) Octave() *MOctave {
	panic("not implemented")
}
func (s MNoteEqual) OffsetInCents(compareFrequency float64) float64 {
	panic("not implemented")
}
func (s MNoteEqual) String() string {
	panic("not implemented")
}
