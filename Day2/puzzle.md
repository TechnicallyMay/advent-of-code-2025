# Day 2: Gift Shop

## Summary
* We get a set of number ranges (e.g., 95-115 meaning 95 through 115 inclusive)
* For each number in the given ranges, determine if it is  a sequence of digits repeated twice (5454 is 54 twice)
* For each number that matches that criteria, sum it up. The answer is the final sum

## Notes
* There may be some numerical pattern that allows us to math it out
* Given a range, we could take the starting number and distance to the end, then calculate the upcoming repeats.
  * If our range is 10-20, we know the next possible repeat from the start is 11. Since We have a distance of 10, we know we can add 1 to reach the repeat.
  * How to determine the upcoming repeats?
  * Probably easiest with text so we can split the string in half.
  * Once we run out of possible repeats with the starting digits, how to increment to the next
    * For example, given 995 we know there are no more possible repeats. How to get to the next number of digits (i.e., get to 1000)
    * Is it just do "1" + ("0" * len("995")) ?
    * Does this work for smaller required increments? 9599 We need to go up to 9600.
    * Do we need to add enough  to increment the first half of the number?
* While a repeat is possible, increment?

* Simple brute force method would be to iterate over each number in the range and check.

### Finding a pattern
| Starting Val | Next Wrap | Required Increment | Next Number to Start Testing | Increment to that thing |
|--------------|-----------|--------------------|------------------------------|-------------------------|
| 9599         | 9696      | 97                 | 9600                         | 97                      |
| 55           | 66        | 11                 | 60                           | 5                       |
| 101222       | 102102    | 880                | 102000                       | 778                     |
| 333999       | 334334    | 335                | 334000                       | 1                       |

Can't find an obvious pattern. How to calculate "next number to start testing?

### Split into 2 idea
* What if we split into 2 numbers? e.x., 333999 => 333 999. Then we know we just need to increment 333 by 1 and set 999 to zero

## Initial Plan

1. Split the smaller number into 2 parts
    - If number of digits in smaller number is not even, skip to 4
2. Determine if it's possible to have a match
    - If 2nd part <= 1st part
3. If possible, add the required amount to the 2nd part
    - Check we're still <= max number of the range
    - No need to repeat - if we find a match that'll be the only possible one
4. Increment first part by 1. Set second part to zero
    - Check if we're still below our range
5. Repeat

## Full Text
You get inside and take the elevator to its only other stop: the gift shop. "Thank you for visiting the North Pole!" gleefully exclaims a nearby sign. You aren't sure who is even allowed to visit the North Pole, but you know you can access the lobby through here, and from there you can access the rest of the North Pole base.

As you make your way through the surprisingly extensive selection, one of the clerks recognizes you and asks for your help.

As it turns out, one of the younger Elves was playing on a gift shop computer and managed to add a whole bunch of invalid product IDs to their gift shop database! Surely, it would be no trouble for you to identify the invalid product IDs for them, right?

They've even checked most of the product ID ranges already; they only have a few product ID ranges (your puzzle input) that you'll need to check. For example:

11-22,95-115,998-1012,1188511880-1188511890,222220-222224,
1698522-1698528,446443-446449,38593856-38593862,565653-565659,
824824821-824824827,2121212118-2121212124

(The ID ranges are wrapped here for legibility; in your input, they appear on a single long line.)

The ranges are separated by commas (,); each range gives its first ID and last ID separated by a dash (-).

Since the young Elf was just doing silly patterns, you can find the invalid IDs by looking for any ID which is made only of some sequence of digits repeated twice. So, 55 (5 twice), 6464 (64 twice), and 123123 (123 twice) would all be invalid IDs.

None of the numbers have leading zeroes; 0101 isn't an ID at all. (101 is a valid ID that you would ignore.)

Your job is to find all of the invalid IDs that appear in the given ranges. In the above example:

    11-22 has two invalid IDs, 11 and 22.
    95-115 has one invalid ID, 99.
    998-1012 has one invalid ID, 1010.
    1188511880-1188511890 has one invalid ID, 1188511885.
    222220-222224 has one invalid ID, 222222.
    1698522-1698528 contains no invalid IDs.
    446443-446449 has one invalid ID, 446446.
    38593856-38593862 has one invalid ID, 38593859.
    The rest of the ranges contain no invalid IDs.

Adding up all the invalid IDs in this example produces 1227775554.

What do you get if you add up all of the invalid IDs?
