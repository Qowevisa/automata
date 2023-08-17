package automata_test

import (
	"testing"

	"github.com/Qowevisa/automata"
)

func TestAutomata(t *testing.T) {
	auto, err := automata.CreateAutomata(
		automata.CreateRule("C1", "val", "C2"),
		automata.CreateRule("C2", "val2", "C3"),
		automata.CreateRule("C3", "a", "C1"),
		automata.CreateRule("C1", "b", "C1"),
		automata.CreateRule("C2", "c", "C1"),
	)
	var expectedError error
	var expectedAutoCellNumbers uint
	expectedError = nil
	expectedAutoCellNumbers = 3
	if err != expectedError {
		t.Fatalf("Expect error to be %v, got: %v", expectedError, err)
	}
	if auto.CellsNumber != expectedAutoCellNumbers {
		t.Fatalf("Expect number of cells in Automata to be %d, got: %d", expectedAutoCellNumbers, auto.CellsNumber)
	}
}
