# frozen_string_literal: true

require 'pry'
require 'pry-byebug'

a = Array.new(12, 0)

lines = IO.readlines('input')
lines.each do |l|
  l.each_char.with_index { |c, i| a[i] += 1 if c.eql?('1') }
end

gamma = a.collect { |c| c > lines.length / 2 ? '1' : '0'}.join
epsilon = a.collect { |c| c > lines.length / 2 ? '0' : '1'}.join

puts gamma.to_i(2) * epsilon.to_i(2)
