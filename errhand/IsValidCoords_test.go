package lemin

import (
	"testing"

	lemin "lemin/functions"
)

func TestIsValidCoords(t *testing.T) {
	data := &lemin.Input{
		Coords: [][]string{{"1", "2"}, {"5", "6"}},
	}

	// Valid coordinates
	coords := []string{"3", "4"}
	result := lemin.IsValidCoords(coords, data)
	if !result {
		t.Errorf("Expected true for valid coordinates, but got false")
	}

	// Invalid coordinates (duplicate)
	coords = []string{"1", "1"}
	result = lemin.IsValidCoords(coords, data)
	if !result {
		t.Errorf("Expected true for duplicate coordinates, but got false")
	}

	// Invalid coordinates (negative X)
	coords = []string{"-1", "2"}
	result = lemin.IsValidCoords(coords, data)
	if !result {
		t.Errorf("Expected true for negative X coordinate, but got false")
	}

	// Invalid coordinates (negative Y)
	coords = []string{"1", "-2"}
	result = lemin.IsValidCoords(coords, data)
	if !result {
		t.Errorf("Expected true for negative Y coordinate, but got false")
	}
}
