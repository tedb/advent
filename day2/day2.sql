# Run as:  sqlite3 < day2/day2.sql

ATTACH DATABASE ':memory:' AS mem;

# Import the `input` file to `input` table
create table input (cmd varchar, x int);
.mode csv input
.separator " "
.import day2/input input

.mode table output
select count(*) as "Imported lines" from input;

select sum(x) filter (where cmd = 'forward') * (sum(x) filter (where cmd = 'down') - sum(x) filter (where cmd = 'up')) from input;
