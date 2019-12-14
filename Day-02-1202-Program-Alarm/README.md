# Day 2: 1202 Program Alarm

An Intcode program is a list of integers separated by commas (like 1, 0, 0, 3, 99). To run one, start by looking at the 
first integer (called position 0). Here, you will find an opcode - either `1`, `2`, or `99`. The opcode indicates what 
to do; for example, 99 means that the program is finished and should immediately halt.

Opcode `1` adds together numbers read from two positions and stores the result in a third position

Opcode `2` works exactly like opcode `1`, except it multiplies the two inputs instead of adding them. Again, the three 
integers after the opcode indicate where the inputs and outputs are, not their values.

For example, suppose you have the following program:

`1, 9, 10, 3, 2, 3, 11, 0, 99, 30, 40, 50`

Result:

`3500, 9, 10, 70, 2, 3, 11, 0, 99, 30, 40, 50`

## Run test

From this path (`advent-code-2019/Day-02-1202-Program-Alarm`) just:

`make test`

## Run

From this path (`advent-code-2019/Day-02-1202-Program-Alarm`) just:

`make run`