package game

import (
    "fmt"
)

type Board struct {
    boxes [9][9]Box
}

func newBoard(values [9][9]int) Board {
    b := Board{}
    for i, y := range values {
        for j, x := range y {
            b.boxes[i][j] = newBox()
            b.boxes[i][j].setValue(x)
        }
    }
    return b
}

// print the board with values not guesses
func (b *Board) printBoard() {
    for i, row := range b.boxes {
        if i == 0 {
            fmt.Println("┌─┬─┬─┬─┬─┬─┬─┬─┬─┐")
        } else {
            fmt.Println("├─┼─┼─┼─┼─┼─┼─┼─┼─┤")
        }
        fmt.Print("│")
        for _, box := range row {
            if box.getValue() == emptyBoxValue {
                fmt.Print(" │")
            } else {
                fmt.Print(box.getValue(), "│") 
            }
        }
        fmt.Println()
    }
    fmt.Println("└─┴─┴─┴─┴─┴─┴─┴─┴─┘")
}

