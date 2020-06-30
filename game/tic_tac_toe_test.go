package game

import (
	"github.com/google/go-cmp/cmp"
	"testing"
)

func TestTicTacToe(t *testing.T)  {
	t.Run("x plays first", func(t *testing.T) {
		got := runGame()
		expected :=  GameState{
			NextPlayer: "X",
			Status:     "on",
		}

		if ! cmp.Equal(got, expected) {
			t.Errorf("got %s, want %s expected", got, expected)
		}
	})
}

type GameState struct {
	NextPlayer string
	Status     string
}

func runGame() GameState {
	return GameState{
		NextPlayer: "X",
		Status:     "on",
	}
}
