import random

def generate_sudoku():
    # Create an empty grid
    grid = [[0 for _ in range(9)] for _ in range(9)]

    # Fill in numbers randomly
    def is_valid(num, row, col):
        # Check if the number already exists in the row, column, or region
        for i in range(9):
            if grid[row][i] == num or grid[i][col] == num:
                return False
        region_row = row // 3 * 3
        region_col = col // 3 * 3
        for i in range(3):
            for j in range(3):
                if grid[region_row + i][region_col + j] == num:
                    return False
        return True

    def fill_grid():
        for row in range(9):
            for col in range(9):
                if grid[row][col] == 0:
                    numbers = list(range(1, 10))
                    random.shuffle(numbers)
                    for num in numbers:
                        if is_valid(num, row, col):
                            grid[row][col] = num
                            if fill_grid():
                                return True
                            grid[row][col] = 0
                    return False
        return True

    fill_grid()
    print('Answer:')
    for row in grid:
        print(''.join(str(x) if x != 0 else '.' for x in row), end='')
    print('')

    # Remove numbers to create the puzzle
    for row in range(9):
        for col in range(9):
            if random.random() < 0.7:  # adjust this value to control difficulty
                grid[row][col] = 0

    return grid

# Example usage:
sudoku_grid = generate_sudoku()
for row in sudoku_grid:
    print(''.join(str(x) if x != 0 else '.' for x in row), end='')

print('')