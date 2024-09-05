package ArraysHashing

// Time Complexity = O(n) because it goes over the numbers once
// Space Complexity = O(n) because it needs to store extra to store unique elements of the input array.
// Much faster but takes up more space.
func containsDuplicate(nums []int) bool {
	numSet := make(map[int]bool)
	for _, num := range nums {
		if numSet[num] {
			return true
		}
		numSet[num] = true
	}
	return false
}

//First try
// Time Complexity = O(n2) because grows quadratically with input size.
// Space Complexity = O(1)
// Time is very slow, Space good because it doesn't change
/* func containsDuplicate(nums []int) bool {

for i := 0; i < len(nums)-1; i++ {
	for x := i + 1; x < len(nums); x++ {
		if nums[i] == nums[x] {
				return true
			}
		}

		}
		return false
		} */

/* func main() {
	numsSetOne := []int{1, 2, 3, 1}
	numsSetTwo := []int{1, 2, 3, 4}
	numsSetThree := []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}

	one := containsDuplicate(numsSetOne)
	two := containsDuplicate(numsSetTwo)
	three := containsDuplicate(numsSetThree)

	fmt.Println(one)
	fmt.Println(two)
	fmt.Println(three)
} */
