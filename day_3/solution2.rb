# frozen_string_literal: true

require 'pry'
require 'pry-byebug'

class Node
  attr_accessor :zero, :one, :count, :type

  def initialize(zero: nil, one: nil, count: 0, type: nil)
    @count = count
    @zero = zero
    @one = one
    @type = type
  end

  def create_branch(type)
    if type.eql?('0')
      target ||= @zero ||= Node.new(type: '0')
    elsif type.eql?('1')
      target ||= @one ||= Node.new(type: '1')
    else
      raise
    end

    target.count += 1
    target
  end

  def next(type = :generator)
    return nil if self.zero.nil? and self.one.nil?
    return self.one if self.zero.nil?
    return self.zero if self.one.nil?

    if type.eql?(:generator)
      return self.zero if self.zero.count > self.one.count
      return self.one
    else
      return self.zero if self.zero.count <= self.one.count
      return self.one
    end
  end

  def generator
    buf = ''
    head = self

    while head = head.next(:generator)
      buf += head.type
    end
    buf
  end

  def scrubber
    buf = ''
    head = self

    while head = head.next(:scrubber)
      buf += head.type
    end
    buf
  end

  def zero_count
    zero&.count || 0
  end

  def one_count
    one&.count || 0
  end
end

root = Node.new
IO.readlines('input').each do |l|
  root.count += 1
  head = root
  l.chomp!
  l.each_char do |c|
    head = head.create_branch(c)
  end
end

puts root.generator.to_i(2),
     root.scrubber.to_i(2),
     root.generator.to_i(2) * root.scrubber.to_i(2)
