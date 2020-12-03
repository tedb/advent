.PHONY: get day1 day2

get:
	./get.sh

day1: get
	sqlite3 < day1/day1ab.sql

day2: get
	nim compile --run --verbosity:0 day2/day2.nim

day3: get
	nim compile --run --verbosity:0 day3/day3.nim