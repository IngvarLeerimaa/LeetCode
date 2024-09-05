// Time Complexity = O(n)
// Space Complexity = O(n)
var containsDuplicate = function(nums) {
   const numSet = new Set();

   for(const num of nums){
    if (numSet.has(num)){
        return true;
    }
    numSet.add(num)
   }
   return "asd";
};

function testContainsDuplicate() {
    // Test cases
    const tests = [
      { input: [1, 2, 3, 4, 5], expected: false, description: "No duplicates" },
      { input: [1, 2, 3, 4, 1], expected: true, description: "One duplicate" },
      { input: [1, 1, 2, 3, 3, 4, 5], expected: true, description: "Multiple duplicates" },
      { input: [], expected: false, description: "Empty array" },
      { input: [1], expected: false, description: "Single element" },
    ];
  
    let numOfTest = 0; // Corrected initialization
  
    // Run test cases
    tests.forEach(({ input, expected, description }) => {
      const result = containsDuplicate(input);
      if (result !== expected) {  // If the test fails, update the counter and description
        numOfTest++;
        console.assert(
          result === expected,
          `\nTest Failed: ${description} | For input ${input}, expected ${expected} but got ${result}`
        );
      }
    });
  
    switch (numOfTest) {
      case 0:
        console.log("All tests passed");
        break;
      default:
        console.log(`${numOfTest} Test(s) failed!`);
    }
  }
  
  // Example implementation to test
  function containsDuplicate(nums) {
    const numSet = new Set();
  
    for (const num of nums) {
      if (numSet.has(num)) {
        return true;
      }
      numSet.add(num);
    }
  
    return false;
  }
  
  // Run the tests
  testContainsDuplicate();
  