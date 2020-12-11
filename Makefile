.PHONY: get day1 day2 day3 day4 day5

get:
	./get.sh

day1: get
	sqlite3 < day1/day1ab.sql

day2: get
	nim compile --run --verbosity:0 day2/day2.nim

day3: get
	nim compile --run --verbosity:0 day3/day3.nim

day4: get
	ruby day4/day4.rb

day5: get
	ruby day5/day5.rb

day7: get
	ruby day7/day7.rb day7/input

day8: get
	ruby day8/day8.rb day8/input
