package vierkantle_test

import (
	"testing"

	"github.com/go-test/deep"
	"github.com/sgielen/vierkantle/pkg/vierkantle"
)

func TestInitialNavigationsFrom(t *testing.T) {
	board := vierkantle.NewBoard(3, 3)

	navigations := board.NextNavigationsFrom(vierkantle.Path{})
	expectedNavigations := []vierkantle.Path{
		{{X: 0, Y: 0}},
		{{X: 1, Y: 0}},
		{{X: 2, Y: 0}},
		{{X: 0, Y: 1}},
		{{X: 1, Y: 1}},
		{{X: 2, Y: 1}},
		{{X: 0, Y: 2}},
		{{X: 1, Y: 2}},
		{{X: 2, Y: 2}},
	}

	if diff := deep.Equal(navigations, expectedNavigations); diff != nil {
		t.Error(diff)
	}
}

func TestNavigationsFromMiddle(t *testing.T) {
	board := vierkantle.NewBoard(3, 3)

	navigations := board.NextNavigationsFrom(vierkantle.Path{{X: 1, Y: 1}})
	expectedNavigations := []vierkantle.Path{
		{{X: 1, Y: 1}, {X: 0, Y: 0}},
		{{X: 1, Y: 1}, {X: 0, Y: 1}},
		{{X: 1, Y: 1}, {X: 0, Y: 2}},
		{{X: 1, Y: 1}, {X: 1, Y: 0}},
		{{X: 1, Y: 1}, {X: 1, Y: 2}},
		{{X: 1, Y: 1}, {X: 2, Y: 0}},
		{{X: 1, Y: 1}, {X: 2, Y: 1}},
		{{X: 1, Y: 1}, {X: 2, Y: 2}},
	}

	if diff := deep.Equal(navigations, expectedNavigations); diff != nil {
		t.Error(diff)
	}
}

func TestNavigationsFromSide(t *testing.T) {
	board := vierkantle.NewBoard(3, 3)

	navigations := board.NextNavigationsFrom(vierkantle.Path{{X: 1, Y: 1}, {X: 0, Y: 1}})
	expectedNavigations := []vierkantle.Path{
		{{X: 1, Y: 1}, {X: 0, Y: 1}, {X: 0, Y: 0}},
		{{X: 1, Y: 1}, {X: 0, Y: 1}, {X: 0, Y: 2}},
		{{X: 1, Y: 1}, {X: 0, Y: 1}, {X: 1, Y: 0}},
		{{X: 1, Y: 1}, {X: 0, Y: 1}, {X: 1, Y: 2}},
	}

	if diff := deep.Equal(navigations, expectedNavigations); diff != nil {
		t.Error(diff)
	}
}
