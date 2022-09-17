def read_file
  SeatMap.new IO.read('data.txt').lines(chomp: true)
end

class SeatMap
  attr_reader :data

  def initialize(data)
    @data = data
  end

  def width
    @data[0].length
  end

  def height
    @data.length
  end

  # If a seat is empty (L) and there are no occupied seats adjacent to it, the seat becomes occupied.
  # If a seat is occupied (#) and four or more seats adjacent to it are also occupied, the seat becomes empty.
  # Otherwise, the seat's state does not change.
  def simulate_iteration(strategy)
    changes = []
    (0...height).each do |r|
      (0...width).each do |c|
        next unless seat? r,c

        nei_taken = strategy.neighbour_taken(self, r,c)

        if free?(r, c) &&  nei_taken== 0
          changes << [r, c, "#"]
        elsif taken?(r, c) && nei_taken >= strategy.limit
          changes << [r,c, "L"]
        end
      end
    end
    changes
  end

  def count_occupied
    cnt = 0
    (0...height).each do |r|
      (0...width).each do |c|
        cnt += 1 if taken? r, c
      end
    end
    cnt
  end

  def apply_changes(changes)
    changes.each { |c| @data[c[0]][c[1]] = c[2] }
  end

  def taken?(r,c)
    @data[r][c] == "#"
  end

  def free?(r,c)
    @data[r][c] == "L"
  end

  def seat?(r,c)
    free?(r,c) || taken?(r,c)
  end
end

class NeighbourStrategyP1
  #lots of allocations, it slows things down
  def neighbour_taken(seat_map, row, col)
    candidates = [
      Coords.new(row-1, col-1), Coords.new(row-1, col), Coords.new(row-1, col+1),
      Coords.new(row, col-1),                       Coords.new(row, col+1),
      Coords.new(row+1, col-1), Coords.new(row+1, col), Coords.new(row+1, col+1),
    ]

    candidates.filter{ |c| c.row < seat_map.height && c.col < seat_map.width && c.row >= 0 && c.col >= 0 }.count {|c| seat_map.taken? c.row, c.col}
  end

  def limit
    4
  end
end

class NeighbourStrategyP2
  def limit
    5
  end

  def neighbour_taken(seat_map, row, col)
    candidates = []

    (row-1).downto(0).each do |r|
      if seat_map.seat?(r, col)
        candidates << Coords.new(r, col)
        break
      end
    end

    (row+1...seat_map.height).each do |r|
      if seat_map.seat?(r, col)
        candidates << Coords.new(r, col)
        break
      end
    end

    (col-1).downto(0).each do |c|
      if seat_map.seat?(row, c)
        candidates << Coords.new(row,c)
        break
      end
    end

    (col+1...seat_map.width).each do |c|
      if seat_map.seat?(row, c)
        candidates << Coords.new(row,c)
        break
      end
    end

    r,c = row-1,col-1
    #to upper left corner
    while r >= 0 && c >= 0
      if seat_map.seat?(r,c)
        candidates << Coords.new(r,c)
        break
      end
      r-=1
      c-=1
    end

    #to lower right corner
    r,c = row+1,col+1
    while r < seat_map.height && c < seat_map.width
      if seat_map.seat?(r,c)
        candidates << Coords.new(r,c)
        break
      end
      r+=1
      c+=1
    end

    #lower left
    r,c = row+1,col-1
    while r < seat_map.height && c >= 0
      if seat_map.seat?(r,c)
        candidates << Coords.new(r,c)
        break
      end
      r+=1
      c-=1
    end

    #upper right
    r,c = row-1,col+1
    while r >= 0 && c < seat_map.width
      if seat_map.seat?(r,c)
        candidates << Coords.new(r,c)
        break
      end
      r-=1
      c+=1
    end

    candidates.count { |c| seat_map.taken? c.row, c.col}
  end
end

class Coords
  attr_reader :row, :col
  def initialize(row, col)
    @row = row
    @col = col
  end
end

def run(strategy)
  m = read_file
  loop do
    changes = m.simulate_iteration strategy
    return m.count_occupied if changes.empty?
    m.apply_changes changes
  end
  raise 'oops'
end

def p1
  run NeighbourStrategyP1.new
end

def p2
  run NeighbourStrategyP2.new
end

exp_p1 = 2270
exp_p2 = 2042
res_p1 = p1
res_p2 = p2

puts res_p1
puts res_p2
raise 'p1' if exp_p1 != res_p1
raise 'p2' if exp_p2 != res_p2