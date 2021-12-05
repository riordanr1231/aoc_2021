# frozen_string_literal: true

require 'pry'
require 'pry-byebug'

pos = depth = 0

lines = IO.readlines('input')
lines.each do |l|
  l =~ /^(\w+)\s+(\d+)$/
  raise if $&.nil?

  case $1
  when 'forward'
    pos += $2.to_i
  when 'up'
    depth -= $2.to_i
  when 'down'
    depth += $2.to_i
  else
    raise
  end
end

pp [pos, depth, pos*depth]
