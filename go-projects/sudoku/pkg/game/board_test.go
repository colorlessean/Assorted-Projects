package game

import "testing"

func TestBoard(t *testing.T) {
    t.Log("Running Board Test")

    boardValues := [9][9]int{}

    b := newBoard(boardValues)

    b.printBoard()
}
