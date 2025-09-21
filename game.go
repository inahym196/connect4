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
	Winner   Piece
}

func initPieces() [][]Piece {
	board := make([][]Piece, BoardHeight)
	for i := range BoardHeight {
		board[i] = make([]Piece, BoardWidth)
	}
	return board
}

func NewGame() *Game {
	return &Game{false, initPieces(), PieceYellow, PieceEmpty}
}

func (g *Game) nextPiece() Piece {
	if g.Next == PieceRed {
		return PieceYellow
	}
	return PieceRed
}

func (g *Game) countDirection(row, col, drow, dcol int, color Piece) int {
	count := 0
	row += drow
	col += dcol
	for 0 <= row && row < BoardHeight && 0 <= col && col < BoardWidth && g.Board[row][col] == color {
		count++
		row += drow
		col += dcol
	}
	return count
}

func (g *Game) CheckWin(row, col int) bool {
	color := g.Board[row][col]
	dirs := [][2]int{
		// {v,h}
		{0, 1},
		{1, 0},
		{1, 1},
		{-1, -1},
	}
	for _, d := range dirs {
		count := 1
		count += g.countDirection(row, col, d[0], d[1], color)
		count += g.countDirection(row, col, -d[0], -d[1], color)
		if count >= 4 {
			return true
		}
	}
	return false
}

func (g *Game) PutPiece(column int) (row int, err error) {
	if !(0 <= column && column < BoardWidth) {
		return -1, ErrColumnOutOfRange
	}
	if g.Board[0][column] != PieceEmpty {
		return -1, ErrColumnFull
	}
	for i := BoardHeight - 1; i >= 0; i-- {
		if g.Board[i][column] == PieceEmpty {
			g.Board[i][column] = g.Next
			g.Next = g.nextPiece()
			return i, nil
		}
	}
	return -1, fmt.Errorf("something is wrong")
}

func (g *Game) PutPieceAndUpdate(column int) error {
	row, err := g.PutPiece(column)
	if err != nil {
		return err
	}
	if g.CheckWin(row, column) {
		g.Finished = true
		g.Winner = g.nextPiece()
	}
	return nil
}
