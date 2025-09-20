package connect4

type GameStatus int

const (
	GameStatusUnknown GameStatus = iota
	GameStatusWaiting
	GameStatusPlaying
	GameStatusFinished
)

const (
	BoardHeight = 6
	BoardWidth  = 7
)

type Piece int

const (
	PieceEmpty Piece = iota
	PieceRed
	PieceYellow
)

type Turn byte

const (
	TurnFirst Turn = iota
	TurnSecond
)

type Game struct {
	Status GameStatus
	Board  [][]Piece
	Turn   Turn
}

func initPieces() [][]Piece {
	board := make([][]Piece, BoardHeight)
	for i := range BoardHeight {
		board[i] = make([]Piece, BoardWidth)
	}
	return board
}

func NewGame() *Game {
	return &Game{GameStatusWaiting, initPieces(), TurnFirst}
}
