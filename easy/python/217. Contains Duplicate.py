from typing import List

class Solution:
    def containsDuplicate(self, nums: List[int]) -> bool:
        num_set = set()
        
        for num in nums:
            if num in num_set:
                return True
            num_set.add(num)
            
        return False
    
def test_contains_duplicate():
    # Test cases
    tests = [
        {"input": [1, 2, 3, 4, 5], "expected": False, "description": "No duplicates"},
        {"input": [1, 2, 3, 4, 1], "expected": True, "description": "One duplicate"},
        {"input": [1, 1, 2, 3, 3, 4, 5], "expected": True, "description": "Multiple duplicates"},
        {"input": [], "expected": False, "description": "Empty list"},
        {"input": [1], "expected": False, "description": "Single element"},
    ]

    # Initialize variables for counting failed tests
    num_of_failed_tests = 0

    # Create an instance of the Solution class
    solution = Solution()

    # Run test cases
    for test in tests:
        input_data = test["input"]
        expected = test["expected"]
        description = test["description"]

        # Call the containsDuplicate method
        result = solution.containsDuplicate(input_data)
        if result != expected:
            num_of_failed_tests += 1
            print(f"Test Failed: {description} | For input {input_data}, expected {expected} but got {result}")

    # Print the result summary
    print("Test Ended")
    if num_of_failed_tests == 0:
        print("All tests passed")
    else:
        print(f"{num_of_failed_tests} Test(s) failed!")


# Run the tests
test_contains_duplicate()
