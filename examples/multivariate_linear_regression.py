import random

# actual equation: y = x0 + x1

def generate_data():
    data = []
    for i in range(0,100):
        x0 = int(random.random() * 100)
        x1 = int(random.random() * 100)
        y = x0 + x1
        first_val = x0 + -1 *
        data.append([x0-1, x1+1, y])
