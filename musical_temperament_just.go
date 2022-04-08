package go_hugipipes_musical_notes

import (
	"errors"
	"fmt"
	"math"
)

//MTemperamentJust represents the Just temperament scale that is often used by bagpipes
type MTemperamentJust struct {
	baseNote   MGNote
	freqC0     float64
	freqCis0   float64
	freqD0     float64
	freqDis0   float64
	freqE0     float64
	freqF0     float64
	freqFis0   float64
	freqG0     float64
	freqGis0   float64
	freqA0     float64
	freqAis0   float64
	freqB0     float64
	allOctaves []MOctave
}

//NewMTemperamentJust ist the constructor for MTemperamentJust
//baseNote is the base note of the scale
//pitchNote is the note that's frequency is set to be exactly a frequency in the pitch
//pitchOctave is the octave of the pitchNote
//pitch is the frequency pitchNote at pitchOctave should have
func NewMTemperamentJust(baseNote MGNote, pitchNote MGNote, pitchOctave MGOctave, pitch float64) *MTemperamentJust {

	pitchNoteFrequency0 := pitch / math.Pow(2, float64(pitchOctave))

	baseNoteFreq := FrequencyOfBaseNoteInJustScale(baseNote, NewSimpleNote(pitchNote, pitchNoteFrequency0))

	freqC0 := 0.0
	freqCis0 := 0.0
	freqD0 := 0.0
	freqDis0 := 0.0
	freqE0 := 0.0
	freqF0 := 0.0
	freqFis0 := 0.0
	freqG0 := 0.0
	freqGis0 := 0.0
	freqA0 := 0.0
	freqAis0 := 0.0
	freqB0 := 0.0

	for _, n := range JustScaleFromBaseNoteWithinOctave(NewSimpleNote(baseNote, baseNoteFreq)) {
		switch n.note {
		case C:
			freqC0 = n.frequency
		case Cis:
			freqCis0 = n.frequency
		case D:
			freqD0 = n.frequency
		case Dis:
			freqDis0 = n.frequency
		case E:
			freqE0 = n.frequency
		case F:
			freqF0 = n.frequency
		case Fis:
			freqFis0 = n.frequency
		case G:
			freqG0 = n.frequency
		case Gis:
			freqGis0 = n.frequency
		case A:
			freqA0 = n.frequency
		case Ais:
			freqAis0 = n.frequency
		case B:
			freqB0 = n.frequency
		}
	}
	return &MTemperamentJust{
		freqC0:     freqC0,
		freqCis0:   freqCis0,
		freqD0:     freqD0,
		freqDis0:   freqDis0,
		freqE0:     freqE0,
		freqF0:     freqF0,
		freqFis0:   freqFis0,
		freqG0:     freqG0,
		freqGis0:   freqGis0,
		freqA0:     freqA0,
		freqAis0:   freqAis0,
		freqB0:     freqB0,
		allOctaves: calcAllOctavesJust(freqC0, freqCis0, freqD0, freqDis0, freqE0, freqF0, freqFis0, freqG0, freqGis0, freqA0, freqAis0, freqB0),
		baseNote:   baseNote,
	}
	/*basePitchFreq := FrequencyOfBaseNoteInJustScale(baseNote, pitchNote, pitch0)
	println("BasePitchFreq: ", basePitchFreq, " ", baseNote)
	freqC0 := justTempFromBaseNote(baseNote, C, basePitchFreq)
	println("C0: ", freqC0, " ", baseNote)
	return &MTemperamentJust{
		freqC0:     freqC0,
		freqCis0:   justTempFromBaseNote(baseNote, Cis, basePitchFreq),
		freqD0:     justTempFromBaseNote(baseNote, D, basePitchFreq),
		freqDis0:   justTempFromBaseNote(baseNote, Dis, basePitchFreq),
		freqE0:     justTempFromBaseNote(baseNote, E, basePitchFreq),
		freqF0:     justTempFromBaseNote(baseNote, F, basePitchFreq),
		freqFis0:   justTempFromBaseNote(baseNote, Fis, basePitchFreq),
		freqG0:     justTempFromBaseNote(baseNote, G, basePitchFreq),
		freqGis0:   justTempFromBaseNote(baseNote, Gis, basePitchFreq),
		freqA0:     justTempFromBaseNote(baseNote, A, basePitchFreq),
		freqAis0:   justTempFromBaseNote(baseNote, Ais, basePitchFreq),
		freqB0:     justTempFromBaseNote(baseNote, B, basePitchFreq),
		allOctaves: calcAllOctavesJust(freqC0),
		baseNote:   baseNote,
	}*/
}

//calcAllOctavesJust calculates all the octaves in the midi range for the Just temperament
func calcAllOctavesJust(freqC0 float64, freqCis0 float64, freqD0 float64, freqDis0 float64, freqE0 float64, freqF0 float64, freqFis0 float64, freqG0 float64, freqGis0 float64, freqA0 float64, freqAis0 float64, freqB0 float64) []MOctave {
	octaves := make([]MOctave, 11)
	octaves[0] = newMOctaveJust(OctaveMinus1, freqC0, freqCis0, freqD0, freqDis0, freqE0, freqF0, freqFis0, freqG0, freqGis0, freqA0, freqAis0, freqB0, AddCents(freqC0/2, -50))
	octaves[1] = newMOctaveJust(Octave0, freqC0, freqCis0, freqD0, freqDis0, freqE0, freqF0, freqFis0, freqG0, freqGis0, freqA0, freqAis0, freqB0, octaves[0].UpperFrequency())
	octaves[2] = newMOctaveJust(Octave1, freqC0, freqCis0, freqD0, freqDis0, freqE0, freqF0, freqFis0, freqG0, freqGis0, freqA0, freqAis0, freqB0, octaves[1].UpperFrequency())
	octaves[3] = newMOctaveJust(Octave2, freqC0, freqCis0, freqD0, freqDis0, freqE0, freqF0, freqFis0, freqG0, freqGis0, freqA0, freqAis0, freqB0, octaves[2].UpperFrequency())
	octaves[4] = newMOctaveJust(Octave3, freqC0, freqCis0, freqD0, freqDis0, freqE0, freqF0, freqFis0, freqG0, freqGis0, freqA0, freqAis0, freqB0, octaves[3].UpperFrequency())
	octaves[5] = newMOctaveJust(Octave4, freqC0, freqCis0, freqD0, freqDis0, freqE0, freqF0, freqFis0, freqG0, freqGis0, freqA0, freqAis0, freqB0, octaves[4].UpperFrequency())
	octaves[6] = newMOctaveJust(Octave5, freqC0, freqCis0, freqD0, freqDis0, freqE0, freqF0, freqFis0, freqG0, freqGis0, freqA0, freqAis0, freqB0, octaves[5].UpperFrequency())
	octaves[7] = newMOctaveJust(Octave6, freqC0, freqCis0, freqD0, freqDis0, freqE0, freqF0, freqFis0, freqG0, freqGis0, freqA0, freqAis0, freqB0, octaves[6].UpperFrequency())
	octaves[8] = newMOctaveJust(Octave7, freqC0, freqCis0, freqD0, freqDis0, freqE0, freqF0, freqFis0, freqG0, freqGis0, freqA0, freqAis0, freqB0, octaves[7].UpperFrequency())
	octaves[9] = newMOctaveJust(Octave8, freqC0, freqCis0, freqD0, freqDis0, freqE0, freqF0, freqFis0, freqG0, freqGis0, freqA0, freqAis0, freqB0, octaves[8].UpperFrequency())
	octaves[10] = newMOctaveJust(Octave9, freqC0, freqCis0, freqD0, freqDis0, freqE0, freqF0, freqFis0, freqG0, freqGis0, freqA0, freqAis0, freqB0, octaves[9].UpperFrequency())
	return octaves
}

//AllOctaves returns all available octaves within the midi range
func (s MTemperamentJust) AllOctaves() []MOctave {
	return s.allOctaves
}

//NoteFromFrequency returns the musical note a Frequency belongs to
func (s MTemperamentJust) NoteFromFrequency(frequency float64) (MNote, error) {
	o, err := s.OctaveFromFrequency(frequency)
	if err != nil {
		return nil, err
	}
	od := o
	for _, n := range od.AllNotes() {
		if n.FrequencyBelongsToNote(frequency) {
			return n, nil
		}
	}
	return nil, errors.New(fmt.Sprintf("Note for frequency %f not found", frequency))
}

//OctaveFromFrequency returns the octave a Frequency belongs to
func (s MTemperamentJust) OctaveFromFrequency(frequency float64) (MOctave, error) {
	for _, o := range s.allOctaves {
		if o.FrequencyBelongsToOctave(frequency) {
			return o, nil
		}
	}
	return nil, errors.New(fmt.Sprintf("Octave for frequency %f not found", frequency))
}

//Octave returns the MOctave for the temperament from MGOctave
func (s MTemperamentJust) Octave(octave MGOctave) MOctave {
	for _, o := range s.allOctaves {
		if o.Octave() == octave {
			return o
		}
	}
	panic(fmt.Sprintf("Octave %d not found", octave))
}
