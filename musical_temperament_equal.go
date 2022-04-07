package go_hugipipes_musical_notes

//MTemperamentEqual represents the equal temperament scale that is most often used in music theory, especially in western countries
type MTemperamentEqual struct {
	concertPitchA float64
}

//NewMTemperamentEqual ist the constructor for MTemperamentEqual
//concertPitchA is the frequency in Hz of A. If passed <=0.0, 440Hz will be assumed
func NewMTemperamentEqual(concertPitchA float64) *MTemperamentEqual {
	if concertPitchA <= 0.0 {
		concertPitchA = 440.0
	}
	return &MTemperamentEqual{concertPitchA: concertPitchA}
}

func (s MTemperamentEqual) AllOctaves() []MOctave {
	panic("unimplemented")
}

func (s MTemperamentEqual) MGNoteToNote(n *MGNote, o *MGOctave) *MNote {
	panic("unimplemented")
}

func (s MTemperamentEqual) NoteFromFrequency(frequency float64) (*MNote, error) {
	panic("unimplemented")
}
func (s MTemperamentEqual) OctaveFromFrequency(frequency float64) (*MOctave, error) {
	panic("unimplemented")
}
