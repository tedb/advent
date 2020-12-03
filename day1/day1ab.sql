# Run as:  sqlite3 < day1/day1ab.sql

ATTACH DATABASE ':memory:' AS mem;

# Import the `input` file to `input` table
create table input (x int);
.mode csv input
.import day1/input input

.mode table output
select count(*) as "Imported lines" from input;

select t1.x * t2.x AS "Part 1 result" from input t1, input t2 where t1.x = 2020-t2.x limit 1;

select t1.x * t2.x * t3.x AS "Part 2 result" from input t1, input t2, input t3 where t1.x + t2.x + t3.x = 2020 limit 1;
