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

type Turn byte

const (
	TurnYellow Turn = iota
	TurnRed
)

type Game struct {
	Finished bool
	Board    [][]Piece
	Turn     Turn
}

func initPieces() [][]Piece {
	board := make([][]Piece, BoardHeight)
	for i := range BoardHeight {
		board[i] = make([]Piece, BoardWidth)
	}
	return board
}

func NewGame() *Game {
	return &Game{false, initPieces(), TurnYellow}
}

func (g *Game) turnColor() Piece {
	return Piece(g.Turn + 1)
}

func (g *Game) nextTurn() Turn {
	if g.Turn == TurnRed {
		return TurnYellow
	}
	return TurnRed
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
			g.Board[i][column] = g.turnColor()
			g.Turn = g.nextTurn()
			break
		}
	}
	return nil
}
