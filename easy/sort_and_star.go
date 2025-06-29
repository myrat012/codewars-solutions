/*
-- Sort and Star --

Link: https://www.codewars.com/kata/57cfdf34902f6ba3d300001e/go

You will be given a list of strings. You must sort it alphabetically (case-sensitive, and based on the ASCII values of the chars) and then return the first value.
The returned value must be a string, and have "***" between each of its letters.
You should not remove or add elements from/to the array.

Example
s = []string{"bitcoin", "take", "over", "the", "world", "maybe", "who", "knows", "perhaps"}
Expect(TwoSort(s)).To(Equal("b***i***t***c***o***i***n"))
*/

package easy

import (
	"sort"
	"strings"
)

func TwoSort(arr []string) string {
	sort.Strings(arr)
	chars := strings.Split(arr[0], "")
	return strings.Join(chars, "***")
}
