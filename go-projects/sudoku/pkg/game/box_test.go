package game

import (
	"testing"
)

func TestBox(t *testing.T) {
    t.Log("Running Box Test") 
    
    // create a new Box
    b := newBox()
    if b.getValue() != emptyBoxValue {
        t.Fatal("Not empty box value on init")
    }
    
    // test get and set
    b.setValue(3)
    if b.getValue() != 3 {
        t.Fatal("Set and/or Get does not return correct value")
    }

    // check guesses are empty on creation
    if b.getGuesses() != nil {
        t.Fatal("Non-empty guess list on creation")
    }

    // set the guesses and get the guesses
    guesses := []int{1, 2, 5, 7}
    
    for _, v := range guesses {
        b.addGuess(v)
    }

    // get the gueses
    ret := b.getGuesses()
    for i, v := range guesses {
        if v != ret[i] {
            t.Fatal("Guesses do not match after set")
        }
    }

    // remove the last element as order is not normally preserved
    b.removeGuess(7)
    guesses = []int{1, 2, 5}
    for i, v := range b.getGuesses() {
        if v != guesses[i] {
            t.Fatal("Guesses do not match after delete")
        }
    }
}

