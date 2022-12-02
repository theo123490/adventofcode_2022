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

    def rps_logic(self, current_move, enemy_move):
        if current_move == enemy_move :
            return 'draw'
        if current_move.win_to == enemy_move.name :
            return 'win'
        if current_move.lose_to == enemy_move.name :
            return 'lose'

class Move:
    def __init__(self, name, win_to, lose_to):
        self.name = name
        self.win_to = win_to
        self.lose_to = lose_to

rock = Move("rock", "scissor", "paper")
paper = Move("paper", "rock", "scissor")
scissor = Move("scissor", "paper", "rock")

self_dict = {
    'A' : 'rock',
    'B' : 'paper',
    'C' : 'scissor'
}

player_self = Player(self_dict)
