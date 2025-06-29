/*
-- Count of positives / sum of negatives --

Link: https://www.codewars.com/kata/576bb71bbbcf0951d5000044/train/go

Given an array of integers.
Return an array, where the first element is the count of positives numbers and the second element is sum of negative numbers. 0 is neither positive nor negative.
If the input is an empty array or is null, return an empty array.

Example
For input [1, 2, 3, 4, 5, 6, 7, 8, 9, 10, -11, -12, -13, -14, -15], you should return [10, -65].
*/
package easy

func CountPositivesSumNegatives(numbers []int) []int {
	res := []int{0, 0}
	for _, v := range numbers {
		if v > 0 {
			res[0] += 1
		} else {
			res[1] += v
		}
	}

	return res // your code here
}
