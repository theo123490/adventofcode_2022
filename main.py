with open("./input.txt") as file:
    lines = [line.rstrip() for line in file]

class Player:
    def __init__(self, move_dict):
        self.move = ""
        self.possible_move = ["rock", "paper", "scissor"]
        self.move_dict = move_dict

    def validify_move(self, input_string):
        if input_string not in self.possible_move:
            raise ValueError("input string is not included in possible move")
    
self_dict = {
    'A' : 'rock',
    'B' : 'paper',
    'C' : 'scissor'
}

player_self = Player(self_dict)
print(player_self)