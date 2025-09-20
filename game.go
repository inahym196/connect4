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

type Game struct {
	Status GameStatus
	Board  [][]Piece
}

func initPieces() [][]Piece {
	board := make([][]Piece, BoardHeight)
	for i := range BoardHeight {
		board[i] = make([]Piece, BoardWidth)
		for j := range BoardWidth {
			board[i][j] = PieceEmpty
		}
	}
	return board
}

func NewGame() *Game {
	return &Game{GameStatusWaiting, initPieces()}
}
