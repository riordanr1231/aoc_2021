# frozen_string_literal: true

require 'pry'
require 'pry-byebug'

aim = pos = depth = 0

lines = IO.readlines('input')
lines.each do |l|
  l =~ /^(\w+)\s+(\d+)$/
  raise if $&.nil?

  direction = $1
  magnitude = $2.to_i

  case direction
  when 'forward'
    pos += magnitude
    depth += magnitude * aim
  when 'up'
    aim -= magnitude
  when 'down'
    aim += magnitude
  else
    raise
  end
end

pp [pos, depth, pos*depth]
