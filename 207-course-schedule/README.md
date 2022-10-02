# Problem
[Leetcode - 207. Course Schedule (Medium)](https://leetcode.com/problems/course-schedule/)

## Takeaway
- Depth first search recursion would be easiest.
- Course ids are incremented from `0` to `numCourses - 1` so we can use an array of arrays instead of a map of arrays.
- We must set a flag to validate for infinite loops.
- We don't need to validate the same course more than once.

## Solution
- First create an array of arrays of course prerequisites `[][]int` of size `numCourses` and iterate through all the `prerequisites` adding them to each course.
- Iterate through each course and validate it and return `false` on the first failure.
  1. Check if the course is currently being validated (infinite loop) if so return `false`.
  2. If the course has no prerequisites then return `true`.
  3. Set a flag to indicate the course is currently being validated.
  4. Iterate through each prerequisites and validate it where we can return `false` on the first failure -> This will recursively repeat step #1.
  5. Remove the flag indicating the course is done validation.
  6. Remove the `prerequisites` for the course so future validations can immediately return `true`.
