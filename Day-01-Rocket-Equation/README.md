# Day 1: The Tyranny of the Rocket Equation

Fuel required to launch a given **module** is based on its **mass**. Specifically, to find the fuel required for a 
module, take its mass, divide by three, round down, and subtract 2.

Fuel itself requires fuel just like a module - take its mass, divide by three, round down, and subtract 2. However, that fuel also requires fuel, and that fuel requires fuel, and so on.

For example:
* A module of mass `14` requires `2` fuel. This fuel requires no further fuel (2 divided by 3 and rounded down is `0`, 
which would call for a negative fuel), so the total fuel required is still just `2`.
* At first, a module of mass `1969` requires `654` fuel. Then, this fuel requires `216` more fuel (`654 / 3 - 2`). `216` 
then requires `70` more fuel, which requires `21` fuel, which requires `5` fuel, which requires no further fuel. 
So, the total fuel required for a module of mass `1969` is `654 + 216 + 70 + 21 + 5 = 966`.
* The fuel required by a module of mass `100756` and its fuel is: 
`33583 + 11192 + 3728 + 1240 + 411 + 135 + 43 + 12 + 2 = 50346`.

**Find out the sum of the fuel requirements from input data.**

## Run test

From this path (`advent-code-2019/Day-01-Rocket-Equation`) just:

`make test`

## Run

From this path (`advent-code-2019/Day-01-Rocket-Equation`) just:

`make run`