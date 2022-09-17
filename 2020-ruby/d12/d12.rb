def read_file
  IO.read('data.txt').lines(chomp: true).map do |line|
    {cmd: line[0], val: line.slice(1, line.length-1).to_i}
  end
end

class State
  attr_reader :x, :y

  def initialize
    @x = 0
    @y = 0
    @direction = 'E'
  end

  def call(instruction)
    cmd = instruction[:cmd]
    val = instruction[:val]

    case cmd
    when 'N', 'E', 'S', 'W' then mov cmd, val
    when 'L', 'R' then deg_to_v cmd, val
    when 'F' then mov @direction, val
    else raise "unknown command: #{cmd}"
    end
  end

  def deg_to_v(dir, val)
    @direction = rotate(@direction, dir, val)
  end

  def mov(dir, val)
    case dir
    when 'N' then @y -= val
    when 'E' then @x += val
    when 'S' then @y += val
    when 'W' then @x -= val
    end
  end
end

def rotate(current_dir, dir, val)
  dir_vals = {
    'E' => 0,
    'S' => 1,
    'W' => 2,
    'N' => 3
  }.freeze

  reverse_map = {}
  dir_vals.each { |key, v| reverse_map[v] = key }

  normalized = ((val % 360) / 90).floor
  normalized *= -1 if dir == 'L'

  new_dir = (dir_vals[current_dir] + normalized) % dir_vals.length
  reverse_map[new_dir]
end

def p1
  data = read_file
  s = State.new
  data.each { |instr| s.call instr }
  manhattan s.x, s.y
end

def manhattan(x,y)
  x.abs + y.abs
end

def p2
  data = read_file
  s = StateP2.new
  data.each { |instr| s.call instr }
  manhattan s.x, s.y
end

puts p1
puts p2

exp_p1 = 1186
exp_p2 = -1

raise 'p1' if exp_p1 != p1
raise 'p2' if exp_p2 != p2