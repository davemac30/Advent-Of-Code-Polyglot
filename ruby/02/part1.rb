#!/usr/bin/env ruby

puts File.readlines('input').reduce(0) { |s, l|
  h, w, d = l.chomp.split(/x/).map{|e| e.to_i}.sort
  s + 3 * h * w + 2 * h * d + 2 * w * d
}
