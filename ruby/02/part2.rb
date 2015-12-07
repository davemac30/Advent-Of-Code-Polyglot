#!/usr/bin/env ruby

puts File.readlines('input').reduce(0) { |s, l|
  h, w, d = l.chomp.split(/x/).map{|e| e.to_i}.sort;
  s + h + h + w + w + h * w * d
}
