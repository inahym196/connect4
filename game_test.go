package connect4_test

import (
	"testing"

	"github.com/inahym196/connect4"
)

func TestNewGame(t *testing.T) {
	g := connect4.NewGame()
	if g == nil {
		t.Fatal("NewGame() returned nil")
	}
	if g.Status != connect4.GameStatusWaiting {
		t.Errorf("expected Status=%d, got %d", connect4.GameStatusWaiting, g.Status)
	}
}
