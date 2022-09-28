# Problem
[Leetcode - 295. Find Median from Data Stream (Hard)](https://leetcode.com/problems/find-median-from-data-stream/)

## Takeaway
- If time complexity wasn't important we could sort the numbers and easily return the median, but this will need to be O(1) when calling `FindMedian`.
- The heap Priority Queue will be required (using the `IntHeap` template from [golang docs](https://pkg.go.dev/container/heap))
- Keep track of the leftmost and rightmost values split in half. The highest value in `left` should be less than or equal to the lowest value in `right`.
- We need the first element of the `left` heap (`index[0]`) to be the max value so we'll need to push values as negative - the alternative option is to create a 2nd `IntHeap` where `Less` returns the opposite.

## Solution
1. Add to the `right` heap.
    - If the highest value in `left` is lower than `num`, then we pop from `left` and add to `right`, then we add `num` to `left`.
    - Otherwise add `num` to `right`.
2. Next we need to even out the `left` heap.
    - If the lowest value in `right` is higher than `num`, then we pop from `right` and add to `left`, then we add `num` to `right`.
    - Otherwise add `num` to `left`.
3. Repeat step #1.
4. Return median:
    - If `left` and `right` length are equal return the combined values of `left` max value and `right` min value divided by 2.
    - Otherwise return the `right` min value.

Example with the values: `[1,5,2,4,3]`

- Add `1` to `right`
    - `left: []`
    - `right: [1]`
- `5` is higher than `right` min value of `1`, move `1` to `left` and add `5` to `right`.
    - `left: [-1]`
    - `right: [5]`
- `2` is less than `right` min value of `5`, add `2` to `right`.
    - `left: [-1]`
    - `right: [2,5]`
- `4` is higher than `right` min value of `2`, move `2` to `left` and add `4` to `right`.
    - `left: [-2,-1]`
    - `right: [4,5]`
- `3` is less than `right` min value of `4`, add `3` to `right`.
    - `left: [-2,-1]`
    - `right: [3,5,4]`
- Return `3` as median.
