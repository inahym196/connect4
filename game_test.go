package connect4_test

import (
	"testing"

	"github.com/inahym196/connect4"
)

func TestBoard_DropPiece(t *testing.T) {
	t.Run("ボード外には置けない", func(t *testing.T) {
		b := connect4.Board{}
		if _, err := b.DropPiece(-1, connect4.PlayerPieceRed); err == nil {
			t.Errorf("expected err, got %v", err)
		}
		if _, err := b.DropPiece(7, connect4.PlayerPieceRed); err == nil {
			t.Errorf("expected err, got %v", err)
		}
	})
	t.Run("最初の一個を置いてみる", func(t *testing.T) {
		b := connect4.Board{}
		col := 0
		pp := connect4.PlayerPieceRed

		row, err := b.DropPiece(col, pp)
		if err != nil {
			t.Fatal(err)
		}

		// 最下段にピースが置かれているか
		if row != connect4.BoardRows-1 {
			t.Errorf("expected row %d, got %d", connect4.BoardRows-1, row)
		}
		// 最下段のピースは置いたピースと一致するか
		if b[col][connect4.BoardRows-1] != connect4.Piece(pp) {
			t.Errorf("expected piece %v at bottom, got %v", connect4.Piece(pp), b[col][connect4.BoardRows-1])
		}
	})
	t.Run("2個目を置いてみる", func(t *testing.T) {
		b := connect4.Board{}
		col := 0
		pp1 := connect4.PlayerPieceYellow
		pp2 := connect4.PlayerPieceRed

		b.DropPiece(col, pp1)
		row, err := b.DropPiece(col, pp2)
		if err != nil {
			t.Fatal(err)
		}

		// 二段目にピースが置かれているか
		if row != connect4.BoardRows-2 {
			t.Errorf("expected row %d, got %d", connect4.BoardRows-2, row)
		}
		// 二段目のピースは置いたピースと一致するか
		if b[col][connect4.BoardRows-2] != connect4.Piece(pp2) {
			t.Errorf("expected piece %v at bottom, got %v", connect4.Piece(pp2), b[col][connect4.BoardRows-1])
		}
	})
	t.Run("満タンのColumnにはピースを置けない", func(t *testing.T) {
		b := connect4.Board{}
		col := 0
		for range connect4.BoardRows {
			if _, err := b.DropPiece(col, connect4.PlayerPieceRed); err != nil {
				t.Fatal(err)
			}
		}

		if _, err := b.DropPiece(col, connect4.PlayerPieceRed); err == nil {
			t.Error(err)
		}
	})

}
func TestBoard_CheckWin(t *testing.T) {
	tests := []struct {
		name   string
		setup  func() connect4.Board
		col    int
		row    int
		expect bool
	}{
		{
			name: "horizontal win",
			setup: func() connect4.Board {
				b := connect4.Board{}
				for i := range 4 {
					b.DropPiece(i, connect4.PlayerPieceRed)
				}
				return b
			},
			col:    3,
			row:    5, // connect4.BoardRows - 1
			expect: true,
		},
		{
			name: "vertical win",
			setup: func() connect4.Board {
				b := connect4.Board{}
				for range 4 {
					b.DropPiece(0, connect4.PlayerPieceYellow)
				}
				return b
			},
			col:    0,
			row:    3,
			expect: true,
		},
		{
			name: "diagonal win ↙︎",
			setup: func() connect4.Board {
				b := connect4.Board{}
				b.DropPiece(0, connect4.PlayerPieceRed) //0,5
				b.DropPiece(1, connect4.PlayerPieceYellow)
				b.DropPiece(1, connect4.PlayerPieceRed) //1,4
				b.DropPiece(2, connect4.PlayerPieceYellow)
				b.DropPiece(2, connect4.PlayerPieceYellow)
				b.DropPiece(2, connect4.PlayerPieceRed) //2,3
				b.DropPiece(3, connect4.PlayerPieceYellow)
				b.DropPiece(3, connect4.PlayerPieceYellow)
				b.DropPiece(3, connect4.PlayerPieceYellow)
				b.DropPiece(3, connect4.PlayerPieceRed) //3,2
				return b
			},
			col:    3,
			row:    2,
			expect: true,
		},
		{
			name: "no win",
			setup: func() connect4.Board {
				b := connect4.Board{}
				b.DropPiece(0, connect4.PlayerPieceRed)
				b.DropPiece(1, connect4.PlayerPieceYellow)
				b.DropPiece(2, connect4.PlayerPieceRed)
				b.DropPiece(3, connect4.PlayerPieceRed)
				b.DropPiece(4, connect4.PlayerPieceRed)
				return b
			},
			col:    4,
			row:    5,
			expect: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := tt.setup()
			got := b.CheckWin(tt.col, tt.row)
			if got != tt.expect {
				t.Errorf("expected checkWin() = %v, got %v", tt.expect, got)
			}
		})
	}
}

func TestNewGame(t *testing.T) {
	nextPiece := connect4.PlayerPieceYellow

	game := connect4.NewGame()

	if game == nil {
		t.Fatal("NewGame() returned nil")
	}
	if game.Finished != false {
		t.Errorf("expected false, got %T", game.Finished)
	}
	if game.Next != nextPiece {
		t.Errorf("expected Turn %v, got %v", nextPiece, game.Next)
	}
	if game.Winner != connect4.PieceEmpty {
		t.Errorf("expected %v, got %v", connect4.PieceEmpty, game.Winner)
	}
}

func TestPutPiece(t *testing.T) {
	t.Run("最初の一個を置いてみる", func(t *testing.T) {
		game := connect4.NewGame()
		col := 0
		nextPiece := connect4.PlayerPieceRed

		if err := game.PutPiece(col); err != nil {
			t.Fatal(err)
		}

		if game.Next != nextPiece {
			t.Errorf("expected Next is %v, got %v", nextPiece, game.Next)
		}
		if game.Finished != false {
			t.Errorf("expected finished is false, got %T", game.Finished)
		}
		if game.Winner != connect4.PieceEmpty {
			t.Errorf("expected winner is %v, got %v", connect4.PieceEmpty, game.Winner)
		}
	})
	t.Run("2個目を置いてみる", func(t *testing.T) {
		game := connect4.NewGame()
		col := 0
		secondPiece := connect4.PlayerPieceYellow

		game.PutPiece(col)
		if err := game.PutPiece(col); err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if game.Next != secondPiece {
			t.Errorf("expected Next %v, got %v", secondPiece, game.Next)
		}
		if game.Finished != false {
			t.Errorf("expected finished is false, got %T", game.Finished)
		}
		if game.Winner != connect4.PieceEmpty {
			t.Errorf("expected winner %v, got %v", connect4.PieceEmpty, game.Winner)
		}
	})
	t.Run("4個連続すると勝てる", func(t *testing.T) {
		game := connect4.NewGame()
		myCol := 0
		myColor := connect4.PieceYellow
		oppCol := connect4.BoardColumns - 1
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
			t.Errorf("expected winner is %v, got %v", myColor, game.Winner)
		}
	})

	t.Run("終了したGameにはピースを置けない", func(t *testing.T) {
		game := connect4.NewGame()
		myCol := 0
		oppCol := connect4.BoardColumns - 1
		moves := []int{myCol, oppCol, myCol, oppCol, myCol, oppCol, myCol}
		for _, lane := range moves {
			err := game.PutPiece(lane)
			if err != nil {
				t.Fatal(err)
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
