# [215. Kth Largest Element in an Array](https://leetcode.com/problems/kth-largest-element-in-an-array/)

Find the kth largest element in an unsorted array. Note that it is the kth largest element in the sorted order, not the kth distinct element.

## Examples

Example 1:

```
Input: [3,2,1,5,6,4] and k = 2
Output: 5
```

Example 2:

```
Input: [3,2,3,1,2,4,5,5,6] and k = 4
Output: 4
```

**Note**:

You may assume k is always valid, 1 ≤ k ≤ array's length.

## Pseudocode
```
FIND-KTH-LARGEST(A, k)
    n = length(A)
    return RANDOMIZED-SELECT(A, 0, n-1, n-k+1)
END


RANDOMIZED-SELECT(A, l, r, i)
    if l = r
        return A[l]
    q = RANDOMIZED-PARTITION(A, l, r)
    k = q - l + 1
    if i = k
        return A[q]
    else if i < k
        return RANDOMIZED-SELECT(A, l, q - 1, i)
    else 
        return RANDOMIZED-SELECT(A, q, r, i - k + 1)
END

RANDOMIZED-PARTITION(A, l, r)
    i = random(l, r)
    exchange A[i] and A[r]
    return PARTITION(A, l, r)
END

PARTITION(A, l, r)
    x = A[r]
    i = l - 1
    for j = l to r - 1
        if A[j] <= x
            i += 1
            exchange A[i] and A[j]
    exchange A[i+1] and A[r]
    return i+1
END

```