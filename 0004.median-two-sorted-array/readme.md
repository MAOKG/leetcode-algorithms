# [4. Median of Two Sorted Arrays](https://leetcode.com/problems/median-of-two-sorted-arrays/)

There are two sorted arrays nums1 and nums2 of size m and n respectively.

Find the median of the two sorted arrays. The overall run time complexity should be O(log (m+n)).

## Examples

Example 1:

```
nums1 = [1, 3]
nums2 = [2]
The median is 2.0
```

Example 2:

```
nums1 = [1, 2]
nums2 = [3, 4]
The median is (2 + 3)/2 = 2.5
```

## Approach

Partitioning `nums1`, `nums2` into two group of halves, such that `nums1[i] >= nums2[j-1] && nums2[j] >= nums1[i-1]`. Becasue `i`, `j` partition the two arrays into halves, we also have `j = (l1 + l2 + 1) / 2 - i`. (plus one in the case of odd total length)

Now, the goal is to find the index `i` in `[0, l1)` that satisfies the above conditions. To do so, we can use a binary search. Once we found the index `i`, we can find the median depends on if `(l1 + l2)` is odd or not.

If it is odd, then median is the max of the first half `max(nums1[i-1], nums2[j-1])`; Else, median equals to the average between max of first half and min of second half `(max(nums1[i-1], nums2[j-1]) + min(nums1[i], nums2[j])) / 2`

## Pseudocode

```
MEDIAN-TWO-ARRAYS(X, Y)
  l1, l2 = length(X), length(Y)
  HANDLE-EDGE-CASES...
  if l1 <= l2
    return FIND-MEDIAN(X, Y, l1, l2, 0, l1)
  else
    return FIND-MEDIAN(Y, X, l2, l1, 0, l2)
  end if
END

FIND-MEDIAN(A, B, l1, l2, low, heigh)
  HANDLE-EDGE-CASES...
  i = (low + heigh)/2
  j = (l1 + l2 + 1)/2 - i
  isEven = (l1 + l2)%2 == 0
  if A[i-1] > B[j]
    return FIND-MEDIAN(A, B, l1, l2, low, i)
  else if B[j-1] > a[i]
    return FIND-MEDIAN(A, B, l1, l2, i+1, heigh)
  else
    maxFirstHalf = MAX(A[i-1], B[j-1])
    minSecondHalf = MIN(A[i], B[j])
    return GET-MEDIAN(maxFirstHalf, minSecondHalf, isEven)
  end if
END

GET-MEDIAN(maxFirstHalf, minSecondHalf, isEven)
  if isEven == false
    return maxFirstHalf
  else
    return (maxFirstHalf + minSecondHalf)/2
  end if
END
```
