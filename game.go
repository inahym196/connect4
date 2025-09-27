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
	switch p {
	case PlayerPieceRed:
		return PlayerPieceYellow
	case PlayerPieceYellow:
		return PlayerPieceRed
	default:
		panic("invalid player piece")
	}
}

type Board [BoardColumns][BoardRows]Piece

func (b *Board) DropPiece(col int, pp PlayerPiece) (row int, err error) {
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
	for 0 <= col && col < len(b) && 0 <= row && row < len(b[col]) && b[col][row] == piece {
		count++
		col += dx
		row += dy
	}
	return count
}

func (b Board) CheckWin(col, row int) bool {
	piece := b[col][row]
	if piece == PieceEmpty {
		return false
	}
	dirs := [][2]int{
		// {dx, dy}
		{0, 1},
		{1, 0},
		{1, 1},
		{1, -1},
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
	finished bool
	board    Board
	next     PlayerPiece
	winner   Piece
}

func NewGame() *Game {
	return &Game{false, Board{}, PlayerPieceYellow, PieceEmpty}
}

func (g Game) IsFinished() bool {
	return g.finished
}

func (g Game) Board() Board {
	return g.board
}

func (g Game) Next() PlayerPiece {
	return g.next
}

func (g Game) Winner() Piece {
	return g.winner
}

func (g *Game) PutPiece(column int) error {
	if g.finished {
		return ErrGameHasAlreadyFinished
	}
	row, err := g.board.DropPiece(column, g.next)
	if err != nil {
		return err
	}
	if g.board.CheckWin(column, row) {
		g.finished = true
		g.winner = Piece(g.next)
	} else {
		g.next = g.next.Opponent()
	}
	return nil
}
