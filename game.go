package connect4

import "fmt"

var (
	ErrColumnOutOfRange = fmt.Errorf("column out of range")
	ErrColumnFull       = fmt.Errorf("column is full")
)

const (
	BoardHeight = 6
	BoardWidth  = 7
)

type Piece int

const (
	PieceEmpty Piece = iota
	PieceYellow
	PieceRed
)

type Game struct {
	Finished bool
	Board    [][]Piece
	Next     Piece
}

func initPieces() [][]Piece {
	board := make([][]Piece, BoardHeight)
	for i := range BoardHeight {
		board[i] = make([]Piece, BoardWidth)
	}
	return board
}

func NewGame() *Game {
	return &Game{false, initPieces(), PieceYellow}
}

func (g *Game) nextPiece() Piece {
	if g.Next == PieceRed {
		return PieceYellow
	}
	return PieceRed
}

func (g *Game) PutPiece(column int) error {
	if !(0 <= column && column < BoardWidth) {
		return ErrColumnOutOfRange
	}
	if g.Board[0][column] != PieceEmpty {
		return ErrColumnFull
	}
	for i := BoardHeight - 1; i >= 0; i-- {
		if g.Board[i][column] == PieceEmpty {
			g.Board[i][column] = g.Next
			g.Next = g.nextPiece()
			break
		}
	}
	return nil
}
