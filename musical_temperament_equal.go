package go_hugipipes_musical_notes

import (
	"errors"
	"fmt"
	"math"
)

//MTemperamentEqual represents the equal temperament scale that is most often used in music theory, especially in western countries
type MTemperamentEqual struct {
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

//NewMTemperamentEqual ist the constructor for MTemperamentEqual
//concertPitchA is the frequency in Hz of A4. If passed <=0.0, 440Hz will be assumed
func NewMTemperamentEqual(concertPitchA4 float64) *MTemperamentEqual {
	if concertPitchA4 <= 0.0 {
		concertPitchA4 = 440.0
	}
	freqA0 := concertPitchA4 / math.Pow(2, 4)
	freqC0 := AddCents(freqA0, -900)
	return &MTemperamentEqual{
		freqC0:     freqC0,
		freqCis0:   AddCents(freqA0, -800),
		freqD0:     AddCents(freqA0, -700),
		freqDis0:   AddCents(freqA0, -600),
		freqE0:     AddCents(freqA0, -500),
		freqF0:     AddCents(freqA0, -400),
		freqFis0:   AddCents(freqA0, -300),
		freqG0:     AddCents(freqA0, -200),
		freqGis0:   AddCents(freqA0, -100),
		freqA0:     freqA0,
		freqAis0:   AddCents(freqA0, 100),
		freqB0:     AddCents(freqA0, 200),
		allOctaves: calcAllOctavesEqual(freqC0),
	}
}

//calcAllOctavesEqual calculates all the octaves in the midi range for the equal temperament
func calcAllOctavesEqual(freqC0 float64) []MOctave {
	octaves := make([]MOctave, 11)
	octaves[0] = newMOctaveEqual(OctaveMinus1, freqC0, AddCents(freqC0/2, -50))
	octaves[1] = newMOctaveEqual(Octave0, freqC0, octaves[0].UpperFrequency())
	octaves[2] = newMOctaveEqual(Octave1, freqC0, octaves[1].UpperFrequency())
	octaves[3] = newMOctaveEqual(Octave2, freqC0, octaves[2].UpperFrequency())
	octaves[4] = newMOctaveEqual(Octave3, freqC0, octaves[3].UpperFrequency())
	octaves[5] = newMOctaveEqual(Octave4, freqC0, octaves[4].UpperFrequency())
	octaves[6] = newMOctaveEqual(Octave5, freqC0, octaves[5].UpperFrequency())
	octaves[7] = newMOctaveEqual(Octave6, freqC0, octaves[6].UpperFrequency())
	octaves[8] = newMOctaveEqual(Octave7, freqC0, octaves[7].UpperFrequency())
	octaves[9] = newMOctaveEqual(Octave8, freqC0, octaves[8].UpperFrequency())
	octaves[10] = newMOctaveEqual(Octave9, freqC0, octaves[9].UpperFrequency())
	return octaves
}

//AllOctaves returns all available octaves within the midi range
func (s MTemperamentEqual) AllOctaves() []MOctave {
	return s.allOctaves
}

//NoteFromFrequency returns the musical note a Frequency belongs to
func (s MTemperamentEqual) NoteFromFrequency(frequency float64) (MNote, error) {
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
func (s MTemperamentEqual) OctaveFromFrequency(frequency float64) (MOctave, error) {
	for _, o := range s.allOctaves {
		if o.FrequencyBelongsToOctave(frequency) {
			return o, nil
		}
	}
	return nil, errors.New(fmt.Sprintf("Octave for frequency %f not found", frequency))
}

//Octave returns the MOctave for the temperament from MGOctave
func (s MTemperamentEqual) Octave(octave MGOctave) MOctave {
	for _, o := range s.allOctaves {
		if o.Octave() == octave {
			return o
		}
	}
	panic(fmt.Sprintf("Octave %d not found", octave))
}
