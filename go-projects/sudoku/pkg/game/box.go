package game

const (
    emptyBoxValue int = 0 
)

type Box struct {
    value int
    guesses []int
}

// function to initialize and return a new box instance
func newBox() Box {
    box := Box{
        value: emptyBoxValue,
        guesses: nil,
    }
    return box
}

// get the value that is set in the box
func (b *Box) getValue() int {
    return b.value
}

// set the value in the box (should only be done once)
func (b *Box) setValue(value int) {
   if b.value != emptyBoxValue {
       return
   }
   b.value = value
}

// get the guesses stored in the struct
func (b *Box) getGuesses() []int {
    return b.guesses 
}

// addGuess if guess does not exist in guess list will add to value to it
func (b *Box) addGuess(guess int) {
    for _, v := range b.guesses {
        if v == guess {
            return
        }
    }
    b.guesses = append(b.guesses, guess)
}

// removeGuess if guess is in list remove it order not preserved
func (b *Box) removeGuess(guess int) {
    for i, v := range b.guesses {
        if v == guess {
            b.guesses[i] = b.guesses[len(b.guesses)-1]
            b.guesses = b.guesses[:len(b.guesses)-1]
        }
    }
}

