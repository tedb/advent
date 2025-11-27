# Run as:  sqlite3 < day1/day1ab.sql

ATTACH DATABASE ':memory:' AS mem;

# Import the `input` file to `input` table
create table input (x int);
.mode csv input
.import day1/input input

.mode table output
select count(*) as "Imported lines" from input;

select sum(y) AS part1 from (select x > lag(x, 1) OVER () AS y from input);

select count(*) AS part2 from (select x, sum(x) OVER (rows  between current row and 2 following) AS this, sum(x) OVER (rows between 1 following and 3 following) AS next from input) where next > this;
