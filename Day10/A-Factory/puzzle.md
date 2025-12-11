# Day 10: Factory
* Segment in brackets is the machine lights.
    * It represents the final state we need the same number of lights to be in.
    * `.` is off `#` is on
* Segments in parens are buttons.
    * Each number represents the index of the light that button toggles.
* Segments in curlies is joltage and we don't need it yet
    * Will make my life suck later

## Plan
* Parse machine state to binary.
    * Handling leading 0s?
```
.##. => 0110
```
* Parse each button to a binary number.
```
(0,3,4) => 10011
```
* Increment from 0...n until the right number of buttons to press is found
* Test if a button combination works by using a running "XOR sum"

```
    0000 <- Starting state of machine
XOR 1010
    1010
XOR 1100
    0110 <- Desired state of machine
```
## Notes
* I can know the number I need to XOR with for any given input
```
1010 <- Current
1001 <- Needed
0011 <- Final
```
* What operations can solve for "Needed"?
```
    1010
??  0011
    1001
```
* Is it just another XOR
* Can I truly "Sum" XORs?
    * (spoiler, no)
    * I have to add them consecutively
    * TODO: Technically in this example, I could get desired output using AND & NOT with the two separate XOR solutions. Might want to see if that continues
```
A   1011
XOR 1101
    0110

B   1111
XOR 0001
    1110

AwB 0110
XOR 1110
    1000

ALL 1011
XOR 1101
    0110
XOR 1111
    1001
XOR 1110
    0111

```
## Full Puzzle

Just across the hall, you find a large factory. Fortunately, the Elves here have plenty of time to decorate. Unfortunately, it's because the factory machines are all offline, and none of the Elves can figure out the initialization procedure.

The Elves do have the manual for the machines, but the section detailing the initialization procedure was eaten by a Shiba Inu. All that remains of the manual are some indicator light diagrams, button wiring schematics, and joltage requirements for each machine.

For example:

[.##.] (3) (1,3) (2) (2,3) (0,2) (0,1) {3,5,4,7}
[...#.] (0,2,3,4) (2,3) (0,4) (0,1,2) (1,2,3,4) {7,5,12,7,2}
[.###.#] (0,1,2,3,4) (0,3,4) (0,1,2,4,5) (1,2) {10,11,11,5,10,5}

The manual describes one machine per line. Each line contains a single indicator light diagram in [square brackets], one or more button wiring schematics in (parentheses), and joltage requirements in {curly braces}.

To start a machine, its indicator lights must match those shown in the diagram, where . means off and # means on. The machine has the number of indicator lights shown, but its indicator lights are all initially off.

So, an indicator light diagram like [.##.] means that the machine has four indicator lights which are initially off and that the goal is to simultaneously configure the first light to be off, the second light to be on, the third to be on, and the fourth to be off.

You can toggle the state of indicator lights by pushing any of the listed buttons. Each button lists which indicator lights it toggles, where 0 means the first light, 1 means the second light, and so on. When you push a button, each listed indicator light either turns on (if it was off) or turns off (if it was on). You have to push each button an integer number of times; there's no such thing as "0.5 presses" (nor can you push a button a negative number of times).

So, a button wiring schematic like (0,3,4) means that each time you push that button, the first, fourth, and fifth indicator lights would all toggle between on and off. If the indicator lights were [#.....], pushing the button would change them to be [...##.] instead.

Because none of the machines are running, the joltage requirements are irrelevant and can be safely ignored.

You can push each button as many times as you like. However, to save on time, you will need to determine the fewest total presses required to correctly configure all indicator lights for all machines in your list.

There are a few ways to correctly configure the first machine:

[.##.] (3) (1,3) (2) (2,3) (0,2) (0,1) {3,5,4,7}

    You could press the first three buttons once each, a total of 3 button presses.
    You could press (1,3) once, (2,3) once, and (0,1) twice, a total of 4 button presses.
    You could press all of the buttons except (1,3) once each, a total of 5 button presses.

However, the fewest button presses required is 2. One way to do this is by pressing the last two buttons ((0,2) and (0,1)) once each.

The second machine can be configured with as few as 3 button presses:

[...#.] (0,2,3,4) (2,3) (0,4) (0,1,2) (1,2,3,4) {7,5,12,7,2}

One way to achieve this is by pressing the last three buttons ((0,4), (0,1,2), and (1,2,3,4)) once each.

The third machine has a total of six indicator lights that need to be configured correctly:

[.###.#] (0,1,2,3,4) (0,3,4) (0,1,2,4,5) (1,2) {10,11,11,5,10,5}

The fewest presses required to correctly configure it is 2; one way to do this is by pressing buttons (0,3,4) and (0,1,2,4,5) once each.

So, the fewest button presses required to correctly configure the indicator lights on all of the machines is 2 + 3 + 2 = 7.

Analyze each machine's indicator light diagram and button wiring schematics. What is the fewest button presses required to correctly configure the indicator lights on all of the machines?
