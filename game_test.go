package connect4_test

import (
	"testing"

	"github.com/inahym196/connect4"
)

func TestNewGame(t *testing.T) {
	nextPiece := connect4.PieceYellow

	game := connect4.NewGame()

	if game == nil {
		t.Fatal("NewGame() returned nil")
	}
	if game.Finished != false {
		t.Errorf("expected false, got %T", game.Finished)
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

	if game.Next != nextPiece {
		t.Errorf("expected Turn=%d, got %d", nextPiece, game.Next)
	}
}

func TestPutPiece(t *testing.T) {
	t.Run(connect4.ErrColumnOutOfRange.Error(), func(t *testing.T) {
		game := connect4.NewGame()

		if err := game.PutPiece(-1); err != connect4.ErrColumnOutOfRange {
			t.Errorf("expected ErrColumnOutOfRange, got %v", err)
		}
		if err := game.PutPiece(connect4.BoardWidth); err != connect4.ErrColumnOutOfRange {
			t.Errorf("expected ErrColumnOutOfRange, got %v", err)
		}
	})

	t.Run("put piece in empty column", func(t *testing.T) {
		game := connect4.NewGame()
		col := 0
		nextPiece := connect4.PieceYellow

		if err := game.PutPiece(col); err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		// 最下段にピースが置かれているか
		if game.Board[connect4.BoardHeight-1][col] != nextPiece {
			t.Errorf("expected piece %v at bottom, got %v", nextPiece, game.Board[connect4.BoardHeight-1][col])
		}

		// ターンが交代しているか
		if game.Next == nextPiece {
			t.Errorf("expected turn to switch, still %v", game.Next)
		}
	})
	t.Run("place piece on partially filled column", func(t *testing.T) {
		game := connect4.NewGame()
		col := 0
		secondPiece := connect4.PieceRed

		game.PutPiece(col)

		if err := game.PutPiece(col); err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		// 二段目にピースが置かれているか
		if game.Board[connect4.BoardHeight-2][col] != secondPiece {
			t.Errorf("expected piece %v at second-to-bottom, got %v", secondPiece, game.Board[connect4.BoardHeight-2][col])
		}
	})

	t.Run(connect4.ErrColumnFull.Error(), func(t *testing.T) {
		game := connect4.NewGame()
		for i := range connect4.BoardHeight {
			game.Board[i][0] = connect4.PieceRed
		}

		if err := game.PutPiece(0); err != connect4.ErrColumnFull {
			t.Errorf("expected ErrColumnFull, got %v", err)
		}
	})
}
