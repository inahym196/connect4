package connect4

import "fmt"

var (
	ErrWidthOutOfRange        = fmt.Errorf("width out of range")
	ErrColumnIsFull           = fmt.Errorf("column is full")
	ErrGameHasAlreadyFinished = fmt.Errorf("game has already finished")
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
	board := make([][]Piece, BoardWidth)
	for i := range BoardWidth {
		board[i] = make([]Piece, BoardHeight)
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

func (g *Game) countDirection(col, height, dx, dy int, color Piece) int {
	count := 0
	col += dx
	height += dy
	for 0 <= col && col < BoardWidth && 0 <= height && height < BoardHeight && g.Board[col][height] == color {
		count++
		col += dx
		height += dy
	}
	return count
}

func (g *Game) CheckWin(col, height int) bool {
	color := g.Board[col][height]
	dirs := [][2]int{
		// {dx, dy}
		{0, 1},
		{1, 0},
		{1, 1},
		{-1, -1},
	}
	for _, d := range dirs {
		count := 1
		count += g.countDirection(col, height, d[0], d[1], color)
		count += g.countDirection(col, height, -d[0], -d[1], color)
		if count >= 4 {
			return true
		}
	}
	return false
}

func (g *Game) putPiece(column int, piece Piece) (height int, err error) {
	if !(0 <= column && column < BoardWidth) {
		return -1, ErrWidthOutOfRange
	}
	for i := BoardHeight - 1; i >= 0; i-- {
		if g.Board[column][i] == PieceEmpty {
			g.Board[column][i] = piece
			return i, nil
		}
	}
	return -1, ErrColumnIsFull
}

func (g *Game) PutPiece(column int) error {
	if g.Finished {
		return ErrGameHasAlreadyFinished
	}
	height, err := g.putPiece(column, g.Next)
	if err != nil {
		return err
	}
	g.Next = g.nextPiece()
	if g.CheckWin(column, height) {
		g.Finished = true
		g.Winner = g.nextPiece()
	}
	return nil
}
