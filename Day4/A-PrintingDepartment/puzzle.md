# Day 4: Printing Department
* Input representing a grid
* Cells can either be empty `.` or have a roll of paper `@`
* Identify rolls of paper where there are < 4 adjacent rolls
  * Including diagonally adjacent

## Notes
### Could I use Bitwise operators?
* **Maybe, but too hard for tonight :D**
* Input
```
001
101
101

```
* Looking at the middle character
* For each row, we can bitshift left and right
* For each bitshifted row, plus the original row, we do a bitwise Or

If I want to take the previous input and pretend we're looking for slots which themselves are 1 with less than 2 adjacent 1s

```
000
001
000
```

New example

```
11001
11111
00110
```

=>

```
11000
11111
    00110
```

## Plan
1. Read in the file line by line
  - Load previous and next lines as well, we only need adjacent lines
  - This is the input to our function, a mini grid
2.k
3. For each roll of paper in the line, count adja

## Full Text
--- Day 4: Printing Department ---

You ride the escalator down to the printing department. They're clearly getting ready for Christmas; they have lots of large rolls of paper everywhere, and there's even a massive printer in the corner (to handle the really big print jobs).

Decorating here will be easy: they can make their own decorations. What you really need is a way to get further into the North Pole base while the elevators are offline.

"Actually, maybe we can help with that," one of the Elves replies when you ask for help. "We're pretty sure there's a cafeteria on the other side of the back wall. If we could break through the wall, you'd be able to keep moving. It's too bad all of our forklifts are so busy moving those big rolls of paper around."

If you can optimize the work the forklifts are doing, maybe they would have time to spare to break through the wall.

The rolls of paper (@) are arranged on a large grid; the Elves even have a helpful diagram (your puzzle input) indicating where everything is located.

For example:

..@@.@@@@.
@@@.@.@.@@
@@@@@.@.@@
@.@@@@..@.
@@.@@@@.@@
.@@@@@@@.@
.@.@.@.@@@
@.@@@.@@@@
.@@@@@@@@.
@.@.@@@.@.

The forklifts can only access a roll of paper if there are fewer than four rolls of paper in the eight adjacent positions. If you can figure out which rolls of paper the forklifts can access, they'll spend less time looking and more time breaking down the wall to the cafeteria.

In this example, there are 13 rolls of paper that can be accessed by a forklift (marked with x):

..xx.xx@x.
x@@.@.@.@@
@@@@@.x.@@
@.@@@@..@.
x@.@@@@.@x
.@@@@@@@.@
.@.@.@.@@@
x.@@@.@@@@
.@@@@@@@@.
x.x.@@@.x.

Consider your complete diagram of the paper roll locations. How many rolls of paper can be accessed by a forklift?
