# [3. Longest Substring Without Repeating Characters](https://leetcode.com/problems/longest-substring-without-repeating-characters/)

Given a string, find the length of the longest substring without repeating characters.

## Examples

Example 1:

```
Input: "abcabcbb"
Output: 3
Explanation: The answer is "abc", with the length of 3.
```

Example 2:

```
Input: "bbbbb"
Output: 1
Explanation: The answer is "b", with the length of 1.
```

Example 3:

```
Input: "pwwkew"
Output: 3
Explanation: The answer is "wke", with the length of 3. Note that the answer must be a substring, "pwke" is a subsequence and not a substring.
```

## Approach

Use a `slide window` to represent the substring, and a `hashmap` to store the last seen locations of the characters in the slide window and check for duplicates

## Pseudocode

```
LENGTH-LONGEST-SUBSTRING(s)
  max = 1
  l = length(s)
  map = HashMap<char, lastSeenIndex>
  for i < l & j < l
    if m[s[j]] is in map
      remove map[s[i]] from map
      i++
    else
      map[s[j]] = j
      j++
      if (j-i) > max
        max = j-i
      end if
    end if
  end for
  return max
END
```
