class TicTacToe
  def initialize
    @board = Array.new(9, ' ')
    @current_player = 'X'
  end

  def print_board
    puts " #{@board[0]} | #{@board[1]} | #{@board[2]}"
    puts "---+---+---"
    puts " #{@board[3]} | #{@board[4]} | #{@board[5]}"
    puts "---+---+---"
    puts " #{@board[6]} | #{@board[7]} | #{@board[8]}"
  end

  def make_move(position)
    if @board[position] == ' '
      @board[position] = @current_player
      @current_player = (@current_player == 'X') ? 'O' : 'X'
      return true
    else
      puts '该位置已经被占用，请选择一个空位置。'
      return false
    end
  end

  def check_winner
    winning_combinations = [
      [0, 1, 2], [3, 4, 5], [6, 7, 8],
      [0, 3, 6], [1, 4, 7], [2, 5, 8],
      [0, 4, 8], [2, 4, 6]
    ]

    winning_combinations.each do |combo|
      a, b, c = combo
      if @board[a] != ' ' && @board[a] == @board[b] && @board[b] == @board[c]
        return @board[a]
      end
    end

    return nil
  end

  def play
    puts '欢迎来到井字游戏！'
    puts '玩家 X 先开始。'

    loop do
      print_board
      puts "玩家 #{@current_player}，请选择一个位置 (0-8): "
      position = gets.chomp.to_i

      if position >= 0 && position <= 8
        if make_move(position)
          winner = check_winner
          if winner
            print_board
            puts "玩家 #{winner} 获胜！恭喜！"
            break
          elsif @board.none? { |cell| cell == ' ' }
            print_board
            puts '平局！'
            break
          end
        end
      else
        puts '请选择有效的位置 (0-8)。'
      end
    end
  end
end

game = TicTacToe.new
game.play
