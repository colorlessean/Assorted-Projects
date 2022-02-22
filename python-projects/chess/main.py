from enum import Enum

class Color(Enum):
    WHITE = 1 
    BLACK = 2

COLORS = {
    Color.WHITE: "W",
    Color.BLACK: "B",
}

class Role(Enum):
    KING = 1 
    QUEEN = 2
    BISHOP = 3
    KNIGHT = 4
    ROOK = 5
    PAWN = 6

ROLES = {
    Role.KING: "K",
    Role.QUEEN: "Q",
    Role.BISHOP: "B",
    Role.KNIGHT: "N",
    Role.ROOK: "R",
    Role.PAWN: "P",
}

# Luckily this is the same for black and white as they are mirrored from player perspective
BACK_ROW_ORDER = [
    Role.ROOK,
    Role.KNIGHT,
    Role.BISHOP,
    Role.QUEEN,
    Role.KING,
    Role.BISHOP,
    Role.KNIGHT,
    Role.ROOK
]

class Piece:
    def __init__(self, color, position, role):
        self.color = color # black/white
        self.position = position # ex: a5
        self.role = role

    def __str__(self):
        return f"{self.color}{self.position}"


class Square:
    def __init__(self, piece):
        self.piece = piece # instance of a piece or None

# board will not be a visual represetnation rather a collection of pieces
class Board:
    # setup an empty board
    def __init__(self):
        # initialize a list of pieces for the "board"
        pieces = []

        # initialize each of the pieces for white then black
        for i in range(0, 8):
            boardCode = chr(i + 97) + "1"
            pieces.append(Piece(Color.WHITE, boardCode, BACK_ROW_ORDER[i]))
        for i in range(0, 8):
            boardCode = chr(i + 97) + "2"
            pieces.append(Piece(Color.WHITE, boardCode, Role.PAWN))

        for i in range(0, 8):
            boardCode = chr(i + 97) + "8"
            pieces.append(Piece(Color.BLACK, boardCode, BACK_ROW_ORDER[i]))
        for i in range(0, 8):
            boardCode = chr(i + 97) + "7"
            pieces.append(Piece(Color.BLACK, boardCode, Role.PAWN))

        self.pieces = pieces

    def __str__(self):
        # lets get a blank board going first
        rows, cols = (8, 8)
        arr = [["  " for i in range(cols)] for j in range(rows)]

        for piece in self.pieces:
            indexX = 8 - int(piece.position[1])
            indexY = ord(piece.position[0]) - 97

            arr[indexX][indexY] = COLORS[piece.color] + ROLES[piece.role] 

        string = "+--+--+--+--+--+--+--+--+\n"

        for i in range(0, 8):
            temp = "|"
            for j in range(0, 8):
                temp += arr[i][j]
                temp += "|"
            string += temp
            string += "\n+--+--+--+--+--+--+--+--+\n"

        return string

def main():
    # initalize a board / game
    board = Board()
    
    # game will probably be in a super loop that will end with a checkmate or withdrawl
    print(board)
    

    print("Success")

if __name__ == "__main__":
    main()
