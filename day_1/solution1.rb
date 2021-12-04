# frozen_string_literal: true

require 'pry'
require 'pry-byebug'

count = 0
lines = IO.readlines('input')
lines.each_cons(2) do |l|
  count += 1 if l.first.to_i < l.last.to_i
end

puts count
