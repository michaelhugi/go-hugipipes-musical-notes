package go_hugipipes_musical_notes

import (
	tu "github.com/informaticon/lib.go.base.test-utils"
	"testing"
)

func checkMNInterface(mn MNote) {

}

func TestAddHalftones(t *testing.T) {
	tu.AssertEq(AddHalftones(C, 0), C, t)
	tu.AssertEq(AddHalftones(C, 1), Cis, t)
	tu.AssertEq(AddHalftones(C, 2), D, t)
	tu.AssertEq(AddHalftones(C, 3), Dis, t)
	tu.AssertEq(AddHalftones(C, 4), E, t)
	tu.AssertEq(AddHalftones(C, 5), F, t)
	tu.AssertEq(AddHalftones(C, 6), Fis, t)
	tu.AssertEq(AddHalftones(C, 7), G, t)
	tu.AssertEq(AddHalftones(C, 8), Gis, t)
	tu.AssertEq(AddHalftones(C, 9), A, t)
	tu.AssertEq(AddHalftones(C, 10), Ais, t)
	tu.AssertEq(AddHalftones(C, 11), B, t)
	tu.AssertEq(AddHalftones(C, 12), C, t)
	tu.AssertEq(AddHalftones(C, 13), Cis, t)
	tu.AssertEq(AddHalftones(C, 14), D, t)
	tu.AssertEq(AddHalftones(C, 15), Dis, t)
	tu.AssertEq(AddHalftones(C, 16), E, t)
	tu.AssertEq(AddHalftones(C, 17), F, t)
	tu.AssertEq(AddHalftones(C, 18), Fis, t)
	tu.AssertEq(AddHalftones(C, 19), G, t)
	tu.AssertEq(AddHalftones(C, 20), Gis, t)
	tu.AssertEq(AddHalftones(C, 21), A, t)
	tu.AssertEq(AddHalftones(C, 22), Ais, t)
	tu.AssertEq(AddHalftones(C, 23), B, t)
	tu.AssertEq(AddHalftones(C, 24), C, t)
	tu.AssertEq(AddHalftones(C, 25), Cis, t)
	tu.AssertEq(AddHalftones(C, 26), D, t)
	tu.AssertEq(AddHalftones(C, 27), Dis, t)
	tu.AssertEq(AddHalftones(C, 28), E, t)

	tu.AssertEq(AddHalftones(F, 0), F, t)
	tu.AssertEq(AddHalftones(F, 1), Fis, t)
	tu.AssertEq(AddHalftones(F, 2), G, t)
	tu.AssertEq(AddHalftones(F, 3), Gis, t)
	tu.AssertEq(AddHalftones(F, 4), A, t)
	tu.AssertEq(AddHalftones(F, 5), Ais, t)
	tu.AssertEq(AddHalftones(F, 6), B, t)
	tu.AssertEq(AddHalftones(F, 7), C, t)
	tu.AssertEq(AddHalftones(F, 8), Cis, t)
	tu.AssertEq(AddHalftones(F, 9), D, t)
	tu.AssertEq(AddHalftones(F, 10), Dis, t)
	tu.AssertEq(AddHalftones(F, 11), E, t)
	tu.AssertEq(AddHalftones(F, 12), F, t)
	tu.AssertEq(AddHalftones(F, 13), Fis, t)
	tu.AssertEq(AddHalftones(F, 14), G, t)
	tu.AssertEq(AddHalftones(F, 15), Gis, t)
	tu.AssertEq(AddHalftones(F, 16), A, t)
	tu.AssertEq(AddHalftones(F, 17), Ais, t)
	tu.AssertEq(AddHalftones(F, 18), B, t)
	tu.AssertEq(AddHalftones(F, 19), C, t)
	tu.AssertEq(AddHalftones(F, 20), Cis, t)
	tu.AssertEq(AddHalftones(F, 21), D, t)
	tu.AssertEq(AddHalftones(F, 22), Dis, t)
	tu.AssertEq(AddHalftones(F, 23), E, t)
	tu.AssertEq(AddHalftones(F, 24), F, t)
	tu.AssertEq(AddHalftones(F, 25), Fis, t)
	tu.AssertEq(AddHalftones(F, 26), G, t)
	tu.AssertEq(AddHalftones(F, 27), Gis, t)
	tu.AssertEq(AddHalftones(F, 28), A, t)

}
