package connect4

import "fmt"

var (
	ErrColumnOutOfRange       = fmt.Errorf("column out of range")
	ErrColumnIsFull           = fmt.Errorf("column is full")
	ErrGameHasAlreadyFinished = fmt.Errorf("game has already finished")
)

const (
	BoardColumns = 7
	BoardRows    = 6
)

type Piece int

const (
	PieceEmpty Piece = iota
	PieceYellow
	PieceRed
)

type PlayerPiece Piece

const (
	PlayerPieceYellow = PlayerPiece(PieceYellow)
	PlayerPieceRed    = PlayerPiece(PieceRed)
)

func (p PlayerPiece) Opponent() PlayerPiece {
	if p == PlayerPieceRed {
		return PlayerPieceYellow
	}
	return PlayerPieceYellow
}

type Board [BoardColumns][BoardRows]Piece

func (b Board) DropPiece(col int, pp PlayerPiece) (row int, err error) {
	if !(0 <= col && col < len(b)) {
		return -1, ErrColumnOutOfRange
	}
	for i := len(b[col]) - 1; i >= 0; i-- {
		if b[col][i] == PieceEmpty {
			b[col][i] = Piece(pp)
			return i, nil
		}
	}
	return -1, ErrColumnIsFull
}

func (b Board) countDirection(col, row, dx, dy int, piece Piece) (count int) {
	col += dx
	row += dy
	for 0 <= col && col < len(b) && 0 <= row && row < len(b) && b[col][row] == piece {
		count++
		col += dx
		row += dy
	}
	return count
}

func (b Board) CheckWin(col, row int) bool {
	piece := b[col][row]
	dirs := [][2]int{
		// {dx, dy}
		{0, 1},
		{1, 0},
		{1, 1},
		{-1, -1},
	}
	for _, d := range dirs {
		count := 1
		count += b.countDirection(col, row, d[0], d[1], piece)
		count += b.countDirection(col, row, -d[0], -d[1], piece)
		if count >= 4 {
			return true
		}
	}
	return false
}

type Game struct {
	Finished bool
	Board    Board
	Next     PlayerPiece
	Winner   Piece
}

func NewGame() *Game {
	return &Game{false, Board{}, PlayerPieceYellow, PieceEmpty}
}

func (g *Game) PutPiece(column int) error {
	if g.Finished {
		return ErrGameHasAlreadyFinished
	}
	next := g.Next
	row, err := g.Board.DropPiece(column, next)
	if err != nil {
		return err
	}
	g.Next = next.Opponent()
	if g.Board.CheckWin(column, row) {
		g.Finished = true
		g.Winner = Piece(next)
	}
	return nil
}
