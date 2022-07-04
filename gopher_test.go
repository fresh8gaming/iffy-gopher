package gopher_test

import (
	"testing"

	iffyGopher "github.com/fresh8gaming/iffy-gopher"
)

func TestGopher(t *testing.T) {
	gopher := &iffyGopher.Gopher{
		Position:    [2]int64{0, 0},
		Field:       [2]int64{5, 5},
		Orientation: "N",
	}

	gopher.Move("F")

	if !(gopher.GetPositionX() == 0 && gopher.GetPositionY() == 1) {
		t.Fail()
	}

	gopher.Move("F")

	if !(gopher.GetPositionX() == 0 && gopher.GetPositionY() == 2) {
		t.Fail()
	}

	gopher.Move("T")
	gopher.Move("F")

	if !(gopher.GetPositionX() == 1 && gopher.GetPositionY() == 2) {
		t.Fail()
	}

	gopher.Move("T")
	gopher.Move("F")

	if !(gopher.GetPositionX() == 1 && gopher.GetPositionY() == 1) {
		t.Fail()
	}
}
