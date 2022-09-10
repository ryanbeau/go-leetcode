# Problem
[Leetcode - 208. Implement Trie (Medium)](https://leetcode.com/problems/implement-trie-prefix-tree/)

## Takeaway
- With the constraints of lowercase english letters we can safely assume only `'a'` to `'z'`.
  > `word` and `prefix` consist only of lowercase English letters
- If the Trie had only the method `Search` we could use a dictionary for the strings, but because of `StartsWith` prefix we need to ensure we use the `'a'` to `'z'` trie data structure.

## Solution
- Within the Trie we can store a 26 length array of the children ( `'a'` to `'z'`).
- With `Insert` we need to flag the final Trie to indicate it's the end of a word.
- Logic used for `Search` and `StartsWith` will be mostly identical we can create a shared method for this to return the last Trie.
  - `Search` will return false if the Trie is `nil`, otherwise it returns boolean flag set from `Insert`.
  - `StartsWith` will return true the Trie is not `nil`.

Below is an ascii diagram of a Trie with the inserted words `"bob"`, `"star"`, `"stop"` & `"stove"` (note: each parent node would have an array of 26 Trie children for the letters `'a'` to `'z'`):
```
           {root}
            /  \
          [b]  [s]
          /      \
        [o]      [t]
        /        / \
      [b]      [a] [o]
               /   / \
             [r] [p] [v]
                       \
                       [e]
```
