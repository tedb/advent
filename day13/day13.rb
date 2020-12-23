ts, ids = ARGF.read.lines
ts = ts.to_i
ids = ids.split(",").reject {|x| x == 'x'}
puts ids.map { |id| id = id.to_i; [id, (id - (ts % id)) ] }.min { |a,b| a[1] <=> b[1] }.reduce(&:*)
