# Day 3: Crossed Wires
Opening the front panel reveals a jumble of wires. Specifically, two wires are connected to a central port and extend 
outward on a grid. You trace the path each wire takes as it leaves the central port, one wire per line of text 
(your puzzle input).

## Part 1

The wires twist and turn, but the two wires occasionally cross paths. To fix the circuit, you need to find the 
intersection point closest to the central port. Because the wires are on a grid, use the Manhattan distance for this
measurement. While the wires do technically cross right at the central port where they both start, this point does not 
count, nor does a wire count as crossing with itself.

For example, if the first wire's path is `R8,U5,L5,D3`, then starting from the central port (o), it goes right 
`8`, up `5`, left `5`, and finally down `3`:

...........\
...........\
...........\
....+----+.\
....|....|.\
....|....|.\
....|....|.\
.........|.\
.o-------+.\
...........\
Then, if the second wire's path is `U7,R6,D4,L4`, it goes up `7`, right `6`, down `4`, and left `4`:

...........\
.+-----+...\
.|.....|...\
.|..+--X-+.\
.|..|..|.|.\
.|.-X--+.|.\
.|..|....|.\
.|.......|.\
.o-------+.\
...........\
These wires cross at two locations (marked X), but the lower-left one is closer to the central port: its distance is 
`3` + `3` = `6`.

## Part 2
It turns out that this circuit is very timing-sensitive; you actually need to minimize the signal delay.

To do this, calculate the number of steps each wire takes to reach each intersection; choose the intersection where the 
sum of both wires' steps is lowest. If a wire visits a position on the grid multiple times, use the steps value from the 
first time it visits that position when calculating the total value of a specific intersection.

In the above example, the intersection closest to the central port is reached after `8 + 5 + 5 + 2 = 20` steps by the 
first wire and `7 + 6 + 4 + 3 = 20` steps by the second wire for a total of `20 + 20 = 40` steps.

## Run test

From this path (`advent-code-2019/Day-XX-`) just:

`make test`

## Run

From this path (`advent-code-2019/Day-XX-`) just:

`make run`