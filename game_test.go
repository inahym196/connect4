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

	if game.Winner != connect4.PieceEmpty {
		t.Errorf("expected PieceEmpty(%d), got %d", connect4.PieceEmpty, game.Winner)
	}
}

func TestCheckWin(t *testing.T) {
	tests := []struct {
		name   string
		setup  func() *connect4.Game
		row    int
		col    int
		expect bool
	}{
		{
			name: "horizontal win",
			setup: func() *connect4.Game {
				g := connect4.NewGame()
				for i := range 4 {
					g.Board[5][i] = connect4.PieceRed
				}
				return g
			},
			row:    5,
			col:    3,
			expect: true,
		},
		{
			name: "vertical win",
			setup: func() *connect4.Game {
				g := connect4.NewGame()
				for i := 2; i < 6; i++ {
					g.Board[i][0] = connect4.PieceYellow
				}
				return g
			},
			row:    5,
			col:    0,
			expect: true,
		},
		{
			name: "diagonal win ↘",
			setup: func() *connect4.Game {
				g := connect4.NewGame()
				g.Board[2][0] = connect4.PieceRed
				g.Board[3][1] = connect4.PieceRed
				g.Board[4][2] = connect4.PieceRed
				g.Board[5][3] = connect4.PieceRed
				return g
			},
			row:    5,
			col:    3,
			expect: true,
		},
		{
			name: "no win",
			setup: func() *connect4.Game {
				g := connect4.NewGame()
				g.Board[5][0] = connect4.PieceRed
				g.Board[5][1] = connect4.PieceYellow
				g.Board[5][2] = connect4.PieceRed
				g.Board[5][3] = connect4.PieceRed
				g.Board[5][4] = connect4.PieceRed
				return g
			},
			row:    5,
			col:    4,
			expect: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := tt.setup()
			got := g.CheckWin(tt.row, tt.col)
			if got != tt.expect {
				t.Errorf("checkWin() = %v, want %v", got, tt.expect)
			}
		})
	}
}

func TestPutPiece(t *testing.T) {

	t.Run("ボード外には置けない", func(t *testing.T) {
		game := connect4.NewGame()

		if err := game.PutPiece(-1); err == nil {
			t.Errorf("expected err, got %v", err)
		}
		if err := game.PutPiece(connect4.BoardWidth); err == nil {
			t.Errorf("expected err, got %v", err)
		}
	})

	t.Run("最初の一個を置いてみる", func(t *testing.T) {
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
		if game.Next == nextPiece {
			t.Errorf("expected turn to switch, still %v", game.Next)
		}
		if game.Finished != false {
			t.Errorf("expected finished is false, got %T", game.Finished)
		}
		if game.Winner != connect4.PieceEmpty {
			t.Errorf("expected winner is still empty(%d), got %d", connect4.PieceEmpty, game.Winner)
		}
	})
	t.Run("2個目を置いてみる", func(t *testing.T) {
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
		if game.Next == secondPiece {
			t.Errorf("expected turn to switch, still %v", game.Next)
		}
		if game.Finished != false {
			t.Errorf("expected finished is false, got %T", game.Finished)
		}
		if game.Winner != connect4.PieceEmpty {
			t.Errorf("expected winner is still empty(%d), got %d", connect4.PieceEmpty, game.Winner)
		}
	})
	t.Run("4個連続すると勝てる", func(t *testing.T) {
		game := connect4.NewGame()
		myCol := 0
		myColor := connect4.PieceYellow
		oppCol := connect4.BoardWidth - 1
		moves := []int{myCol, oppCol, myCol, oppCol, myCol, oppCol, myCol}
		for _, col := range moves {
			err := game.PutPiece(col)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
		}

		if game.Finished != true {
			t.Errorf("expected finished is true, got %T", game.Finished)
		}
		if game.Winner != myColor {
			t.Errorf("expected winner is %d, got %d", myColor, game.Winner)
		}
	})

	t.Run("満タンのレーンにはピースを置けない", func(t *testing.T) {
		game := connect4.NewGame()
		col := 0
		for range connect4.BoardHeight {
			game.PutPiece(col)
		}

		if err := game.PutPiece(col); err != connect4.ErrColumnFull {
			t.Errorf("expected ErrColumnFull, got %v", err)
		}
	})

	t.Run("終了したGameにはピースを置けない", func(t *testing.T) {
		game := connect4.NewGame()
		myCol := 0
		oppCol := connect4.BoardWidth - 1
		moves := []int{myCol, oppCol, myCol, oppCol, myCol, oppCol, myCol}
		for _, col := range moves {
			err := game.PutPiece(col)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
		}
		if game.Finished != true {
			t.Fatalf("expected finished is true, got %T", game.Finished)
		}
		if err := game.PutPiece(myCol); err != connect4.ErrGameHasAlreadyFinished {
			t.Errorf("expected error is %v, got: %v", connect4.ErrGameHasAlreadyFinished, err)
		}
	})

}
