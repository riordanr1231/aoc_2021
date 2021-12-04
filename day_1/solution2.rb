# frozen_string_literal: true

require 'pry'
require 'pry-byebug'

prev = nil
count = 0
lines = IO.readlines('input')
lines.each_cons(3) do |l|
  curr = l[0].to_i + l[1].to_i + l[2].to_i
  count += 1 if not prev.nil? and curr > prev
  prev = curr
end

puts count
