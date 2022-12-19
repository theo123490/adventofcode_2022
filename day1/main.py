with open("./input_2.txt") as file:
    lines = [line.rstrip() for line in file]

calorie_bank = 0
elf_list = []

def assign_bank(elf_list, calorie_bank):
    elf_list.append(calorie_bank)
    return

for line in lines:
    if len(line) <= 0 :
        if calorie_bank <= 0 : 
            continue
        assign_bank(elf_list, calorie_bank)
        calorie_bank = 0
    else:
        calorie = int(line)
        calorie_bank += calorie

assign_bank(elf_list, calorie_bank)
print(max(elf_list))