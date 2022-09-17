def read_file
  d = IO.readlines('data.txt').map(&:to_i)
  d << 0
  d.sort!

  d << d[-1]+3
end

def p1
  adapters = read_file
  # one_jumps = 0
  # three_jumps = 0
  # adapters.slice(0, adapters.length-1).zip(adapters.slice(1, adapters.length - 1)).each do |pair|
  #   diff = pair[1] - pair[0]
  #
  #   one_jumps += 1 if diff == 1
  #   three_jumps += 1 if diff == 3
  # end
  # one_jumps * three_jumps

  diffs = adapters.slice(0, adapters.length-1).zip(adapters.slice(1, adapters.length - 1)).
    map { |pair| pair[1] - pair[0] }.
    group_by { |x| x }  #.group_by(&:itself)
  diffs[3].length * diffs[1].length
end

class Memoized
  def initialize(adapters)
    @adapters = adapters
    @cache = {}
  end

  def solve(idx)
    return @cache[idx] if @cache.has_key? idx
    return 1 if idx >= @adapters.length-1

    res = 0
    #take this and next one - if there's still less <= 3 -> it's ok. Memoize so we can finish fast
    (idx+1...@adapters.length).each do |i|
      res += solve i if @adapters[i] - @adapters[idx] <= 3
    end
    @cache[idx] = res
    res
  end
end

def p2
  data = read_file
  m = Memoized.new data
  m.solve 0
end
exp_p1 = 1656
exp_p2 = 56_693_912_375_296

p1_res = p1
p2_res = p2
puts p1_res
puts p2_res

raise 'p1' if exp_p1 != p1_res
raise 'p2' if exp_p2 != p2_res
puts 'all ok'