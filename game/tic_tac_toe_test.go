package game

import (
	"testing"
)

func TestTicTacToe(t *testing.T)  {
	t.Run("x plays first", func(t *testing.T) {
		got := runGame()
		expected :=  GameState{
			nextPlayer: "X",
			status: "on",
		}

		if got.nextPlayer != expected.nextPlayer {
			t.Errorf("got %s, want %s expected", got.nextPlayer, expected.nextPlayer)
		}
	})
}

type GameState struct {
	nextPlayer string
	status string
}

func runGame() *GameState {
	return &GameState{
		nextPlayer: "X",
		status: "on",
	}
}
