with open("./input_2.txt") as file:
    lines = [line.rstrip() for line in file]

class Player:
    def __init__(self, move_dict, point_dict):
        self.move = ""
        self.possible_move = ["rock", "paper", "scissor"]
        self.move_dict = move_dict
        self.point = 0
        self.point_dict = point_dict

    def translate_move(self, data_move):
        self.move = self.move_dict[data_move]

    def rps_logic(self, move, enemy_move):
        self.move = move
        self.enemy_move = enemy_move
        if self.move == enemy_move :
            self.rps_status = 'draw'
        if self.move.win_to == enemy_move.name :
            self.rps_status = 'win'
        if self.move.lose_to == enemy_move.name :
            self.rps_status = 'lose'

    def calculate_point(self):
        status_point = point_dict[self.rps_status]
        move_point = point_dict[self.move.name]
        self.point += status_point
        self.point += move_point

    def execute_rps(self, enemy_move):
        self.rps_logic(self.move, enemy_move)
        self.calculate_point()

class Move:
    def __init__(self, name, win_to, lose_to):
        self.name = name
        self.win_to = win_to
        self.lose_to = lose_to

rock = Move("rock", "scissor", "paper")
paper = Move("paper", "rock", "scissor")
scissor = Move("scissor", "paper", "rock")

point_dict = {
    'rock' : 1 ,
    'paper' : 2 ,
    'scissor' : 3 ,
    'lose' : 0 ,
    'draw' : 3 ,
    'win' : 6
}

enemy_dict = {
    'A' : rock,
    'B' : paper,
    'C' : scissor
}

self_dict = {
    'X' : rock,
    'Y' : paper,
    'Z' : scissor
}

player_self = Player(self_dict, point_dict)
player_enemy = Player(enemy_dict, point_dict)

for line in lines :
    enemy_move_data, self_move_data = line.split(' ')
    player_enemy.translate_move(enemy_move_data)
    player_self.translate_move(self_move_data)
    player_self.execute_rps(player_enemy.move)

print(player_self.point)