


import random

def CreatePhoneNumber():
    with open("numbers", "a") as file:
        for x in range(10):
            file.write(str(random.randint(0,9)))
        file.write("\n")

for x in range(5):
    CreatePhoneNumber()


