class TicTacToe:
    def __init__(self):
        self.board = [' ' for _ in range(9)]
        self.current_player = 'X'

    def print_board(self):
        print(f" {self.board[0]} | {self.board[1]} | {self.board[2]} ")
        print("---+---+---")
        print(f" {self.board[3]} | {self.board[4]} | {self.board[5]} ")
        print("---+---+---")
        print(f" {self.board[6]} | {self.board[7]} | {self.board[8]} ")

    def make_move(self, position):
        if self.board[position] == ' ':
            self.board[position] = self.current_player
            self.current_player = 'X' if self.current_player == 'O' else 'O'
            return True
        else:
            print("该位置已经被占用，请选择一个空位置。")
            return False

    def check_winner(self):
        winning_combinations = [
            [0, 1, 2], [3, 4, 5], [6, 7, 8],
            [0, 3, 6], [1, 4, 7], [2, 5, 8],
            [0, 4, 8], [2, 4, 6]
        ]

        for combo in winning_combinations:
            a, b, c = combo
            if (
                self.board[a] != ' ' and
                self.board[a] == self.board[b] == self.board[c]
            ):
                return self.board[a]

        return None

    def play(self):
        print("欢迎来到井字游戏！")
        print("玩家 X 先开始。")

        while True:
            self.print_board()
            try:
                position = int(input(f"玩家 {self.current_player}，请选择一个位置 (0-8): "))
                if 0 <= position <= 8:
                    if self.make_move(position):
                        winner = self.check_winner()
                        if winner:
                            self.print_board()
                            print(f"玩家 {winner} 获胜！恭喜！")
                            break
                        elif ' ' not in self.board:
                            self.print_board()
                            print("平局！")
                            break
                else:
                    print("请选择有效的位置 (0-8)。")
            except ValueError:
                print("请输入一个有效的数字。")

if __name__ == "__main__":
    game = TicTacToe()
    game.play()
