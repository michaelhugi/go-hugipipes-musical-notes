package go_hugipipes_musical_notes

import (
	"math"
	"testing"
)
import tu "github.com/informaticon/lib.go.base.test-utils"

func TestAddCents(t *testing.T) {
	tu.AssertEq(roundFloat(AddCents(16.35, 100), 2), 17.32, t)
	tu.AssertEq(roundFloat(AddCents(16.35, 1200), 2), 16.35*2, t)
	tu.AssertEq(roundFloat(AddCents(16.35, 300), 2), 19.44, t)
	tu.AssertEq(roundFloat(AddCents(523.25, -300), 2), 440.00, t)
}

func TestCents(t *testing.T) {
	tu.AssertEq(roundFloat(Cents(440.00, 466.16), 0), 100, t)
	tu.AssertEq(roundFloat(Cents(440.00, 880), 0), 1200, t)
	tu.AssertEq(roundFloat(Cents(440.00, 932.33), 0), 1300, t)
	tu.AssertEq(roundFloat(Cents(1661.22, 1567.98), 0), -100, t)
}

func roundFloat(v float64, digits float64) float64 {
	f := math.Pow(10, digits)
	v = v * f
	v = math.Round(v)
	return v / f
}
