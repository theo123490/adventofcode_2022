with open("./input.txt") as file:
    lines = [line.rstrip() for line in file]

class Player:
    def __init__(self, move_dict):
        self.move = ""
        self.possible_move = ["rock", "paper", "scissor"]
        self.move_dict = move_dict

    def translate_move(self, data_move):
        self.move = self.move_dict[data_move]
        return self.move

self_dict = {
    'A' : 'rock',
    'B' : 'paper',
    'C' : 'scissor'
}

player_self = Player(self_dict)
