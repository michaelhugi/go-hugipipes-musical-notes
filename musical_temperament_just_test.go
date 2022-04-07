package go_hugipipes_musical_notes

import (
	tu "github.com/informaticon/lib.go.base.test-utils"
	"testing"
)

func TestNewMTemperamentJustCA440(t *testing.T) {
	tj := NewMTemperamentJust(C, A, Octave4, 440)
	tu.AssertEq(tj.freqC0, 27.5/scaleJust9, t)
	tu.AssertEq(tj.baseNote, C, t)
	tu.AssertEq(roundFloat(AddCents(tj.freqC0, intervalToCents(scaleJust1)), 2), roundFloat(tj.freqCis0, 2), t)
	tu.AssertEq(roundFloat(AddCents(tj.freqC0, intervalToCents(scaleJust2)), 2), roundFloat(tj.freqD0, 2), t)
	tu.AssertEq(roundFloat(AddCents(tj.freqC0, intervalToCents(scaleJust3)), 2), roundFloat(tj.freqDis0, 2), t)
	tu.AssertEq(roundFloat(AddCents(tj.freqC0, intervalToCents(scaleJust4)), 2), roundFloat(tj.freqE0, 2), t)
	tu.AssertEq(roundFloat(AddCents(tj.freqC0, intervalToCents(scaleJust5)), 2), roundFloat(tj.freqF0, 2), t)
	tu.AssertEq(roundFloat(AddCents(tj.freqC0, intervalToCents(scaleJust6)), 2), roundFloat(tj.freqFis0, 2), t)
	tu.AssertEq(roundFloat(AddCents(tj.freqC0, intervalToCents(scaleJust7)), 2), roundFloat(tj.freqG0, 2), t)
	tu.AssertEq(roundFloat(AddCents(tj.freqC0, intervalToCents(scaleJust8)), 2), roundFloat(tj.freqGis0, 2), t)
	tu.AssertEq(roundFloat(AddCents(tj.freqC0, intervalToCents(scaleJust9)), 2), roundFloat(tj.freqA0, 2), t)
	tu.AssertEq(roundFloat(AddCents(tj.freqC0, intervalToCents(scaleJust10)), 2), roundFloat(tj.freqAis0, 2), t)
	tu.AssertEq(roundFloat(AddCents(tj.freqC0, intervalToCents(scaleJust11)), 2), roundFloat(tj.freqB0, 2), t)
	tu.AssertEq(roundFloat(tj.freqA0, 2), 27.50, t)
}

func TestNewMTemperamentJustCisA440(t *testing.T) {
	tj := NewMTemperamentJust(Cis, A, Octave4, 440)
	tu.AssertEq(tj.freqCis0, 27.5/scaleJust8, t)
	tu.AssertEq(tj.baseNote, Cis, t)
	tu.AssertEq(roundFloat(AddCents(tj.freqCis0, intervalToCents(scaleJust1)), 2), roundFloat(tj.freqD0, 2), t)
	tu.AssertEq(roundFloat(AddCents(tj.freqCis0, intervalToCents(scaleJust2)), 2), roundFloat(tj.freqDis0, 2), t)
	tu.AssertEq(roundFloat(AddCents(tj.freqCis0, intervalToCents(scaleJust3)), 2), roundFloat(tj.freqE0, 2), t)
	tu.AssertEq(roundFloat(AddCents(tj.freqCis0, intervalToCents(scaleJust4)), 2), roundFloat(tj.freqF0, 2), t)
	tu.AssertEq(roundFloat(AddCents(tj.freqCis0, intervalToCents(scaleJust5)), 2), roundFloat(tj.freqFis0, 2), t)
	tu.AssertEq(roundFloat(AddCents(tj.freqCis0, intervalToCents(scaleJust6)), 2), roundFloat(tj.freqG0, 2), t)
	tu.AssertEq(roundFloat(AddCents(tj.freqCis0, intervalToCents(scaleJust7)), 2), roundFloat(tj.freqGis0, 2), t)
	tu.AssertEq(roundFloat(AddCents(tj.freqCis0, intervalToCents(scaleJust8)), 2), roundFloat(tj.freqA0, 2), t)
	tu.AssertEq(roundFloat(AddCents(tj.freqCis0, intervalToCents(scaleJust9)), 2), roundFloat(tj.freqAis0, 2), t)
	tu.AssertEq(roundFloat(AddCents(tj.freqCis0, intervalToCents(scaleJust10)), 2), roundFloat(tj.freqB0, 2), t)
	tu.AssertEq(roundFloat(AddCents(tj.freqCis0, intervalToCents(1/scaleJust1)), 2), roundFloat(tj.freqC0, 2), t)
	tu.AssertEq(roundFloat(tj.freqA0, 2), 27.50, t)
}

func TestNewMTemperamentJustDA440(t *testing.T) {
	/*	tj := NewMTemperamentJust(D, A, Octave4, 440)
		tu.AssertEq(tj.freqD0, 27.5/scaleJust7, t)
		tu.AssertEq(tj.baseNote, D, t)
		tu.AssertEq(roundFloat(AddCents(tj.freqD0, intervalToCents(scaleJust1)), 2), roundFloat(tj.freqDis0, 2), t)
		tu.AssertEq(roundFloat(AddCents(tj.freqD0, intervalToCents(scaleJust2)), 2), roundFloat(tj.freqE0, 2), t)
		tu.AssertEq(roundFloat(AddCents(tj.freqD0, intervalToCents(scaleJust3)), 2), roundFloat(tj.freqF0, 2), t)
		tu.AssertEq(roundFloat(AddCents(tj.freqD0, intervalToCents(scaleJust4)), 2), roundFloat(tj.freqFis0, 2), t)
		tu.AssertEq(roundFloat(AddCents(tj.freqD0, intervalToCents(scaleJust5)), 2), roundFloat(tj.freqG0, 2), t)
		tu.AssertEq(roundFloat(AddCents(tj.freqD0, intervalToCents(scaleJust6)), 2), roundFloat(tj.freqGis0, 2), t)
		tu.AssertEq(roundFloat(AddCents(tj.freqD0, intervalToCents(scaleJust7)), 2), roundFloat(tj.freqA0, 2), t)
		tu.AssertEq(roundFloat(AddCents(tj.freqD0, intervalToCents(scaleJust8)), 2), roundFloat(tj.freqAis0, 2), t)
		tu.AssertEq(roundFloat(AddCents(tj.freqD0, intervalToCents(scaleJust9)), 2), roundFloat(tj.freqB0, 2), t)
		tu.AssertEq(roundFloat(AddCents(tj.freqD0, intervalToCents(1+scaleJust8)), 2), roundFloat(tj.freqC0, 2), t)
		tu.AssertEq(roundFloat(AddCents(tj.freqD0, intervalToCents(1/scaleJust1)), 2), roundFloat(tj.freqCis0, 2), t)
		tu.AssertEq(roundFloat(tj.freqA0, 2), 27.50, t)*/
}
