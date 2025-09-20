package connect4_test

import (
	"testing"

	"github.com/inahym196/connect4"
)

func TestNewGame(t *testing.T) {
	game := connect4.NewGame()
	if game == nil {
		t.Fatal("NewGame() returned nil")
	}
	if game.Status != connect4.GameStatusWaiting {
		t.Errorf("expected Status=%d, got %d", connect4.GameStatusWaiting, game.Status)
	}

	if len(game.Board) != connect4.BoardHeight {
		t.Fatalf("expected %d rows, got %d", connect4.BoardHeight, len(game.Board))
	}

	for i, row := range game.Board {
		if len(row) != connect4.BoardWidth {
			t.Fatalf("row %d: expected %d columns, got %d", i, connect4.BoardWidth, len(row))
		}
		for j, p := range row {
			if p != connect4.PieceEmpty {
				t.Fatalf("cell (%d,%d): expected PieceEmpty, got %v", i, j, p)
			}
		}
	}
}
